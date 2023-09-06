// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: policy.proto

package protobuf

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type PolicyStatus int32

const (
	PolicyStatus_Failure  PolicyStatus = 0
	PolicyStatus_Applied  PolicyStatus = 1
	PolicyStatus_Deleted  PolicyStatus = 2
	PolicyStatus_Modified PolicyStatus = 3
	PolicyStatus_NotExist PolicyStatus = 4
	PolicyStatus_Invalid  PolicyStatus = 5
)

// Enum value maps for PolicyStatus.
var (
	PolicyStatus_name = map[int32]string{
		0: "Failure",
		1: "Applied",
		2: "Deleted",
		3: "Modified",
		4: "NotExist",
		5: "Invalid",
	}
	PolicyStatus_value = map[string]int32{
		"Failure":  0,
		"Applied":  1,
		"Deleted":  2,
		"Modified": 3,
		"NotExist": 4,
		"Invalid":  5,
	}
)

func (x PolicyStatus) Enum() *PolicyStatus {
	p := new(PolicyStatus)
	*p = x
	return p
}

func (x PolicyStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PolicyStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_policy_proto_enumTypes[0].Descriptor()
}

func (PolicyStatus) Type() protoreflect.EnumType {
	return &file_policy_proto_enumTypes[0]
}

func (x PolicyStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PolicyStatus.Descriptor instead.
func (PolicyStatus) EnumDescriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{0}
}

// Health check
type HealthCheckReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce int32 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *HealthCheckReq) Reset() {
	*x = HealthCheckReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckReq) ProtoMessage() {}

func (x *HealthCheckReq) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckReq.ProtoReflect.Descriptor instead.
func (*HealthCheckReq) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{0}
}

func (x *HealthCheckReq) GetNonce() int32 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

// healthcheck reply message
type HealthCheckReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retval int32 `protobuf:"varint,1,opt,name=Retval,proto3" json:"Retval,omitempty"`
}

func (x *HealthCheckReply) Reset() {
	*x = HealthCheckReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckReply) ProtoMessage() {}

func (x *HealthCheckReply) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[1]
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
	return file_policy_proto_rawDescGZIP(), []int{1}
}

func (x *HealthCheckReply) GetRetval() int32 {
	if x != nil {
		return x.Retval
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status PolicyStatus `protobuf:"varint,1,opt,name=status,proto3,enum=policy.PolicyStatus" json:"status,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetStatus() PolicyStatus {
	if x != nil {
		return x.Status
	}
	return PolicyStatus_Failure
}

type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy []byte `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{3}
}

func (x *Policy) GetPolicy() []byte {
	if x != nil {
		return x.Policy
	}
	return nil
}

type ContainerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyList    []string `protobuf:"bytes,1,rep,name=policyList,proto3" json:"policyList,omitempty"`
	PolicyEnabled int32    `protobuf:"varint,2,opt,name=policyEnabled,proto3" json:"policyEnabled,omitempty"`
}

func (x *ContainerData) Reset() {
	*x = ContainerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerData) ProtoMessage() {}

func (x *ContainerData) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerData.ProtoReflect.Descriptor instead.
func (*ContainerData) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{4}
}

func (x *ContainerData) GetPolicyList() []string {
	if x != nil {
		return x.PolicyList
	}
	return nil
}

func (x *ContainerData) GetPolicyEnabled() int32 {
	if x != nil {
		return x.PolicyEnabled
	}
	return 0
}

type HostSecurityPolicies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyList []string `protobuf:"bytes,1,rep,name=policyList,proto3" json:"policyList,omitempty"`
}

func (x *HostSecurityPolicies) Reset() {
	*x = HostSecurityPolicies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostSecurityPolicies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostSecurityPolicies) ProtoMessage() {}

func (x *HostSecurityPolicies) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostSecurityPolicies.ProtoReflect.Descriptor instead.
func (*HostSecurityPolicies) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{5}
}

func (x *HostSecurityPolicies) GetPolicyList() []string {
	if x != nil {
		return x.PolicyList
	}
	return nil
}

type Karmorresponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerList []string                         `protobuf:"bytes,1,rep,name=containerList,proto3" json:"containerList,omitempty"`
	ContainerMap  map[string]*ContainerData        `protobuf:"bytes,2,rep,name=containerMap,proto3" json:"containerMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	HostMap       map[string]*HostSecurityPolicies `protobuf:"bytes,3,rep,name=hostMap,proto3" json:"hostMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Karmorresponse) Reset() {
	*x = Karmorresponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Karmorresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Karmorresponse) ProtoMessage() {}

func (x *Karmorresponse) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Karmorresponse.ProtoReflect.Descriptor instead.
func (*Karmorresponse) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{6}
}

func (x *Karmorresponse) GetContainerList() []string {
	if x != nil {
		return x.ContainerList
	}
	return nil
}

func (x *Karmorresponse) GetContainerMap() map[string]*ContainerData {
	if x != nil {
		return x.ContainerMap
	}
	return nil
}

func (x *Karmorresponse) GetHostMap() map[string]*HostSecurityPolicies {
	if x != nil {
		return x.HostMap
	}
	return nil
}

var File_policy_proto protoreflect.FileDescriptor

var file_policy_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x2a, 0x0a, 0x10, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x52, 0x65, 0x74, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x52, 0x65, 0x74, 0x76, 0x61, 0x6c, 0x22, 0x38, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x20, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x22, 0x55, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x36, 0x0a, 0x14, 0x48, 0x6f,
	0x73, 0x74, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0xf5, 0x02, 0x0a, 0x0e, 0x6b, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x0c, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6b, 0x61, 0x72, 0x6d, 0x6f,
	0x72, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x12, 0x3d, 0x0a, 0x07, 0x68, 0x6f, 0x73,
	0x74, 0x4d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x6b, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x1a, 0x56, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x58, 0x0a, 0x0c, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x5e, 0x0a, 0x0c, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x64, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10,
	0x02, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x03, 0x12,
	0x0c, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0x04, 0x12, 0x0b, 0x0a,
	0x07, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x05, 0x32, 0x49, 0x0a, 0x06, 0x6b, 0x61,
	0x72, 0x6d, 0x6f, 0x72, 0x12, 0x3f, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x4b, 0x61, 0x72, 0x6d, 0x6f,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6b, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x74, 0x0a, 0x0d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0e, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x10, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x68,
	0x6f, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0e, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x10, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc3, 0x01, 0x0a, 0x13,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0x16, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x37, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x10, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x0e, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x28, 0x01, 0x30, 0x01, 0x12, 0x32, 0x0a,
	0x0a, 0x68, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x10, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x0e, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6b, 0x75, 0x62, 0x65, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x2f, 0x4b, 0x75, 0x62, 0x65, 0x41, 0x72,
	0x6d, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_policy_proto_rawDescOnce sync.Once
	file_policy_proto_rawDescData = file_policy_proto_rawDesc
)

func file_policy_proto_rawDescGZIP() []byte {
	file_policy_proto_rawDescOnce.Do(func() {
		file_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_policy_proto_rawDescData)
	})
	return file_policy_proto_rawDescData
}

var file_policy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_policy_proto_goTypes = []interface{}{
	(PolicyStatus)(0),            // 0: policy.PolicyStatus
	(*HealthCheckReq)(nil),       // 1: policy.HealthCheckReq
	(*HealthCheckReply)(nil),     // 2: policy.HealthCheckReply
	(*Response)(nil),             // 3: policy.response
	(*Policy)(nil),               // 4: policy.policy
	(*ContainerData)(nil),        // 5: policy.ContainerData
	(*HostSecurityPolicies)(nil), // 6: policy.HostSecurityPolicies
	(*Karmorresponse)(nil),       // 7: policy.karmorresponse
	nil,                          // 8: policy.karmorresponse.ContainerMapEntry
	nil,                          // 9: policy.karmorresponse.HostMapEntry
	(*empty.Empty)(nil),          // 10: google.protobuf.Empty
}
var file_policy_proto_depIdxs = []int32{
	0,  // 0: policy.response.status:type_name -> policy.PolicyStatus
	8,  // 1: policy.karmorresponse.containerMap:type_name -> policy.karmorresponse.ContainerMapEntry
	9,  // 2: policy.karmorresponse.hostMap:type_name -> policy.karmorresponse.HostMapEntry
	5,  // 3: policy.karmorresponse.ContainerMapEntry.value:type_name -> policy.ContainerData
	6,  // 4: policy.karmorresponse.HostMapEntry.value:type_name -> policy.HostSecurityPolicies
	10, // 5: policy.karmor.getKarmorData:input_type -> google.protobuf.Empty
	4,  // 6: policy.PolicyService.containerPolicy:input_type -> policy.policy
	4,  // 7: policy.PolicyService.hostPolicy:input_type -> policy.policy
	1,  // 8: policy.PolicyStreamService.HealthCheck:input_type -> policy.HealthCheckReq
	3,  // 9: policy.PolicyStreamService.containerPolicy:input_type -> policy.response
	3,  // 10: policy.PolicyStreamService.hostPolicy:input_type -> policy.response
	7,  // 11: policy.karmor.getKarmorData:output_type -> policy.karmorresponse
	3,  // 12: policy.PolicyService.containerPolicy:output_type -> policy.response
	3,  // 13: policy.PolicyService.hostPolicy:output_type -> policy.response
	2,  // 14: policy.PolicyStreamService.HealthCheck:output_type -> policy.HealthCheckReply
	4,  // 15: policy.PolicyStreamService.containerPolicy:output_type -> policy.policy
	4,  // 16: policy.PolicyStreamService.hostPolicy:output_type -> policy.policy
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_policy_proto_init() }
func file_policy_proto_init() {
	if File_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckReq); i {
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
		file_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerData); i {
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
		file_policy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostSecurityPolicies); i {
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
		file_policy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Karmorresponse); i {
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
			RawDescriptor: file_policy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_policy_proto_goTypes,
		DependencyIndexes: file_policy_proto_depIdxs,
		EnumInfos:         file_policy_proto_enumTypes,
		MessageInfos:      file_policy_proto_msgTypes,
	}.Build()
	File_policy_proto = out.File
	file_policy_proto_rawDesc = nil
	file_policy_proto_goTypes = nil
	file_policy_proto_depIdxs = nil
}
