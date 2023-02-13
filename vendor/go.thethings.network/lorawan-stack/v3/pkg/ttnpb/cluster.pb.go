// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: lorawan-stack/api/cluster.proto

package ttnpb

import (
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

// PeerInfo
type PeerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Port on which the gRPC server is exposed.
	GrpcPort uint32 `protobuf:"varint,1,opt,name=grpc_port,json=grpcPort,proto3" json:"grpc_port,omitempty"`
	// Indicates whether the gRPC server uses TLS.
	Tls bool `protobuf:"varint,2,opt,name=tls,proto3" json:"tls,omitempty"`
	// Roles of the peer.
	Roles []ClusterRole `protobuf:"varint,3,rep,packed,name=roles,proto3,enum=ttn.lorawan.v3.ClusterRole" json:"roles,omitempty"`
	// Tags of the peer
	Tags map[string]string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PeerInfo) Reset() {
	*x = PeerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lorawan_stack_api_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerInfo) ProtoMessage() {}

func (x *PeerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_lorawan_stack_api_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerInfo.ProtoReflect.Descriptor instead.
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return file_lorawan_stack_api_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *PeerInfo) GetGrpcPort() uint32 {
	if x != nil {
		return x.GrpcPort
	}
	return 0
}

func (x *PeerInfo) GetTls() bool {
	if x != nil {
		return x.Tls
	}
	return false
}

func (x *PeerInfo) GetRoles() []ClusterRole {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *PeerInfo) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_lorawan_stack_api_cluster_proto protoreflect.FileDescriptor

var file_lorawan_stack_api_cluster_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76,
	0x33, 0x1a, 0x1d, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdd, 0x01, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a,
	0x09, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x67, 0x72, 0x70, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6c,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x12, 0x31, 0x0a, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x74, 0x74,
	0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12,
	0x36, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x6f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e,
	0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x74,
	0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lorawan_stack_api_cluster_proto_rawDescOnce sync.Once
	file_lorawan_stack_api_cluster_proto_rawDescData = file_lorawan_stack_api_cluster_proto_rawDesc
)

func file_lorawan_stack_api_cluster_proto_rawDescGZIP() []byte {
	file_lorawan_stack_api_cluster_proto_rawDescOnce.Do(func() {
		file_lorawan_stack_api_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_lorawan_stack_api_cluster_proto_rawDescData)
	})
	return file_lorawan_stack_api_cluster_proto_rawDescData
}

var file_lorawan_stack_api_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_lorawan_stack_api_cluster_proto_goTypes = []interface{}{
	(*PeerInfo)(nil), // 0: ttn.lorawan.v3.PeerInfo
	nil,              // 1: ttn.lorawan.v3.PeerInfo.TagsEntry
	(ClusterRole)(0), // 2: ttn.lorawan.v3.ClusterRole
}
var file_lorawan_stack_api_cluster_proto_depIdxs = []int32{
	2, // 0: ttn.lorawan.v3.PeerInfo.roles:type_name -> ttn.lorawan.v3.ClusterRole
	1, // 1: ttn.lorawan.v3.PeerInfo.tags:type_name -> ttn.lorawan.v3.PeerInfo.TagsEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_lorawan_stack_api_cluster_proto_init() }
func file_lorawan_stack_api_cluster_proto_init() {
	if File_lorawan_stack_api_cluster_proto != nil {
		return
	}
	file_lorawan_stack_api_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_lorawan_stack_api_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerInfo); i {
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
			RawDescriptor: file_lorawan_stack_api_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lorawan_stack_api_cluster_proto_goTypes,
		DependencyIndexes: file_lorawan_stack_api_cluster_proto_depIdxs,
		MessageInfos:      file_lorawan_stack_api_cluster_proto_msgTypes,
	}.Build()
	File_lorawan_stack_api_cluster_proto = out.File
	file_lorawan_stack_api_cluster_proto_rawDesc = nil
	file_lorawan_stack_api_cluster_proto_goTypes = nil
	file_lorawan_stack_api_cluster_proto_depIdxs = nil
}
