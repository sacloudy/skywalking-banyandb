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
// source: banyandb/common/v1/common.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Catalog int32

const (
	Catalog_CATALOG_UNSPECIFIED Catalog = 0
	Catalog_CATALOG_STREAM      Catalog = 1
	Catalog_CATALOG_MEASURE     Catalog = 2
)

// Enum value maps for Catalog.
var (
	Catalog_name = map[int32]string{
		0: "CATALOG_UNSPECIFIED",
		1: "CATALOG_STREAM",
		2: "CATALOG_MEASURE",
	}
	Catalog_value = map[string]int32{
		"CATALOG_UNSPECIFIED": 0,
		"CATALOG_STREAM":      1,
		"CATALOG_MEASURE":     2,
	}
)

func (x Catalog) Enum() *Catalog {
	p := new(Catalog)
	*p = x
	return p
}

func (x Catalog) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Catalog) Descriptor() protoreflect.EnumDescriptor {
	return file_banyandb_common_v1_common_proto_enumTypes[0].Descriptor()
}

func (Catalog) Type() protoreflect.EnumType {
	return &file_banyandb_common_v1_common_proto_enumTypes[0]
}

func (x Catalog) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Catalog.Descriptor instead.
func (Catalog) EnumDescriptor() ([]byte, []int) {
	return file_banyandb_common_v1_common_proto_rawDescGZIP(), []int{0}
}

type IntervalRule_Unit int32

const (
	IntervalRule_UNIT_UNSPECIFIED IntervalRule_Unit = 0
	IntervalRule_UNIT_HOUR        IntervalRule_Unit = 1
	IntervalRule_UNIT_DAY         IntervalRule_Unit = 2
)

// Enum value maps for IntervalRule_Unit.
var (
	IntervalRule_Unit_name = map[int32]string{
		0: "UNIT_UNSPECIFIED",
		1: "UNIT_HOUR",
		2: "UNIT_DAY",
	}
	IntervalRule_Unit_value = map[string]int32{
		"UNIT_UNSPECIFIED": 0,
		"UNIT_HOUR":        1,
		"UNIT_DAY":         2,
	}
)

func (x IntervalRule_Unit) Enum() *IntervalRule_Unit {
	p := new(IntervalRule_Unit)
	*p = x
	return p
}

func (x IntervalRule_Unit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IntervalRule_Unit) Descriptor() protoreflect.EnumDescriptor {
	return file_banyandb_common_v1_common_proto_enumTypes[1].Descriptor()
}

func (IntervalRule_Unit) Type() protoreflect.EnumType {
	return &file_banyandb_common_v1_common_proto_enumTypes[1]
}

func (x IntervalRule_Unit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IntervalRule_Unit.Descriptor instead.
func (IntervalRule_Unit) EnumDescriptor() ([]byte, []int) {
	return file_banyandb_common_v1_common_proto_rawDescGZIP(), []int{1, 0}
}

// Metadata is for multi-tenant, multi-model use
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// group contains a set of options, like retention policy, max
	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	// name of the entity
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Id   uint32 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// readonly. create_revision is the revision of last creation on this key.
	CreateRevision int64 `protobuf:"varint,4,opt,name=create_revision,json=createRevision,proto3" json:"create_revision,omitempty"`
	// readonly. mod_revision is the revision of last modification on this key.
	ModRevision int64 `protobuf:"varint,5,opt,name=mod_revision,json=modRevision,proto3" json:"mod_revision,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_common_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_common_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_banyandb_common_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Metadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metadata) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Metadata) GetCreateRevision() int64 {
	if x != nil {
		return x.CreateRevision
	}
	return 0
}

func (x *Metadata) GetModRevision() int64 {
	if x != nil {
		return x.ModRevision
	}
	return 0
}

// IntervalRule is a structured duration
type IntervalRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unit can only be UNIT_HOUR or UNIT_DAY
	Unit IntervalRule_Unit `protobuf:"varint,1,opt,name=unit,proto3,enum=banyandb.common.v1.IntervalRule_Unit" json:"unit,omitempty"`
	Num  uint32            `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *IntervalRule) Reset() {
	*x = IntervalRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_common_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntervalRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntervalRule) ProtoMessage() {}

func (x *IntervalRule) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_common_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntervalRule.ProtoReflect.Descriptor instead.
func (*IntervalRule) Descriptor() ([]byte, []int) {
	return file_banyandb_common_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *IntervalRule) GetUnit() IntervalRule_Unit {
	if x != nil {
		return x.Unit
	}
	return IntervalRule_UNIT_UNSPECIFIED
}

func (x *IntervalRule) GetNum() uint32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type ResourceOpts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// shard_num is the number of shards
	ShardNum uint32 `protobuf:"varint,1,opt,name=shard_num,json=shardNum,proto3" json:"shard_num,omitempty"`
	// block_interval indicates the length of a block
	// block_interval should be less than or equal to segment_interval
	BlockInterval *IntervalRule `protobuf:"bytes,2,opt,name=block_interval,json=blockInterval,proto3" json:"block_interval,omitempty"`
	// segment_interval indicates the length of a segment
	SegmentInterval *IntervalRule `protobuf:"bytes,3,opt,name=segment_interval,json=segmentInterval,proto3" json:"segment_interval,omitempty"`
	// ttl indicates time to live, how long the data will be cached
	Ttl *IntervalRule `protobuf:"bytes,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *ResourceOpts) Reset() {
	*x = ResourceOpts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_common_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceOpts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceOpts) ProtoMessage() {}

func (x *ResourceOpts) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_common_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceOpts.ProtoReflect.Descriptor instead.
func (*ResourceOpts) Descriptor() ([]byte, []int) {
	return file_banyandb_common_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceOpts) GetShardNum() uint32 {
	if x != nil {
		return x.ShardNum
	}
	return 0
}

func (x *ResourceOpts) GetBlockInterval() *IntervalRule {
	if x != nil {
		return x.BlockInterval
	}
	return nil
}

func (x *ResourceOpts) GetSegmentInterval() *IntervalRule {
	if x != nil {
		return x.SegmentInterval
	}
	return nil
}

func (x *ResourceOpts) GetTtl() *IntervalRule {
	if x != nil {
		return x.Ttl
	}
	return nil
}

// Group is an internal object for Group management
type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// metadata define the group's identity
	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// catalog denotes which type of data the group contains
	Catalog Catalog `protobuf:"varint,2,opt,name=catalog,proto3,enum=banyandb.common.v1.Catalog" json:"catalog,omitempty"`
	// resourceOpts indicates the structure of the underlying kv storage
	ResourceOpts *ResourceOpts `protobuf:"bytes,3,opt,name=resource_opts,json=resourceOpts,proto3" json:"resource_opts,omitempty"`
	// updated_at indicates when resources of the group are updated
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_common_v1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_common_v1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_banyandb_common_v1_common_proto_rawDescGZIP(), []int{3}
}

func (x *Group) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Group) GetCatalog() Catalog {
	if x != nil {
		return x.Catalog
	}
	return Catalog_CATALOG_UNSPECIFIED
}

func (x *Group) GetResourceOpts() *ResourceOpts {
	if x != nil {
		return x.ResourceOpts
	}
	return nil
}

func (x *Group) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_banyandb_common_v1_common_proto protoreflect.FileDescriptor

var file_banyandb_common_v1_common_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x99, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x27, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x5f,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x6d, 0x6f, 0x64, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xa9, 0x01, 0x0a, 0x0c,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x43, 0x0a, 0x04,
	0x75, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x62, 0x61, 0x6e,
	0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x55, 0x6e, 0x69,
	0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x75, 0x6e, 0x69,
	0x74, 0x12, 0x19, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x39, 0x0a, 0x04,
	0x55, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e,
	0x49, 0x54, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e, 0x49,
	0x54, 0x5f, 0x44, 0x41, 0x59, 0x10, 0x02, 0x22, 0x9c, 0x02, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x09, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x2a, 0x02, 0x20, 0x00, 0x52, 0x08, 0x73, 0x68, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x51,
	0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x12, 0x55, 0x0a, 0x10, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x61,
	0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x3c, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x22, 0x8e, 0x02, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x42, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x52, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x4f, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4f, 0x70, 0x74, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x4b, 0x0a, 0x07, 0x43, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x41, 0x54, 0x41, 0x4c, 0x4f, 0x47, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43,
	0x41, 0x54, 0x41, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x10, 0x01, 0x12,
	0x13, 0x0a, 0x0f, 0x43, 0x41, 0x54, 0x41, 0x4c, 0x4f, 0x47, 0x5f, 0x4d, 0x45, 0x41, 0x53, 0x55,
	0x52, 0x45, 0x10, 0x02, 0x42, 0x6e, 0x0a, 0x28, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x61,
	0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x61,
	0x63, 0x68, 0x65, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x62,
	0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_banyandb_common_v1_common_proto_rawDescOnce sync.Once
	file_banyandb_common_v1_common_proto_rawDescData = file_banyandb_common_v1_common_proto_rawDesc
)

func file_banyandb_common_v1_common_proto_rawDescGZIP() []byte {
	file_banyandb_common_v1_common_proto_rawDescOnce.Do(func() {
		file_banyandb_common_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_banyandb_common_v1_common_proto_rawDescData)
	})
	return file_banyandb_common_v1_common_proto_rawDescData
}

var file_banyandb_common_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_banyandb_common_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_banyandb_common_v1_common_proto_goTypes = []interface{}{
	(Catalog)(0),                  // 0: banyandb.common.v1.Catalog
	(IntervalRule_Unit)(0),        // 1: banyandb.common.v1.IntervalRule.Unit
	(*Metadata)(nil),              // 2: banyandb.common.v1.Metadata
	(*IntervalRule)(nil),          // 3: banyandb.common.v1.IntervalRule
	(*ResourceOpts)(nil),          // 4: banyandb.common.v1.ResourceOpts
	(*Group)(nil),                 // 5: banyandb.common.v1.Group
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_banyandb_common_v1_common_proto_depIdxs = []int32{
	1, // 0: banyandb.common.v1.IntervalRule.unit:type_name -> banyandb.common.v1.IntervalRule.Unit
	3, // 1: banyandb.common.v1.ResourceOpts.block_interval:type_name -> banyandb.common.v1.IntervalRule
	3, // 2: banyandb.common.v1.ResourceOpts.segment_interval:type_name -> banyandb.common.v1.IntervalRule
	3, // 3: banyandb.common.v1.ResourceOpts.ttl:type_name -> banyandb.common.v1.IntervalRule
	2, // 4: banyandb.common.v1.Group.metadata:type_name -> banyandb.common.v1.Metadata
	0, // 5: banyandb.common.v1.Group.catalog:type_name -> banyandb.common.v1.Catalog
	4, // 6: banyandb.common.v1.Group.resource_opts:type_name -> banyandb.common.v1.ResourceOpts
	6, // 7: banyandb.common.v1.Group.updated_at:type_name -> google.protobuf.Timestamp
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_banyandb_common_v1_common_proto_init() }
func file_banyandb_common_v1_common_proto_init() {
	if File_banyandb_common_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_banyandb_common_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_banyandb_common_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntervalRule); i {
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
		file_banyandb_common_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceOpts); i {
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
		file_banyandb_common_v1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
			RawDescriptor: file_banyandb_common_v1_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_banyandb_common_v1_common_proto_goTypes,
		DependencyIndexes: file_banyandb_common_v1_common_proto_depIdxs,
		EnumInfos:         file_banyandb_common_v1_common_proto_enumTypes,
		MessageInfos:      file_banyandb_common_v1_common_proto_msgTypes,
	}.Build()
	File_banyandb_common_v1_common_proto = out.File
	file_banyandb_common_v1_common_proto_rawDesc = nil
	file_banyandb_common_v1_common_proto_goTypes = nil
	file_banyandb_common_v1_common_proto_depIdxs = nil
}
