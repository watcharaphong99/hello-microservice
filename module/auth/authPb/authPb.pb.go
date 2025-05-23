// Version

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: module/auth/authPb/authPb.proto

package hello_microservice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Structures
type CredentialSearchReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CredentialSearchReq) Reset() {
	*x = CredentialSearchReq{}
	mi := &file_module_auth_authPb_authPb_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CredentialSearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CredentialSearchReq) ProtoMessage() {}

func (x *CredentialSearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_module_auth_authPb_authPb_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CredentialSearchReq.ProtoReflect.Descriptor instead.
func (*CredentialSearchReq) Descriptor() ([]byte, []int) {
	return file_module_auth_authPb_authPb_proto_rawDescGZIP(), []int{0}
}

func (x *CredentialSearchReq) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type CredentialSearchRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsValid       bool                   `protobuf:"varint,1,opt,name=isValid,proto3" json:"isValid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CredentialSearchRes) Reset() {
	*x = CredentialSearchRes{}
	mi := &file_module_auth_authPb_authPb_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CredentialSearchRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CredentialSearchRes) ProtoMessage() {}

func (x *CredentialSearchRes) ProtoReflect() protoreflect.Message {
	mi := &file_module_auth_authPb_authPb_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CredentialSearchRes.ProtoReflect.Descriptor instead.
func (*CredentialSearchRes) Descriptor() ([]byte, []int) {
	return file_module_auth_authPb_authPb_proto_rawDescGZIP(), []int{1}
}

func (x *CredentialSearchRes) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

type RolesCountReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RolesCountReq) Reset() {
	*x = RolesCountReq{}
	mi := &file_module_auth_authPb_authPb_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RolesCountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RolesCountReq) ProtoMessage() {}

func (x *RolesCountReq) ProtoReflect() protoreflect.Message {
	mi := &file_module_auth_authPb_authPb_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RolesCountReq.ProtoReflect.Descriptor instead.
func (*RolesCountReq) Descriptor() ([]byte, []int) {
	return file_module_auth_authPb_authPb_proto_rawDescGZIP(), []int{2}
}

type RolesCountRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Count         int64                  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RolesCountRes) Reset() {
	*x = RolesCountRes{}
	mi := &file_module_auth_authPb_authPb_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RolesCountRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RolesCountRes) ProtoMessage() {}

func (x *RolesCountRes) ProtoReflect() protoreflect.Message {
	mi := &file_module_auth_authPb_authPb_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RolesCountRes.ProtoReflect.Descriptor instead.
func (*RolesCountRes) Descriptor() ([]byte, []int) {
	return file_module_auth_authPb_authPb_proto_rawDescGZIP(), []int{3}
}

func (x *RolesCountRes) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_module_auth_authPb_authPb_proto protoreflect.FileDescriptor

const file_module_auth_authPb_authPb_proto_rawDesc = "" +
	"\n" +
	"\x1fmodule/auth/authPb/authPb.proto\"7\n" +
	"\x13CredentialSearchReq\x12 \n" +
	"\vaccessToken\x18\x01 \x01(\tR\vaccessToken\"/\n" +
	"\x13CredentialSearchRes\x12\x18\n" +
	"\aisValid\x18\x01 \x01(\bR\aisValid\"\x0f\n" +
	"\rRolesCountReq\"%\n" +
	"\rRolesCountRes\x12\x14\n" +
	"\x05count\x18\x01 \x01(\x03R\x05count2~\n" +
	"\x0fAuthGrpcService\x12>\n" +
	"\x10CredentialSearch\x12\x14.CredentialSearchReq\x1a\x14.CredentialSearchRes\x12+\n" +
	"\tRoleCount\x12\x0e.RolesCountReq\x1a\x0e.RolesCountResB)Z'github.com/bunkai888/hello-microserviceb\x06proto3"

var (
	file_module_auth_authPb_authPb_proto_rawDescOnce sync.Once
	file_module_auth_authPb_authPb_proto_rawDescData []byte
)

func file_module_auth_authPb_authPb_proto_rawDescGZIP() []byte {
	file_module_auth_authPb_authPb_proto_rawDescOnce.Do(func() {
		file_module_auth_authPb_authPb_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_module_auth_authPb_authPb_proto_rawDesc), len(file_module_auth_authPb_authPb_proto_rawDesc)))
	})
	return file_module_auth_authPb_authPb_proto_rawDescData
}

var file_module_auth_authPb_authPb_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_module_auth_authPb_authPb_proto_goTypes = []any{
	(*CredentialSearchReq)(nil), // 0: CredentialSearchReq
	(*CredentialSearchRes)(nil), // 1: CredentialSearchRes
	(*RolesCountReq)(nil),       // 2: RolesCountReq
	(*RolesCountRes)(nil),       // 3: RolesCountRes
}
var file_module_auth_authPb_authPb_proto_depIdxs = []int32{
	0, // 0: AuthGrpcService.CredentialSearch:input_type -> CredentialSearchReq
	2, // 1: AuthGrpcService.RoleCount:input_type -> RolesCountReq
	1, // 2: AuthGrpcService.CredentialSearch:output_type -> CredentialSearchRes
	3, // 3: AuthGrpcService.RoleCount:output_type -> RolesCountRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_module_auth_authPb_authPb_proto_init() }
func file_module_auth_authPb_authPb_proto_init() {
	if File_module_auth_authPb_authPb_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_module_auth_authPb_authPb_proto_rawDesc), len(file_module_auth_authPb_authPb_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_module_auth_authPb_authPb_proto_goTypes,
		DependencyIndexes: file_module_auth_authPb_authPb_proto_depIdxs,
		MessageInfos:      file_module_auth_authPb_authPb_proto_msgTypes,
	}.Build()
	File_module_auth_authPb_authPb_proto = out.File
	file_module_auth_authPb_authPb_proto_goTypes = nil
	file_module_auth_authPb_authPb_proto_depIdxs = nil
}
