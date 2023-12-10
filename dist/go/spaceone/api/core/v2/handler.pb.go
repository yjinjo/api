// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: spaceone/api/core/v2/handler.proto

package v2

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type AuthorizationRequest_OwnerType int32

const (
	AuthorizationRequest_NONE AuthorizationRequest_OwnerType = 0
	AuthorizationRequest_USER AuthorizationRequest_OwnerType = 1
	AuthorizationRequest_APP  AuthorizationRequest_OwnerType = 2
)

// Enum value maps for AuthorizationRequest_OwnerType.
var (
	AuthorizationRequest_OwnerType_name = map[int32]string{
		0: "NONE",
		1: "USER",
		2: "APP",
	}
	AuthorizationRequest_OwnerType_value = map[string]int32{
		"NONE": 0,
		"USER": 1,
		"APP":  2,
	}
)

func (x AuthorizationRequest_OwnerType) Enum() *AuthorizationRequest_OwnerType {
	p := new(AuthorizationRequest_OwnerType)
	*p = x
	return p
}

func (x AuthorizationRequest_OwnerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthorizationRequest_OwnerType) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_core_v2_handler_proto_enumTypes[0].Descriptor()
}

func (AuthorizationRequest_OwnerType) Type() protoreflect.EnumType {
	return &file_spaceone_api_core_v2_handler_proto_enumTypes[0]
}

func (x AuthorizationRequest_OwnerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthorizationRequest_OwnerType.Descriptor instead.
func (AuthorizationRequest_OwnerType) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_core_v2_handler_proto_rawDescGZIP(), []int{0, 0}
}

type AuthorizationResponse_RoleType int32

const (
	AuthorizationResponse_NONE             AuthorizationResponse_RoleType = 0
	AuthorizationResponse_SYSTEM           AuthorizationResponse_RoleType = 1
	AuthorizationResponse_SYSTEM_ADMIN     AuthorizationResponse_RoleType = 2
	AuthorizationResponse_DOMAIN_ADMIN     AuthorizationResponse_RoleType = 3
	AuthorizationResponse_WORKSPACE_OWNER  AuthorizationResponse_RoleType = 4
	AuthorizationResponse_WORKSPACE_MEMBER AuthorizationResponse_RoleType = 5
	AuthorizationResponse_USER             AuthorizationResponse_RoleType = 6
)

// Enum value maps for AuthorizationResponse_RoleType.
var (
	AuthorizationResponse_RoleType_name = map[int32]string{
		0: "NONE",
		1: "SYSTEM",
		2: "SYSTEM_ADMIN",
		3: "DOMAIN_ADMIN",
		4: "WORKSPACE_OWNER",
		5: "WORKSPACE_MEMBER",
		6: "USER",
	}
	AuthorizationResponse_RoleType_value = map[string]int32{
		"NONE":             0,
		"SYSTEM":           1,
		"SYSTEM_ADMIN":     2,
		"DOMAIN_ADMIN":     3,
		"WORKSPACE_OWNER":  4,
		"WORKSPACE_MEMBER": 5,
		"USER":             6,
	}
)

func (x AuthorizationResponse_RoleType) Enum() *AuthorizationResponse_RoleType {
	p := new(AuthorizationResponse_RoleType)
	*p = x
	return p
}

func (x AuthorizationResponse_RoleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthorizationResponse_RoleType) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_core_v2_handler_proto_enumTypes[1].Descriptor()
}

func (AuthorizationResponse_RoleType) Type() protoreflect.EnumType {
	return &file_spaceone_api_core_v2_handler_proto_enumTypes[1]
}

func (x AuthorizationResponse_RoleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthorizationResponse_RoleType.Descriptor instead.
func (AuthorizationResponse_RoleType) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_core_v2_handler_proto_rawDescGZIP(), []int{1, 0}
}

type AuthorizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scope     string                         `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope,omitempty"`
	OwnerType AuthorizationRequest_OwnerType `protobuf:"varint,2,opt,name=owner_type,json=ownerType,proto3,enum=spaceone.api.core.v2.AuthorizationRequest_OwnerType" json:"owner_type,omitempty"`
	Audience  string                         `protobuf:"bytes,3,opt,name=audience,proto3" json:"audience,omitempty"`
	TokenId   string                         `protobuf:"bytes,4,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	// +optional
	WorkspaceId string `protobuf:"bytes,6,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// +optional
	DomainId string `protobuf:"bytes,7,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *AuthorizationRequest) Reset() {
	*x = AuthorizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_core_v2_handler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationRequest) ProtoMessage() {}

func (x *AuthorizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_core_v2_handler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationRequest.ProtoReflect.Descriptor instead.
func (*AuthorizationRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_core_v2_handler_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorizationRequest) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *AuthorizationRequest) GetOwnerType() AuthorizationRequest_OwnerType {
	if x != nil {
		return x.OwnerType
	}
	return AuthorizationRequest_NONE
}

func (x *AuthorizationRequest) GetAudience() string {
	if x != nil {
		return x.Audience
	}
	return ""
}

func (x *AuthorizationRequest) GetTokenId() string {
	if x != nil {
		return x.TokenId
	}
	return ""
}

func (x *AuthorizationRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *AuthorizationRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type AuthorizationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleType     AuthorizationResponse_RoleType `protobuf:"varint,1,opt,name=role_type,json=roleType,proto3,enum=spaceone.api.core.v2.AuthorizationResponse_RoleType" json:"role_type,omitempty"`
	UserProjects []string                       `protobuf:"bytes,2,rep,name=user_projects,json=userProjects,proto3" json:"user_projects,omitempty"`
}

func (x *AuthorizationResponse) Reset() {
	*x = AuthorizationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_core_v2_handler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationResponse) ProtoMessage() {}

func (x *AuthorizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_core_v2_handler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationResponse.ProtoReflect.Descriptor instead.
func (*AuthorizationResponse) Descriptor() ([]byte, []int) {
	return file_spaceone_api_core_v2_handler_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorizationResponse) GetRoleType() AuthorizationResponse_RoleType {
	if x != nil {
		return x.RoleType
	}
	return AuthorizationResponse_NONE
}

func (x *AuthorizationResponse) GetUserProjects() []string {
	if x != nil {
		return x.UserProjects
	}
	return nil
}

type AuthenticationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainId string `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *AuthenticationRequest) Reset() {
	*x = AuthenticationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_core_v2_handler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationRequest) ProtoMessage() {}

func (x *AuthenticationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_core_v2_handler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationRequest.ProtoReflect.Descriptor instead.
func (*AuthenticationRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_core_v2_handler_proto_rawDescGZIP(), []int{2}
}

func (x *AuthenticationRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type AuthenticationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainId  string `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	PublicKey string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *AuthenticationResponse) Reset() {
	*x = AuthenticationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_core_v2_handler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationResponse) ProtoMessage() {}

func (x *AuthenticationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_core_v2_handler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationResponse.ProtoReflect.Descriptor instead.
func (*AuthenticationResponse) Descriptor() ([]byte, []int) {
	return file_spaceone_api_core_v2_handler_proto_rawDescGZIP(), []int{3}
}

func (x *AuthenticationResponse) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *AuthenticationResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type EventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service  string          `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Resource string          `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	Verb     string          `protobuf:"bytes,3,opt,name=verb,proto3" json:"verb,omitempty"`
	Status   string          `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Message  *_struct.Struct `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EventRequest) Reset() {
	*x = EventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_core_v2_handler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventRequest) ProtoMessage() {}

func (x *EventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_core_v2_handler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventRequest.ProtoReflect.Descriptor instead.
func (*EventRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_core_v2_handler_proto_rawDescGZIP(), []int{4}
}

func (x *EventRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *EventRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *EventRequest) GetVerb() string {
	if x != nil {
		return x.Verb
	}
	return ""
}

func (x *EventRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *EventRequest) GetMessage() *_struct.Struct {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_spaceone_api_core_v2_handler_proto protoreflect.FileDescriptor

var file_spaceone_api_core_v2_handler_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x02, 0x0a, 0x14, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x53, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x09, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53,
	0x45, 0x52, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x50, 0x50, 0x10, 0x02, 0x22, 0x8a, 0x02,
	0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22,
	0x79, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10,
	0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x41, 0x44, 0x4d, 0x49,
	0x4e, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x41, 0x44,
	0x4d, 0x49, 0x4e, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41,
	0x43, 0x45, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x57, 0x4f,
	0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x05,
	0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x06, 0x22, 0x34, 0x0a, 0x15, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x22, 0x54, 0x0a, 0x16, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0xa3, 0x01, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x76, 0x65, 0x72, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x76, 0x65, 0x72,
	0x62, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x3b, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73,
	0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_spaceone_api_core_v2_handler_proto_rawDescOnce sync.Once
	file_spaceone_api_core_v2_handler_proto_rawDescData = file_spaceone_api_core_v2_handler_proto_rawDesc
)

func file_spaceone_api_core_v2_handler_proto_rawDescGZIP() []byte {
	file_spaceone_api_core_v2_handler_proto_rawDescOnce.Do(func() {
		file_spaceone_api_core_v2_handler_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_core_v2_handler_proto_rawDescData)
	})
	return file_spaceone_api_core_v2_handler_proto_rawDescData
}

var file_spaceone_api_core_v2_handler_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_spaceone_api_core_v2_handler_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spaceone_api_core_v2_handler_proto_goTypes = []interface{}{
	(AuthorizationRequest_OwnerType)(0), // 0: spaceone.api.core.v2.AuthorizationRequest.OwnerType
	(AuthorizationResponse_RoleType)(0), // 1: spaceone.api.core.v2.AuthorizationResponse.RoleType
	(*AuthorizationRequest)(nil),        // 2: spaceone.api.core.v2.AuthorizationRequest
	(*AuthorizationResponse)(nil),       // 3: spaceone.api.core.v2.AuthorizationResponse
	(*AuthenticationRequest)(nil),       // 4: spaceone.api.core.v2.AuthenticationRequest
	(*AuthenticationResponse)(nil),      // 5: spaceone.api.core.v2.AuthenticationResponse
	(*EventRequest)(nil),                // 6: spaceone.api.core.v2.EventRequest
	(*_struct.Struct)(nil),              // 7: google.protobuf.Struct
}
var file_spaceone_api_core_v2_handler_proto_depIdxs = []int32{
	0, // 0: spaceone.api.core.v2.AuthorizationRequest.owner_type:type_name -> spaceone.api.core.v2.AuthorizationRequest.OwnerType
	1, // 1: spaceone.api.core.v2.AuthorizationResponse.role_type:type_name -> spaceone.api.core.v2.AuthorizationResponse.RoleType
	7, // 2: spaceone.api.core.v2.EventRequest.message:type_name -> google.protobuf.Struct
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_spaceone_api_core_v2_handler_proto_init() }
func file_spaceone_api_core_v2_handler_proto_init() {
	if File_spaceone_api_core_v2_handler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_core_v2_handler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationRequest); i {
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
		file_spaceone_api_core_v2_handler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationResponse); i {
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
		file_spaceone_api_core_v2_handler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationRequest); i {
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
		file_spaceone_api_core_v2_handler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationResponse); i {
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
		file_spaceone_api_core_v2_handler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventRequest); i {
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
			RawDescriptor: file_spaceone_api_core_v2_handler_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spaceone_api_core_v2_handler_proto_goTypes,
		DependencyIndexes: file_spaceone_api_core_v2_handler_proto_depIdxs,
		EnumInfos:         file_spaceone_api_core_v2_handler_proto_enumTypes,
		MessageInfos:      file_spaceone_api_core_v2_handler_proto_msgTypes,
	}.Build()
	File_spaceone_api_core_v2_handler_proto = out.File
	file_spaceone_api_core_v2_handler_proto_rawDesc = nil
	file_spaceone_api_core_v2_handler_proto_goTypes = nil
	file_spaceone_api_core_v2_handler_proto_depIdxs = nil
}
