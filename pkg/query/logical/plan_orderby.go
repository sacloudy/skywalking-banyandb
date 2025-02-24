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

package logical

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/pkg/errors"

	databasev1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/database/v1"
	modelv1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/model/v1"
	streamv1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/stream/v1"
	"github.com/apache/skywalking-banyandb/pkg/convert"
)

type UnresolvedOrderBy struct {
	sort                modelv1.Sort
	targetIndexRuleName string
}

func (u *UnresolvedOrderBy) Analyze(s Schema) (*OrderBy, error) {
	if u == nil {
		// return a default orderBy sub-plan
		return &OrderBy{
			Sort: modelv1.Sort_SORT_UNSPECIFIED,
		}, nil
	}

	if u.targetIndexRuleName == "" {
		return &OrderBy{
			Sort: u.sort,
		}, nil
	}

	defined, indexRule := s.IndexRuleDefined(u.targetIndexRuleName)
	if !defined {
		return nil, errors.Wrap(ErrIndexNotDefined, u.targetIndexRuleName)
	}

	projFieldSpecs, err := s.CreateTagRef(NewTags("", indexRule.GetTags()...))
	if err != nil {
		return nil, ErrTagNotDefined
	}

	return &OrderBy{
		Sort:      u.sort,
		Index:     indexRule,
		fieldRefs: projFieldSpecs[0],
	}, nil
}

type OrderBy struct {
	// orderByIndex describes the indexRule used to sort the elements/
	// It can be null since by default we may sort by created-time.
	Index *databasev1.IndexRule
	// while orderBySort describes the Sort direction
	Sort modelv1.Sort
	// TODO: support multiple tags. Currently only the first member will be used for sorting.
	fieldRefs []*TagRef
}

func (o *OrderBy) Equal(other interface{}) bool {
	if otherOrderBy, ok := other.(*OrderBy); ok {
		if o == nil && otherOrderBy == nil {
			return true
		}
		if o != nil && otherOrderBy == nil || o == nil && otherOrderBy != nil {
			return false
		}
		return o.Sort == otherOrderBy.Sort &&
			o.Index.GetMetadata().GetName() == otherOrderBy.Index.GetMetadata().GetName()
	}

	return false
}

func (o *OrderBy) String() string {
	return fmt.Sprintf("OrderBy: %v, sort=%s", o.Index.GetTags(), o.Sort.String())
}

func NewOrderBy(indexRuleName string, sort modelv1.Sort) *UnresolvedOrderBy {
	return &UnresolvedOrderBy{
		sort:                sort,
		targetIndexRuleName: indexRuleName,
	}
}

func getRawTagValue(typedPair *modelv1.Tag) ([]byte, error) {
	switch v := typedPair.GetValue().Value.(type) {
	case *modelv1.TagValue_Str:
		return []byte(v.Str.GetValue()), nil
	case *modelv1.TagValue_Int:
		return convert.Int64ToBytes(v.Int.GetValue()), nil
	default:
		return nil, errors.New("unsupported data types")
	}
}

// SortedByIndex is used to test whether the given entities are sorted by the sortDirection
// The given entities MUST satisfy both the positive check and the negative check for the reversed direction
func SortedByIndex(elements []*streamv1.Element, tagFamilyIdx, tagIdx int, sortDirection modelv1.Sort) bool {
	if modelv1.Sort_SORT_UNSPECIFIED == sortDirection {
		sortDirection = modelv1.Sort_SORT_ASC
	}
	if len(elements) == 1 {
		return true
	}
	return sort.SliceIsSorted(elements, sortByIndex(elements, tagFamilyIdx, tagIdx, sortDirection)) &&
		!sort.SliceIsSorted(elements, sortByIndex(elements, tagFamilyIdx, tagIdx, reverseSortDirection(sortDirection)))
}

func SortedByTimestamp(elements []*streamv1.Element, sortDirection modelv1.Sort) bool {
	if modelv1.Sort_SORT_UNSPECIFIED == sortDirection {
		sortDirection = modelv1.Sort_SORT_ASC
	}
	if len(elements) == 1 {
		return true
	}
	return sort.SliceIsSorted(elements, sortByTimestamp(elements, sortDirection)) &&
		!sort.SliceIsSorted(elements, sortByTimestamp(elements, reverseSortDirection(sortDirection)))
}

func reverseSortDirection(sort modelv1.Sort) modelv1.Sort {
	if sort == modelv1.Sort_SORT_DESC {
		return modelv1.Sort_SORT_ASC
	}
	return modelv1.Sort_SORT_DESC
}

func sortByIndex(entities []*streamv1.Element, tagFamilyIdx, tagIdx int, sortDirection modelv1.Sort) func(i, j int) bool {
	return func(i, j int) bool {
		iPair := entities[i].GetTagFamilies()[tagFamilyIdx].GetTags()[tagIdx]
		jPair := entities[j].GetTagFamilies()[tagFamilyIdx].GetTags()[tagIdx]
		lField, _ := getRawTagValue(iPair)
		rField, _ := getRawTagValue(jPair)
		comp := bytes.Compare(lField, rField)
		if sortDirection == modelv1.Sort_SORT_ASC {
			return comp == -1
		}
		return comp == 1
	}
}

func sortByTimestamp(entities []*streamv1.Element, sortDirection modelv1.Sort) func(i, j int) bool {
	return func(i, j int) bool {
		iPair := entities[i].GetTimestamp().AsTime().UnixNano()
		jPair := entities[j].GetTimestamp().AsTime().UnixNano()
		if sortDirection == modelv1.Sort_SORT_DESC {
			return iPair > jPair
		}
		return iPair < jPair
	}
}
