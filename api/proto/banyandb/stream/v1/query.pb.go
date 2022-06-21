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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: banyandb/stream/v1/query.proto

package v1

import (
	v11 "github.com/apache/skywalking-banyandb/api/proto/banyandb/common/v1"
	v1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/model/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Element represents
// (stream context) a Span defined in Google Dapper paper or equivalently a Segment in Skywalking.
// (Log context) a log
type Element struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// element_id could be span_id of a Span or segment_id of a Segment in the context of stream
	ElementId string `protobuf:"bytes,1,opt,name=element_id,json=elementId,proto3" json:"element_id,omitempty"`
	// timestamp represents a millisecond
	// 1) either the start time of a Span/Segment,
	// 2) or the timestamp of a log
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// fields contains all indexed Field. Some typical names,
	// - stream_id
	// - duration
	// - service_name
	// - service_instance_id
	// - end_time_milliseconds
	TagFamilies []*v1.TagFamily `protobuf:"bytes,3,rep,name=tag_families,json=tagFamilies,proto3" json:"tag_families,omitempty"`
}

func (x *Element) Reset() {
	*x = Element{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_stream_v1_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Element) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Element) ProtoMessage() {}

func (x *Element) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_stream_v1_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Element.ProtoReflect.Descriptor instead.
func (*Element) Descriptor() ([]byte, []int) {
	return file_banyandb_stream_v1_query_proto_rawDescGZIP(), []int{0}
}

func (x *Element) GetElementId() string {
	if x != nil {
		return x.ElementId
	}
	return ""
}

func (x *Element) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Element) GetTagFamilies() []*v1.TagFamily {
	if x != nil {
		return x.TagFamilies
	}
	return nil
}

// QueryResponse is the response for a query to the Query module.
type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// elements are the actual data returned
	Elements []*Element `protobuf:"bytes,1,rep,name=elements,proto3" json:"elements,omitempty"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_stream_v1_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_stream_v1_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_banyandb_stream_v1_query_proto_rawDescGZIP(), []int{1}
}

func (x *QueryResponse) GetElements() []*Element {
	if x != nil {
		return x.Elements
	}
	return nil
}

// QueryRequest is the request contract for query.
type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// metadata is required
	Metadata *v11.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// time_range is a range query with begin/end time of entities in the timeunit of milliseconds.
	// In the context of stream, it represents the range of the `startTime` for spans/segments,
	// while in the context of Log, it means the range of the timestamp(s) for logs.
	// it is always recommended to specify time range for performance reason
	TimeRange *v1.TimeRange `protobuf:"bytes,2,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	// offset is used to support pagination, together with the following limit
	Offset uint32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	// limit is used to impose a boundary on the number of records being returned
	Limit uint32 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// order_by is given to specify the sort for a field. So far, only fields in the type of Integer are supported
	OrderBy *v1.QueryOrder `protobuf:"bytes,5,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// tag_families are indexed.
	Criteria []*v1.Criteria `protobuf:"bytes,6,rep,name=criteria,proto3" json:"criteria,omitempty"`
	// projection can be used to select the key names of the element in the response
	Projection *v1.TagProjection `protobuf:"bytes,7,opt,name=projection,proto3" json:"projection,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_stream_v1_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_stream_v1_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_banyandb_stream_v1_query_proto_rawDescGZIP(), []int{2}
}

func (x *QueryRequest) GetMetadata() *v11.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *QueryRequest) GetTimeRange() *v1.TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

func (x *QueryRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *QueryRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *QueryRequest) GetOrderBy() *v1.QueryOrder {
	if x != nil {
		return x.OrderBy
	}
	return nil
}

func (x *QueryRequest) GetCriteria() []*v1.Criteria {
	if x != nil {
		return x.Criteria
	}
	return nil
}

func (x *QueryRequest) GetProjection() *v1.TagProjection {
	if x != nil {
		return x.Projection
	}
	return nil
}

var File_banyandb_stream_v1_query_proto protoreflect.FileDescriptor

var file_banyandb_stream_v1_query_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x07, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3f, 0x0a, 0x0c, 0x74, 0x61,
	0x67, 0x5f, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x0b,
	0x74, 0x61, 0x67, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x69, 0x65, 0x73, 0x22, 0x48, 0x0a, 0x0d, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x08,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xe8, 0x02, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61,
	0x6e, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x3b, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x37, 0x0a, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e,
	0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x52, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x40,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x6e, 0x0a, 0x28, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73,
	0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e,
	0x64, 0x62, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x5a, 0x42, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f,
	0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x62, 0x61, 0x6e, 0x79, 0x61,
	0x6e, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61,
	0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_banyandb_stream_v1_query_proto_rawDescOnce sync.Once
	file_banyandb_stream_v1_query_proto_rawDescData = file_banyandb_stream_v1_query_proto_rawDesc
)

func file_banyandb_stream_v1_query_proto_rawDescGZIP() []byte {
	file_banyandb_stream_v1_query_proto_rawDescOnce.Do(func() {
		file_banyandb_stream_v1_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_banyandb_stream_v1_query_proto_rawDescData)
	})
	return file_banyandb_stream_v1_query_proto_rawDescData
}

var file_banyandb_stream_v1_query_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_banyandb_stream_v1_query_proto_goTypes = []interface{}{
	(*Element)(nil),               // 0: banyandb.stream.v1.Element
	(*QueryResponse)(nil),         // 1: banyandb.stream.v1.QueryResponse
	(*QueryRequest)(nil),          // 2: banyandb.stream.v1.QueryRequest
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*v1.TagFamily)(nil),          // 4: banyandb.model.v1.TagFamily
	(*v11.Metadata)(nil),          // 5: banyandb.common.v1.Metadata
	(*v1.TimeRange)(nil),          // 6: banyandb.model.v1.TimeRange
	(*v1.QueryOrder)(nil),         // 7: banyandb.model.v1.QueryOrder
	(*v1.Criteria)(nil),           // 8: banyandb.model.v1.Criteria
	(*v1.TagProjection)(nil),      // 9: banyandb.model.v1.TagProjection
}
var file_banyandb_stream_v1_query_proto_depIdxs = []int32{
	3, // 0: banyandb.stream.v1.Element.timestamp:type_name -> google.protobuf.Timestamp
	4, // 1: banyandb.stream.v1.Element.tag_families:type_name -> banyandb.model.v1.TagFamily
	0, // 2: banyandb.stream.v1.QueryResponse.elements:type_name -> banyandb.stream.v1.Element
	5, // 3: banyandb.stream.v1.QueryRequest.metadata:type_name -> banyandb.common.v1.Metadata
	6, // 4: banyandb.stream.v1.QueryRequest.time_range:type_name -> banyandb.model.v1.TimeRange
	7, // 5: banyandb.stream.v1.QueryRequest.order_by:type_name -> banyandb.model.v1.QueryOrder
	8, // 6: banyandb.stream.v1.QueryRequest.criteria:type_name -> banyandb.model.v1.Criteria
	9, // 7: banyandb.stream.v1.QueryRequest.projection:type_name -> banyandb.model.v1.TagProjection
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_banyandb_stream_v1_query_proto_init() }
func file_banyandb_stream_v1_query_proto_init() {
	if File_banyandb_stream_v1_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_banyandb_stream_v1_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Element); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_banyandb_stream_v1_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_banyandb_stream_v1_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_banyandb_stream_v1_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_banyandb_stream_v1_query_proto_goTypes,
		DependencyIndexes: file_banyandb_stream_v1_query_proto_depIdxs,
		MessageInfos:      file_banyandb_stream_v1_query_proto_msgTypes,
	}.Build()
	File_banyandb_stream_v1_query_proto = out.File
	file_banyandb_stream_v1_query_proto_rawDesc = nil
	file_banyandb_stream_v1_query_proto_goTypes = nil
	file_banyandb_stream_v1_query_proto_depIdxs = nil
}
