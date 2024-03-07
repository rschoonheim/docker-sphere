// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: api/environment/environment.proto

package environment

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_environment_environment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_environment_environment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_environment_environment_proto_rawDescGZIP(), []int{0}
}

type GetEnvironmentHealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DockerStatus string `protobuf:"bytes,1,opt,name=docker_status,json=dockerStatus,proto3" json:"docker_status,omitempty"`
}

func (x *GetEnvironmentHealthResponse) Reset() {
	*x = GetEnvironmentHealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_environment_environment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEnvironmentHealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEnvironmentHealthResponse) ProtoMessage() {}

func (x *GetEnvironmentHealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_environment_environment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEnvironmentHealthResponse.ProtoReflect.Descriptor instead.
func (*GetEnvironmentHealthResponse) Descriptor() ([]byte, []int) {
	return file_api_environment_environment_proto_rawDescGZIP(), []int{1}
}

func (x *GetEnvironmentHealthResponse) GetDockerStatus() string {
	if x != nil {
		return x.DockerStatus
	}
	return ""
}

var File_api_environment_environment_proto protoreflect.FileDescriptor

var file_api_environment_environment_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x43, 0x0a, 0x1c,
	0x47, 0x65, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x32, 0x71, 0x0a, 0x12, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x45, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12,
	0x15, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_environment_environment_proto_rawDescOnce sync.Once
	file_api_environment_environment_proto_rawDescData = file_api_environment_environment_proto_rawDesc
)

func file_api_environment_environment_proto_rawDescGZIP() []byte {
	file_api_environment_environment_proto_rawDescOnce.Do(func() {
		file_api_environment_environment_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_environment_environment_proto_rawDescData)
	})
	return file_api_environment_environment_proto_rawDescData
}

var file_api_environment_environment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_environment_environment_proto_goTypes = []interface{}{
	(*Empty)(nil),                        // 0: environment.v1.Empty
	(*GetEnvironmentHealthResponse)(nil), // 1: environment.v1.GetEnvironmentHealthResponse
}
var file_api_environment_environment_proto_depIdxs = []int32{
	0, // 0: environment.v1.EnvironmentService.GetEnvironmentHealth:input_type -> environment.v1.Empty
	1, // 1: environment.v1.EnvironmentService.GetEnvironmentHealth:output_type -> environment.v1.GetEnvironmentHealthResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_environment_environment_proto_init() }
func file_api_environment_environment_proto_init() {
	if File_api_environment_environment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_environment_environment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_api_environment_environment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEnvironmentHealthResponse); i {
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
			RawDescriptor: file_api_environment_environment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_environment_environment_proto_goTypes,
		DependencyIndexes: file_api_environment_environment_proto_depIdxs,
		MessageInfos:      file_api_environment_environment_proto_msgTypes,
	}.Build()
	File_api_environment_environment_proto = out.File
	file_api_environment_environment_proto_rawDesc = nil
	file_api_environment_environment_proto_goTypes = nil
	file_api_environment_environment_proto_depIdxs = nil
}
