// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: RefreshAuthorizationPolicyProtocol.proto

package hadoop_common

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

//*
//  Refresh service acl request.
type RefreshServiceAclRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefreshServiceAclRequestProto) Reset() {
	*x = RefreshServiceAclRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RefreshAuthorizationPolicyProtocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshServiceAclRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshServiceAclRequestProto) ProtoMessage() {}

func (x *RefreshServiceAclRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_RefreshAuthorizationPolicyProtocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshServiceAclRequestProto.ProtoReflect.Descriptor instead.
func (*RefreshServiceAclRequestProto) Descriptor() ([]byte, []int) {
	return file_RefreshAuthorizationPolicyProtocol_proto_rawDescGZIP(), []int{0}
}

//*
// void response
type RefreshServiceAclResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefreshServiceAclResponseProto) Reset() {
	*x = RefreshServiceAclResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RefreshAuthorizationPolicyProtocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshServiceAclResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshServiceAclResponseProto) ProtoMessage() {}

func (x *RefreshServiceAclResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_RefreshAuthorizationPolicyProtocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshServiceAclResponseProto.ProtoReflect.Descriptor instead.
func (*RefreshServiceAclResponseProto) Descriptor() ([]byte, []int) {
	return file_RefreshAuthorizationPolicyProtocol_proto_rawDescGZIP(), []int{1}
}

var File_RefreshAuthorizationPolicyProtocol_proto protoreflect.FileDescriptor

var file_RefreshAuthorizationPolicyProtocol_proto_rawDesc = []byte{
	0x0a, 0x28, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x68, 0x61, 0x64, 0x6f,
	0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x1f, 0x0a, 0x1d, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x1e, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9d, 0x01, 0x0a,
	0x29, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x70, 0x0a, 0x11, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x6c, 0x12,
	0x2c, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x2e,
	0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x8f, 0x01, 0x0a,
	0x20, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x68, 0x61, 0x64, 0x6f,
	0x6f, 0x70, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x42, 0x28, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x6c, 0x69, 0x73,
	0x74, 0x61, 0x72, 0x2f, 0x68, 0x64, 0x66, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x68, 0x61, 0x64, 0x6f, 0x6f,
	0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0xa0, 0x01, 0x01,
}

var (
	file_RefreshAuthorizationPolicyProtocol_proto_rawDescOnce sync.Once
	file_RefreshAuthorizationPolicyProtocol_proto_rawDescData = file_RefreshAuthorizationPolicyProtocol_proto_rawDesc
)

func file_RefreshAuthorizationPolicyProtocol_proto_rawDescGZIP() []byte {
	file_RefreshAuthorizationPolicyProtocol_proto_rawDescOnce.Do(func() {
		file_RefreshAuthorizationPolicyProtocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_RefreshAuthorizationPolicyProtocol_proto_rawDescData)
	})
	return file_RefreshAuthorizationPolicyProtocol_proto_rawDescData
}

var file_RefreshAuthorizationPolicyProtocol_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_RefreshAuthorizationPolicyProtocol_proto_goTypes = []interface{}{
	(*RefreshServiceAclRequestProto)(nil),  // 0: hadoop.common.RefreshServiceAclRequestProto
	(*RefreshServiceAclResponseProto)(nil), // 1: hadoop.common.RefreshServiceAclResponseProto
}
var file_RefreshAuthorizationPolicyProtocol_proto_depIdxs = []int32{
	0, // 0: hadoop.common.RefreshAuthorizationPolicyProtocolService.refreshServiceAcl:input_type -> hadoop.common.RefreshServiceAclRequestProto
	1, // 1: hadoop.common.RefreshAuthorizationPolicyProtocolService.refreshServiceAcl:output_type -> hadoop.common.RefreshServiceAclResponseProto
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RefreshAuthorizationPolicyProtocol_proto_init() }
func file_RefreshAuthorizationPolicyProtocol_proto_init() {
	if File_RefreshAuthorizationPolicyProtocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RefreshAuthorizationPolicyProtocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshServiceAclRequestProto); i {
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
		file_RefreshAuthorizationPolicyProtocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshServiceAclResponseProto); i {
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
			RawDescriptor: file_RefreshAuthorizationPolicyProtocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_RefreshAuthorizationPolicyProtocol_proto_goTypes,
		DependencyIndexes: file_RefreshAuthorizationPolicyProtocol_proto_depIdxs,
		MessageInfos:      file_RefreshAuthorizationPolicyProtocol_proto_msgTypes,
	}.Build()
	File_RefreshAuthorizationPolicyProtocol_proto = out.File
	file_RefreshAuthorizationPolicyProtocol_proto_rawDesc = nil
	file_RefreshAuthorizationPolicyProtocol_proto_goTypes = nil
	file_RefreshAuthorizationPolicyProtocol_proto_depIdxs = nil
}
