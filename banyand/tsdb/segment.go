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
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/apache/skywalking-banyandb/banyand/kv"
	"github.com/apache/skywalking-banyandb/banyand/tsdb/bucket"
	"github.com/apache/skywalking-banyandb/pkg/logger"
	"github.com/apache/skywalking-banyandb/pkg/timestamp"
)

var ErrEndOfSegment = errors.New("reached the end of the segment")

type segment struct {
	id     uint16
	path   string
	suffix string

	blockCloseCh chan uint16
	globalIndex  kv.Store
	l            *logger.Logger
	timestamp.TimeRange
	bucket.Reporter
	blockController     *blockController
	blockManageStrategy *bucket.Strategy
}

func openSegment(ctx context.Context, startTime time.Time, path, suffix string, segmentSize, blockSize IntervalRule, blockQueue bucket.Queue) (s *segment, err error) {
	suffixInteger, err := strconv.Atoi(suffix)
	if err != nil {
		return nil, err
	}
	// TODO: fix id overflow
	id := uint16(segmentSize.Unit)<<12 | ((uint16(suffixInteger) << 4) >> 4)
	timeRange := timestamp.NewTimeRange(startTime, segmentSize.NextTime(startTime), true, false)
	l := logger.Fetch(ctx, "segment")
	s = &segment{
		id:              id,
		path:            path,
		suffix:          suffix,
		l:               l,
		blockController: newBlockController(id, path, timeRange, blockSize, l, blockQueue),
		TimeRange:       timeRange,
		Reporter:        bucket.NewTimeBasedReporter(timeRange),
		blockCloseCh:    make(chan uint16, 24),
	}
	err = s.blockController.open(context.WithValue(ctx, logger.ContextKey, s.l))
	if err != nil {
		return nil, err
	}

	indexPath, err := mkdir(globalIndexTemplate, path)
	if err != nil {
		return nil, err
	}
	if s.globalIndex, err = kv.OpenStore(0, indexPath, kv.StoreWithLogger(s.l)); err != nil {
		return nil, err
	}
	s.blockManageStrategy, err = bucket.NewStrategy(s.blockController, bucket.WithLogger(s.l))
	if err != nil {
		return nil, err
	}
	s.blockManageStrategy.Run()
	s.blockController.listenToBlockClosing(s.blockCloseCh)
	return s, nil
}

func (s *segment) close() {
	s.blockManageStrategy.Close()
	s.blockController.close()
	s.globalIndex.Close()
	s.Stop()
}

func (s *segment) notifyCloseBlock(id uint16) {
	s.blockCloseCh <- id
}

func (s segment) String() string {
	return s.Reporter.String()
}

type blockController struct {
	sync.RWMutex
	segID        uint16
	location     string
	segTimeRange timestamp.TimeRange
	blockSize    IntervalRule
	lst          []*block
	blockQueue   bucket.Queue
	closing      chan struct{}

	l *logger.Logger
}

func newBlockController(segID uint16, location string, segTimeRange timestamp.TimeRange,
	blockSize IntervalRule, l *logger.Logger, blockQueue bucket.Queue) *blockController {
	return &blockController{
		segID:        segID,
		location:     location,
		blockSize:    blockSize,
		segTimeRange: segTimeRange,
		closing:      make(chan struct{}),
		blockQueue:   blockQueue,
		l:            l,
	}
}

func (bc *blockController) Current() bucket.Reporter {
	bc.RLock()
	defer bc.RUnlock()
	now := time.Now()
	for _, s := range bc.lst {
		if s.suffix == bc.Format(now) {
			return s
		}
	}
	// return the latest segment before now
	if len(bc.lst) > 0 {
		return bc.lst[len(bc.lst)-1]
	}
	return nil
}

func (bc *blockController) Next() (bucket.Reporter, error) {
	b := bc.Current().(*block)
	reporter, err := bc.create(context.TODO(),
		bc.blockSize.NextTime(b.Start))
	if errors.Is(err, ErrEndOfSegment) {
		return nil, bucket.ErrNoMoreBucket
	}
	return reporter, err
}

func (bc *blockController) OnMove(prev bucket.Reporter, next bucket.Reporter) {
	bc.blockQueue.Push(blockIDAndSegID{
		segID:   bc.segID,
		blockID: prev.(*block).blockID,
	})
}

func (bc *blockController) listenToBlockClosing(closeCh <-chan uint16) {
	go func() {
		for {
			select {
			case id := <-closeCh:
				b := bc.removeBlock(id)
				if b != nil {
					b.close()
					continue
				}
				bc.l.Warn().Uint16("blockID", id).Uint16("segID", id).
					Msg("block is absent on removing it from the segment")
			case <-bc.closing:
				return
			}
		}
	}()
}

func (bc *blockController) removeBlock(blockID uint16) *block {
	bc.Lock()
	defer bc.Unlock()
	for i := range bc.lst {
		b := bc.lst[i]
		if b.blockID == blockID {
			//remove the block from the internal lst
			bc.lst = append(bc.lst[:i], bc.lst[i+1:]...)
			return b
		}
	}
	return nil
}

func (bc *blockController) Format(tm time.Time) string {
	switch bc.blockSize.Unit {
	case HOUR:
		return tm.Format(blockHourFormat)
	case DAY:
		return tm.Format(blockDayFormat)
	case MILLISECOND:
		return tm.Format(millisecondFormat)
	}
	panic("invalid interval unit")
}

func (bc *blockController) Parse(value string) (time.Time, error) {
	switch bc.blockSize.Unit {
	case HOUR:
		return time.Parse(blockHourFormat, value)
	case DAY:
		return time.Parse(blockDayFormat, value)
	case MILLISECOND:
		return time.Parse(millisecondFormat, value)
	}
	panic("invalid interval unit")
}

func (bc *blockController) span(timeRange timestamp.TimeRange) (bb []*block) {
	bc.RLock()
	defer bc.RUnlock()
	last := len(bc.lst) - 1
	for i := range bc.lst {
		b := bc.lst[last-i]
		if b.Overlapping(timeRange) {
			bb = append(bb, b)
		}
	}
	return bb
}

func (bc *blockController) get(blockID uint16) *block {
	bc.RLock()
	defer bc.RUnlock()
	last := len(bc.lst) - 1
	for i := range bc.lst {
		b := bc.lst[last-i]
		if b.blockID == blockID {
			return b
		}
	}
	return nil
}

func (bc *blockController) startTime(suffix string) (time.Time, error) {
	t, err := bc.Parse(suffix)
	if err != nil {
		return time.Time{}, err
	}
	startTime := bc.segTimeRange.Start
	switch bc.blockSize.Unit {
	case HOUR:
		return time.Date(startTime.Year(), startTime.Month(),
			startTime.Day(), t.Hour(), 0, 0, 0, startTime.Location()), nil
	case DAY:
		return time.Date(startTime.Year(), startTime.Month(),
			t.Day(), t.Hour(), 0, 0, 0, startTime.Location()), nil
	case MILLISECOND:
		return time.Parse(millisecondFormat, suffix)
	}
	panic("invalid interval unit")
}

func (bc *blockController) open(ctx context.Context) error {
	err := WalkDir(
		bc.location,
		segPathPrefix,
		func(suffix, absolutePath string) error {
			_, err := bc.load(ctx, suffix, absolutePath)
			return err
		})
	if err != nil {
		return err
	}
	if bc.Current() == nil {
		_, err = bc.create(ctx, time.Now())
		if err != nil {
			return err
		}
	}
	return nil
}

func (bc *blockController) create(ctx context.Context, startTime time.Time) (*block, error) {
	if startTime.Before(bc.segTimeRange.Start) {
		startTime = bc.segTimeRange.Start
	}
	if !startTime.Before(bc.segTimeRange.End) {
		return nil, ErrEndOfSegment
	}
	suffix := bc.Format(startTime)
	segPath, err := mkdir(blockTemplate, bc.location, suffix)
	if err != nil {
		return nil, err
	}
	return bc.load(ctx, suffix, segPath)
}

func (bc *blockController) load(ctx context.Context, suffix, path string) (b *block, err error) {
	starTime, err := bc.startTime(suffix)
	if err != nil {
		return nil, err
	}
	if b, err = newBlock(ctx, blockOpts{
		segID:     bc.segID,
		path:      path,
		startTime: starTime,
		suffix:    suffix,
		blockSize: bc.blockSize,
	}); err != nil {
		return nil, err
	}
	{
		bc.Lock()
		defer bc.Unlock()
		bc.lst = append(bc.lst, b)
	}
	return b, nil
}

func (bc *blockController) close() {
	bc.closing <- struct{}{}
	for _, s := range bc.lst {
		s.close()
	}
}