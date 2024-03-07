// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/deployment/deployment.proto

package deployment

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

type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status int64  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_deployment_deployment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_api_deployment_deployment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_api_deployment_deployment_proto_rawDescGZIP(), []int{0}
}

func (x *Deployment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Deployment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Deployment) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type CreateDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateDeploymentRequest) Reset() {
	*x = CreateDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_deployment_deployment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeploymentRequest) ProtoMessage() {}

func (x *CreateDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_deployment_deployment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeploymentRequest.ProtoReflect.Descriptor instead.
func (*CreateDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_api_deployment_deployment_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDeploymentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateDeploymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateDeploymentResponse) Reset() {
	*x = CreateDeploymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_deployment_deployment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeploymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeploymentResponse) ProtoMessage() {}

func (x *CreateDeploymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_deployment_deployment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeploymentResponse.ProtoReflect.Descriptor instead.
func (*CreateDeploymentResponse) Descriptor() ([]byte, []int) {
	return file_api_deployment_deployment_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDeploymentResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDeploymentRequest) Reset() {
	*x = GetDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_deployment_deployment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeploymentRequest) ProtoMessage() {}

func (x *GetDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_deployment_deployment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeploymentRequest.ProtoReflect.Descriptor instead.
func (*GetDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_api_deployment_deployment_proto_rawDescGZIP(), []int{3}
}

func (x *GetDeploymentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetDeploymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status int64  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetDeploymentResponse) Reset() {
	*x = GetDeploymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_deployment_deployment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeploymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeploymentResponse) ProtoMessage() {}

func (x *GetDeploymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_deployment_deployment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeploymentResponse.ProtoReflect.Descriptor instead.
func (*GetDeploymentResponse) Descriptor() ([]byte, []int) {
	return file_api_deployment_deployment_proto_rawDescGZIP(), []int{4}
}

func (x *GetDeploymentResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetDeploymentResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetDeploymentResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ListDeploymentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListDeploymentsRequest) Reset() {
	*x = ListDeploymentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_deployment_deployment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeploymentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeploymentsRequest) ProtoMessage() {}

func (x *ListDeploymentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_deployment_deployment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeploymentsRequest.ProtoReflect.Descriptor instead.
func (*ListDeploymentsRequest) Descriptor() ([]byte, []int) {
	return file_api_deployment_deployment_proto_rawDescGZIP(), []int{5}
}

func (x *ListDeploymentsRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListDeploymentsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListDeploymentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deployments []*Deployment `protobuf:"bytes,1,rep,name=deployments,proto3" json:"deployments,omitempty"`
}

func (x *ListDeploymentsResponse) Reset() {
	*x = ListDeploymentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_deployment_deployment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeploymentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeploymentsResponse) ProtoMessage() {}

func (x *ListDeploymentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_deployment_deployment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeploymentsResponse.ProtoReflect.Descriptor instead.
func (*ListDeploymentsResponse) Descriptor() ([]byte, []int) {
	return file_api_deployment_deployment_proto_rawDescGZIP(), []int{6}
}

func (x *ListDeploymentsResponse) GetDeployments() []*Deployment {
	if x != nil {
		return x.Deployments
	}
	return nil
}

var File_api_deployment_deployment_proto protoreflect.FileDescriptor

var file_api_deployment_deployment_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x22, 0x48, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2d, 0x0a, 0x17, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x18, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x53, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x46, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x56, 0x0a, 0x17, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x32, 0xb6, 0x02, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x0f, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x73, 0x63, 0x68, 0x6f, 0x6f,
	0x6e, 0x68, 0x65, 0x69, 0x6d, 0x2f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2d, 0x73, 0x70, 0x68,
	0x65, 0x72, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_deployment_deployment_proto_rawDescOnce sync.Once
	file_api_deployment_deployment_proto_rawDescData = file_api_deployment_deployment_proto_rawDesc
)

func file_api_deployment_deployment_proto_rawDescGZIP() []byte {
	file_api_deployment_deployment_proto_rawDescOnce.Do(func() {
		file_api_deployment_deployment_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_deployment_deployment_proto_rawDescData)
	})
	return file_api_deployment_deployment_proto_rawDescData
}

var file_api_deployment_deployment_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_deployment_deployment_proto_goTypes = []interface{}{
	(*Deployment)(nil),               // 0: deployment.v1.Deployment
	(*CreateDeploymentRequest)(nil),  // 1: deployment.v1.CreateDeploymentRequest
	(*CreateDeploymentResponse)(nil), // 2: deployment.v1.CreateDeploymentResponse
	(*GetDeploymentRequest)(nil),     // 3: deployment.v1.GetDeploymentRequest
	(*GetDeploymentResponse)(nil),    // 4: deployment.v1.GetDeploymentResponse
	(*ListDeploymentsRequest)(nil),   // 5: deployment.v1.ListDeploymentsRequest
	(*ListDeploymentsResponse)(nil),  // 6: deployment.v1.ListDeploymentsResponse
}
var file_api_deployment_deployment_proto_depIdxs = []int32{
	0, // 0: deployment.v1.ListDeploymentsResponse.deployments:type_name -> deployment.v1.Deployment
	1, // 1: deployment.v1.DeploymentService.CreateDeployment:input_type -> deployment.v1.CreateDeploymentRequest
	3, // 2: deployment.v1.DeploymentService.GetDeployment:input_type -> deployment.v1.GetDeploymentRequest
	5, // 3: deployment.v1.DeploymentService.ListDeployments:input_type -> deployment.v1.ListDeploymentsRequest
	2, // 4: deployment.v1.DeploymentService.CreateDeployment:output_type -> deployment.v1.CreateDeploymentResponse
	4, // 5: deployment.v1.DeploymentService.GetDeployment:output_type -> deployment.v1.GetDeploymentResponse
	6, // 6: deployment.v1.DeploymentService.ListDeployments:output_type -> deployment.v1.ListDeploymentsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_deployment_deployment_proto_init() }
func file_api_deployment_deployment_proto_init() {
	if File_api_deployment_deployment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_deployment_deployment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployment); i {
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
		file_api_deployment_deployment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeploymentRequest); i {
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
		file_api_deployment_deployment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeploymentResponse); i {
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
		file_api_deployment_deployment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeploymentRequest); i {
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
		file_api_deployment_deployment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeploymentResponse); i {
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
		file_api_deployment_deployment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeploymentsRequest); i {
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
		file_api_deployment_deployment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeploymentsResponse); i {
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
			RawDescriptor: file_api_deployment_deployment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_deployment_deployment_proto_goTypes,
		DependencyIndexes: file_api_deployment_deployment_proto_depIdxs,
		MessageInfos:      file_api_deployment_deployment_proto_msgTypes,
	}.Build()
	File_api_deployment_deployment_proto = out.File
	file_api_deployment_deployment_proto_rawDesc = nil
	file_api_deployment_deployment_proto_goTypes = nil
	file_api_deployment_deployment_proto_depIdxs = nil
}