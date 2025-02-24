// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package tsdb

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"

	"github.com/apache/skywalking-banyandb/api/common"
	"github.com/apache/skywalking-banyandb/banyand/kv"
	"github.com/apache/skywalking-banyandb/banyand/observability"
	"github.com/apache/skywalking-banyandb/banyand/tsdb/bucket"
	"github.com/apache/skywalking-banyandb/pkg/encoding"
	"github.com/apache/skywalking-banyandb/pkg/index"
	"github.com/apache/skywalking-banyandb/pkg/index/inverted"
	"github.com/apache/skywalking-banyandb/pkg/index/lsm"
	"github.com/apache/skywalking-banyandb/pkg/logger"
	"github.com/apache/skywalking-banyandb/pkg/timestamp"
)

const (
	componentMain              = "main"
	componentSecondInvertedIdx = "inverted"
	componentSecondLSMIdx      = "lsm"

	defaultMainMemorySize = 8 << 20
)

type block struct {
	path       string
	l          *logger.Logger
	queue      bucket.Queue
	suffix     string
	ref        *atomic.Int32
	closed     *atomic.Bool
	deleted    *atomic.Bool
	lock       sync.RWMutex
	position   common.Position
	memSize    int64
	lsmMemSize int64

	store         kv.TimeSeriesStore
	invertedIndex index.Store
	lsmIndex      index.Store
	closableLst   []io.Closer
	timestamp.TimeRange
	bucket.Reporter
	segID          uint16
	blockID        uint16
	encodingMethod EncodingMethod
	flushCh        chan struct{}
	stopCh         chan struct{}
}

type blockOpts struct {
	segID     uint16
	blockSize IntervalRule
	startTime time.Time
	suffix    string
	path      string
	queue     bucket.Queue
}

func newBlock(ctx context.Context, opts blockOpts) (b *block, err error) {
	suffixInteger, err := strconv.Atoi(opts.suffix)
	if err != nil {
		return nil, err
	}
	id := GenerateInternalID(opts.blockSize.Unit, suffixInteger)
	timeRange := timestamp.NewTimeRange(opts.startTime, opts.blockSize.NextTime(opts.startTime), true, false)
	clock, _ := timestamp.GetClock(ctx)
	b = &block{
		segID:     opts.segID,
		blockID:   id,
		path:      opts.path,
		l:         logger.Fetch(ctx, "block"),
		TimeRange: timeRange,
		Reporter:  bucket.NewTimeBasedReporter(timeRange, clock),
		flushCh:   make(chan struct{}, 1),
		ref:       &atomic.Int32{},
		closed:    &atomic.Bool{},
		deleted:   &atomic.Bool{},
		queue:     opts.queue,
	}
	b.options(ctx)
	position := ctx.Value(common.PositionKey)
	if position != nil {
		b.position = position.(common.Position)
	}

	return b, b.open()
}

func (b *block) options(ctx context.Context) {
	var options DatabaseOpts
	o := ctx.Value(optionsKey)
	if o != nil {
		options = o.(DatabaseOpts)
	}
	if options.EncodingMethod.EncoderPool == nil {
		options.EncodingMethod.EncoderPool = encoding.NewPlainEncoderPool("tsdb", 0)
	}
	if options.EncodingMethod.EncoderPool == nil {
		options.EncodingMethod.DecoderPool = encoding.NewPlainDecoderPool("tsdb", 0)
	}
	b.encodingMethod = options.EncodingMethod
	if options.BlockMemSize < 1 {
		b.memSize = defaultMainMemorySize
	} else {
		b.memSize = options.BlockMemSize
	}
	b.lsmMemSize = b.memSize / 8
	if b.lsmMemSize < defaultKVMemorySize {
		b.lsmMemSize = defaultKVMemorySize
	}
}

func (b *block) open() (err error) {
	if b.deleted.Load() {
		return nil
	}
	if b.store, err = kv.OpenTimeSeriesStore(
		0,
		path.Join(b.path, componentMain),
		kv.TSSWithEncoding(b.encodingMethod.EncoderPool, b.encodingMethod.DecoderPool),
		kv.TSSWithLogger(b.l.Named(componentMain)),
		kv.TSSWithMemTableSize(b.memSize),
		kv.TSSWithFlushCallback(func() {
			b.flushCh <- struct{}{}
		}),
	); err != nil {
		return err
	}
	b.closableLst = append(b.closableLst, b.store)
	if b.invertedIndex, err = inverted.NewStore(inverted.StoreOpts{
		Path:   path.Join(b.path, componentSecondInvertedIdx),
		Logger: b.l.Named(componentSecondInvertedIdx),
	}); err != nil {
		return err
	}
	if b.lsmIndex, err = lsm.NewStore(lsm.StoreOpts{
		Path:         path.Join(b.path, componentSecondLSMIdx),
		Logger:       b.l.Named(componentSecondLSMIdx),
		MemTableSize: b.lsmMemSize,
	}); err != nil {
		return err
	}
	b.closableLst = append(b.closableLst, b.invertedIndex, b.lsmIndex)
	b.ref.Store(0)
	stopCh := make(chan struct{})
	b.stopCh = stopCh
	go func() {
		for {
			select {
			case <-b.flushCh:
				b.flush()
			case <-stopCh:
				return
			}
		}
	}()
	b.closed.Store(false)
	return nil
}

func (b *block) delegate() (BlockDelegate, error) {
	if b.deleted.Load() {
		return nil, errors.WithMessagef(ErrBlockAbsent, "block %d is deleted", b.blockID)
	}
	if b.incRef() {
		return &bDelegate{
			delegate: b,
		}, nil
	}
	b.lock.Lock()
	defer b.lock.Unlock()
	b.queue.Push(BlockID{
		BlockID: b.blockID,
		SegID:   b.segID,
	})
	// TODO: remove the block which fails to open from the queue
	err := b.open()
	if err != nil {
		b.l.Error().Err(err).Stringer("block", b).Msg("fail to open block")
		return nil, err
	}
	b.incRef()
	return &bDelegate{
		delegate: b,
	}, nil
}

func (b *block) incRef() bool {
loop:
	if b.Closed() {
		return false
	}
	r := b.ref.Load()
	if b.ref.CompareAndSwap(r, r+1) {
		return true
	}
	runtime.Gosched()
	goto loop
}

func (b *block) Done() {
loop:
	r := b.ref.Load()
	if r < 1 {
		return
	}
	if b.ref.CompareAndSwap(r, r-1) {
		return
	}
	runtime.Gosched()
	goto loop
}

func (b *block) waitDone() {
loop:
	if b.ref.Load() < 1 {
		b.ref.Store(0)
		return
	}
	runtime.Gosched()
	goto loop
}

func (b *block) flush() {
	for i := 0; i < 10; i++ {
		err := b.invertedIndex.Flush()
		if err == nil {
			break
		}
		time.Sleep(time.Second)
		b.l.Warn().Err(err).Int("retried", i).Msg("failed to flush inverted index")
	}
}

func (b *block) close() {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.closed.Load() {
		return
	}
	b.closed.Store(true)
	b.waitDone()
	for _, closer := range b.closableLst {
		_ = closer.Close()
	}
	close(b.stopCh)
}

func (b *block) stopThenClose() {
	if b.Reporter != nil {
		b.Stop()
	}
	b.close()
}

func (b *block) delete() error {
	if b.deleted.Load() {
		return nil
	}
	b.deleted.Store(true)
	if b.Reporter != nil {
		b.Stop()
	}
	b.close()
	return os.RemoveAll(b.path)
}

func (b *block) Closed() bool {
	return b.closed.Load()
}

func (b *block) String() string {
	return fmt.Sprintf("BlockID-%d-%d", b.segID, b.blockID)
}

func (b *block) stats() (names []string, stats []observability.Statistics) {
	names = append(names, componentMain, componentSecondInvertedIdx, componentSecondLSMIdx)
	if b.Closed() {
		stats = make([]observability.Statistics, 3)
		return
	}
	stats = append(stats, b.store.Stats(), b.invertedIndex.Stats(), b.lsmIndex.Stats())
	return names, stats
}

type BlockDelegate interface {
	io.Closer
	contains(ts time.Time) bool
	write(key []byte, val []byte, ts time.Time) error
	writePrimaryIndex(field index.Field, id common.ItemID) error
	writeLSMIndex(fields []index.Field, id common.ItemID) error
	writeInvertedIndex(fields []index.Field, id common.ItemID) error
	dataReader() kv.TimeSeriesReader
	lsmIndexReader() index.Searcher
	invertedIndexReader() index.Searcher
	primaryIndexReader() index.FieldIterable
	identity() (segID uint16, blockID uint16)
	startTime() time.Time
	String() string
}

var _ BlockDelegate = (*bDelegate)(nil)

type bDelegate struct {
	delegate *block
}

func (d *bDelegate) dataReader() kv.TimeSeriesReader {
	return d.delegate.store
}

func (d *bDelegate) lsmIndexReader() index.Searcher {
	return d.delegate.lsmIndex
}

func (d *bDelegate) invertedIndexReader() index.Searcher {
	return d.delegate.invertedIndex
}

func (d *bDelegate) primaryIndexReader() index.FieldIterable {
	return d.delegate.lsmIndex
}

func (d *bDelegate) startTime() time.Time {
	return d.delegate.Start
}

func (d *bDelegate) identity() (segID uint16, blockID uint16) {
	return d.delegate.segID, d.delegate.blockID
}

func (d *bDelegate) write(key []byte, val []byte, ts time.Time) error {
	return d.delegate.store.Put(key, val, uint64(ts.UnixNano()))
}

func (d *bDelegate) writePrimaryIndex(field index.Field, id common.ItemID) error {
	return d.delegate.lsmIndex.Write([]index.Field{field}, id)
}

func (d *bDelegate) writeLSMIndex(fields []index.Field, id common.ItemID) error {
	return d.delegate.lsmIndex.Write(fields, id)
}

func (d *bDelegate) writeInvertedIndex(fields []index.Field, id common.ItemID) error {
	return d.delegate.invertedIndex.Write(fields, id)
}

func (d *bDelegate) contains(ts time.Time) bool {
	return d.delegate.Contains(uint64(ts.UnixNano()))
}

func (d *bDelegate) String() string {
	return d.delegate.String()
}

func (d *bDelegate) Close() error {
	d.delegate.Done()
	return nil
}
