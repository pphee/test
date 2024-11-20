// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: qdrant.proto

package qdrant

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

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qdrant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qdrant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_qdrant_proto_rawDescGZIP(), []int{0}
}

type HealthCheckReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Version string  `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Commit  *string `protobuf:"bytes,3,opt,name=commit,proto3,oneof" json:"commit,omitempty"`
}

func (x *HealthCheckReply) Reset() {
	*x = HealthCheckReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qdrant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckReply) ProtoMessage() {}

func (x *HealthCheckReply) ProtoReflect() protoreflect.Message {
	mi := &file_qdrant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckReply.ProtoReflect.Descriptor instead.
func (*HealthCheckReply) Descriptor() ([]byte, []int) {
	return file_qdrant_proto_rawDescGZIP(), []int{1}
}

func (x *HealthCheckReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *HealthCheckReply) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *HealthCheckReply) GetCommit() string {
	if x != nil && x.Commit != nil {
		return *x.Commit
	}
	return ""
}

var File_qdrant_proto protoreflect.FileDescriptor

var file_qdrant_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x1a, 0x19, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x14, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x14, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6a, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x06, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x32, 0x4f, 0x0a, 0x06, 0x51, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x0b,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1a, 0x2e, 0x71, 0x64,
	0x72, 0x61, 0x6e, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74,
	0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qdrant_proto_rawDescOnce sync.Once
	file_qdrant_proto_rawDescData = file_qdrant_proto_rawDesc
)

func file_qdrant_proto_rawDescGZIP() []byte {
	file_qdrant_proto_rawDescOnce.Do(func() {
		file_qdrant_proto_rawDescData = protoimpl.X.CompressGZIP(file_qdrant_proto_rawDescData)
	})
	return file_qdrant_proto_rawDescData
}

var file_qdrant_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_qdrant_proto_goTypes = []any{
	(*HealthCheckRequest)(nil), // 0: qdrant.HealthCheckRequest
	(*HealthCheckReply)(nil),   // 1: qdrant.HealthCheckReply
}
var file_qdrant_proto_depIdxs = []int32{
	0, // 0: qdrant.Qdrant.HealthCheck:input_type -> qdrant.HealthCheckRequest
	1, // 1: qdrant.Qdrant.HealthCheck:output_type -> qdrant.HealthCheckReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_qdrant_proto_init() }
func file_qdrant_proto_init() {
	if File_qdrant_proto != nil {
		return
	}
	file_collections_service_proto_init()
	file_points_service_proto_init()
	file_snapshots_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_qdrant_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HealthCheckRequest); i {
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
		file_qdrant_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*HealthCheckReply); i {
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
	file_qdrant_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_qdrant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qdrant_proto_goTypes,
		DependencyIndexes: file_qdrant_proto_depIdxs,
		MessageInfos:      file_qdrant_proto_msgTypes,
	}.Build()
	File_qdrant_proto = out.File
	file_qdrant_proto_rawDesc = nil
	file_qdrant_proto_goTypes = nil
	file_qdrant_proto_depIdxs = nil
}