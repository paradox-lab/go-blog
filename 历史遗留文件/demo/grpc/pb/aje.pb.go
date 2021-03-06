// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: aje.proto

package aje

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The request message containing the user's name.
type RunRequst struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *RunRequst) Reset() {
	*x = RunRequst{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aje_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunRequst) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunRequst) ProtoMessage() {}

func (x *RunRequst) ProtoReflect() protoreflect.Message {
	mi := &file_aje_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunRequst.ProtoReflect.Descriptor instead.
func (*RunRequst) Descriptor() ([]byte, []int) {
	return file_aje_proto_rawDescGZIP(), []int{0}
}

func (x *RunRequst) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

// The response message containing the greetings
type RunReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RunReply) Reset() {
	*x = RunReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aje_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunReply) ProtoMessage() {}

func (x *RunReply) ProtoReflect() protoreflect.Message {
	mi := &file_aje_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunReply.ProtoReflect.Descriptor instead.
func (*RunReply) Descriptor() ([]byte, []int) {
	return file_aje_proto_rawDescGZIP(), []int{1}
}

func (x *RunReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_aje_proto protoreflect.FileDescriptor

var file_aje_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x6a, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x6a, 0x65,
	0x22, 0x1f, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x22, 0x20, 0x0a, 0x08, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x32, 0x3b, 0x0a, 0x08, 0x53, 0x65, 0x6c, 0x65, 0x6e, 0x69, 0x75, 0x6d, 0x12,
	0x2f, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4a, 0x64, 0x43, 0x66, 0x73, 0x74, 0x12,
	0x0e, 0x2e, 0x61, 0x6a, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x61, 0x6a, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aje_proto_rawDescOnce sync.Once
	file_aje_proto_rawDescData = file_aje_proto_rawDesc
)

func file_aje_proto_rawDescGZIP() []byte {
	file_aje_proto_rawDescOnce.Do(func() {
		file_aje_proto_rawDescData = protoimpl.X.CompressGZIP(file_aje_proto_rawDescData)
	})
	return file_aje_proto_rawDescData
}

var file_aje_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aje_proto_goTypes = []interface{}{
	(*RunRequst)(nil), // 0: aje.RunRequst
	(*RunReply)(nil),  // 1: aje.RunReply
}
var file_aje_proto_depIdxs = []int32{
	0, // 0: aje.Selenium.UploadJdCfst:input_type -> aje.RunRequst
	1, // 1: aje.Selenium.UploadJdCfst:output_type -> aje.RunReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aje_proto_init() }
func file_aje_proto_init() {
	if File_aje_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aje_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunRequst); i {
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
		file_aje_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunReply); i {
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
			RawDescriptor: file_aje_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aje_proto_goTypes,
		DependencyIndexes: file_aje_proto_depIdxs,
		MessageInfos:      file_aje_proto_msgTypes,
	}.Build()
	File_aje_proto = out.File
	file_aje_proto_rawDesc = nil
	file_aje_proto_goTypes = nil
	file_aje_proto_depIdxs = nil
}
