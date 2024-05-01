// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v3.6.1
// source: spaceone/api/inventory/v1/metric_example.proto

package v1

import (
	v2 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// {
//
// }
type CreateMetricExampleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricId string          `protobuf:"bytes,1,opt,name=metric_id,json=metricId,proto3" json:"metric_id,omitempty"`
	Name     string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Options  *_struct.Struct `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *CreateMetricExampleRequest) Reset() {
	*x = CreateMetricExampleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMetricExampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMetricExampleRequest) ProtoMessage() {}

func (x *CreateMetricExampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMetricExampleRequest.ProtoReflect.Descriptor instead.
func (*CreateMetricExampleRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_metric_example_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMetricExampleRequest) GetMetricId() string {
	if x != nil {
		return x.MetricId
	}
	return ""
}

func (x *CreateMetricExampleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateMetricExampleRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *CreateMetricExampleRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

// {
//
// }
type UpdateMetricExampleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExampleId string `protobuf:"bytes,1,opt,name=example_id,json=exampleId,proto3" json:"example_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Options *_struct.Struct `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *UpdateMetricExampleRequest) Reset() {
	*x = UpdateMetricExampleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMetricExampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMetricExampleRequest) ProtoMessage() {}

func (x *UpdateMetricExampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMetricExampleRequest.ProtoReflect.Descriptor instead.
func (*UpdateMetricExampleRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_metric_example_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateMetricExampleRequest) GetExampleId() string {
	if x != nil {
		return x.ExampleId
	}
	return ""
}

func (x *UpdateMetricExampleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateMetricExampleRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *UpdateMetricExampleRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

// {
//
// }
type MetricExampleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExampleId string `protobuf:"bytes,1,opt,name=example_id,json=exampleId,proto3" json:"example_id,omitempty"`
}

func (x *MetricExampleRequest) Reset() {
	*x = MetricExampleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricExampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricExampleRequest) ProtoMessage() {}

func (x *MetricExampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricExampleRequest.ProtoReflect.Descriptor instead.
func (*MetricExampleRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_metric_example_proto_rawDescGZIP(), []int{2}
}

func (x *MetricExampleRequest) GetExampleId() string {
	if x != nil {
		return x.ExampleId
	}
	return ""
}

// {
//
// }
type MetricExampleQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	ExampleId string `protobuf:"bytes,2,opt,name=example_id,json=exampleId,proto3" json:"example_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	MetricId string `protobuf:"bytes,4,opt,name=metric_id,json=metricId,proto3" json:"metric_id,omitempty"`
	// +optional
	NamespaceId string `protobuf:"bytes,5,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
}

func (x *MetricExampleQuery) Reset() {
	*x = MetricExampleQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricExampleQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricExampleQuery) ProtoMessage() {}

func (x *MetricExampleQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricExampleQuery.ProtoReflect.Descriptor instead.
func (*MetricExampleQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_metric_example_proto_rawDescGZIP(), []int{3}
}

func (x *MetricExampleQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *MetricExampleQuery) GetExampleId() string {
	if x != nil {
		return x.ExampleId
	}
	return ""
}

func (x *MetricExampleQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetricExampleQuery) GetMetricId() string {
	if x != nil {
		return x.MetricId
	}
	return ""
}

func (x *MetricExampleQuery) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

// {
//
// }
type MetricExampleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExampleId   string          `protobuf:"bytes,1,opt,name=example_id,json=exampleId,proto3" json:"example_id,omitempty"`
	Name        string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Options     *_struct.Struct `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	Tags        *_struct.Struct `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	DomainId    string          `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId string          `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	UserId      string          `protobuf:"bytes,23,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MetricId    string          `protobuf:"bytes,24,opt,name=metric_id,json=metricId,proto3" json:"metric_id,omitempty"`
	NamespaceId string          `protobuf:"bytes,25,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	CreatedAt   string          `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string          `protobuf:"bytes,32,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *MetricExampleInfo) Reset() {
	*x = MetricExampleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricExampleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricExampleInfo) ProtoMessage() {}

func (x *MetricExampleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricExampleInfo.ProtoReflect.Descriptor instead.
func (*MetricExampleInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_metric_example_proto_rawDescGZIP(), []int{4}
}

func (x *MetricExampleInfo) GetExampleId() string {
	if x != nil {
		return x.ExampleId
	}
	return ""
}

func (x *MetricExampleInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetricExampleInfo) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *MetricExampleInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *MetricExampleInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *MetricExampleInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *MetricExampleInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *MetricExampleInfo) GetMetricId() string {
	if x != nil {
		return x.MetricId
	}
	return ""
}

func (x *MetricExampleInfo) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *MetricExampleInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *MetricExampleInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type MetricExamplesInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results    []*MetricExampleInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount int32                `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *MetricExamplesInfo) Reset() {
	*x = MetricExamplesInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricExamplesInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricExamplesInfo) ProtoMessage() {}

func (x *MetricExamplesInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricExamplesInfo.ProtoReflect.Descriptor instead.
func (*MetricExamplesInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_metric_example_proto_rawDescGZIP(), []int{5}
}

func (x *MetricExamplesInfo) GetResults() []*MetricExampleInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *MetricExamplesInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type MetricExampleStatQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query *v2.StatisticsQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *MetricExampleStatQuery) Reset() {
	*x = MetricExampleStatQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricExampleStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricExampleStatQuery) ProtoMessage() {}

func (x *MetricExampleStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricExampleStatQuery.ProtoReflect.Descriptor instead.
func (*MetricExampleStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_metric_example_proto_rawDescGZIP(), []int{6}
}

func (x *MetricExampleStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_spaceone_api_inventory_v1_metric_example_proto protoreflect.FileDescriptor

var file_spaceone_api_inventory_v1_metric_example_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0xaf, 0x01, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x35, 0x0a, 0x14, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x64, 0x22,
	0xba, 0x01, 0x0a, 0x12, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0xfd, 0x02, 0x0a,
	0x11, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7d, 0x0a, 0x12,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x46, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x55, 0x0a, 0x16, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x32, 0xff, 0x06, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x35, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22,
	0x23, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x35, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22,
	0x23, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28,
	0x3a, 0x01, 0x2a, 0x22, 0x23, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x91, 0x01, 0x0a, 0x03, 0x67, 0x65, 0x74,
	0x12, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x92, 0x01, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x1a, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a, 0x22, 0x21,
	0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x80, 0x01, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x12, 0x31, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01,
	0x2a, 0x22, 0x21, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_inventory_v1_metric_example_proto_rawDescOnce sync.Once
	file_spaceone_api_inventory_v1_metric_example_proto_rawDescData = file_spaceone_api_inventory_v1_metric_example_proto_rawDesc
)

func file_spaceone_api_inventory_v1_metric_example_proto_rawDescGZIP() []byte {
	file_spaceone_api_inventory_v1_metric_example_proto_rawDescOnce.Do(func() {
		file_spaceone_api_inventory_v1_metric_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_inventory_v1_metric_example_proto_rawDescData)
	})
	return file_spaceone_api_inventory_v1_metric_example_proto_rawDescData
}

var file_spaceone_api_inventory_v1_metric_example_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_spaceone_api_inventory_v1_metric_example_proto_goTypes = []interface{}{
	(*CreateMetricExampleRequest)(nil), // 0: spaceone.api.inventory.v1.CreateMetricExampleRequest
	(*UpdateMetricExampleRequest)(nil), // 1: spaceone.api.inventory.v1.UpdateMetricExampleRequest
	(*MetricExampleRequest)(nil),       // 2: spaceone.api.inventory.v1.MetricExampleRequest
	(*MetricExampleQuery)(nil),         // 3: spaceone.api.inventory.v1.MetricExampleQuery
	(*MetricExampleInfo)(nil),          // 4: spaceone.api.inventory.v1.MetricExampleInfo
	(*MetricExamplesInfo)(nil),         // 5: spaceone.api.inventory.v1.MetricExamplesInfo
	(*MetricExampleStatQuery)(nil),     // 6: spaceone.api.inventory.v1.MetricExampleStatQuery
	(*_struct.Struct)(nil),             // 7: google.protobuf.Struct
	(*v2.Query)(nil),                   // 8: spaceone.api.core.v2.Query
	(*v2.StatisticsQuery)(nil),         // 9: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),                // 10: google.protobuf.Empty
}
var file_spaceone_api_inventory_v1_metric_example_proto_depIdxs = []int32{
	7,  // 0: spaceone.api.inventory.v1.CreateMetricExampleRequest.options:type_name -> google.protobuf.Struct
	7,  // 1: spaceone.api.inventory.v1.CreateMetricExampleRequest.tags:type_name -> google.protobuf.Struct
	7,  // 2: spaceone.api.inventory.v1.UpdateMetricExampleRequest.options:type_name -> google.protobuf.Struct
	7,  // 3: spaceone.api.inventory.v1.UpdateMetricExampleRequest.tags:type_name -> google.protobuf.Struct
	8,  // 4: spaceone.api.inventory.v1.MetricExampleQuery.query:type_name -> spaceone.api.core.v2.Query
	7,  // 5: spaceone.api.inventory.v1.MetricExampleInfo.options:type_name -> google.protobuf.Struct
	7,  // 6: spaceone.api.inventory.v1.MetricExampleInfo.tags:type_name -> google.protobuf.Struct
	4,  // 7: spaceone.api.inventory.v1.MetricExamplesInfo.results:type_name -> spaceone.api.inventory.v1.MetricExampleInfo
	9,  // 8: spaceone.api.inventory.v1.MetricExampleStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	0,  // 9: spaceone.api.inventory.v1.MetricExample.create:input_type -> spaceone.api.inventory.v1.CreateMetricExampleRequest
	1,  // 10: spaceone.api.inventory.v1.MetricExample.update:input_type -> spaceone.api.inventory.v1.UpdateMetricExampleRequest
	2,  // 11: spaceone.api.inventory.v1.MetricExample.delete:input_type -> spaceone.api.inventory.v1.MetricExampleRequest
	2,  // 12: spaceone.api.inventory.v1.MetricExample.get:input_type -> spaceone.api.inventory.v1.MetricExampleRequest
	3,  // 13: spaceone.api.inventory.v1.MetricExample.list:input_type -> spaceone.api.inventory.v1.MetricExampleQuery
	6,  // 14: spaceone.api.inventory.v1.MetricExample.stat:input_type -> spaceone.api.inventory.v1.MetricExampleStatQuery
	4,  // 15: spaceone.api.inventory.v1.MetricExample.create:output_type -> spaceone.api.inventory.v1.MetricExampleInfo
	4,  // 16: spaceone.api.inventory.v1.MetricExample.update:output_type -> spaceone.api.inventory.v1.MetricExampleInfo
	10, // 17: spaceone.api.inventory.v1.MetricExample.delete:output_type -> google.protobuf.Empty
	4,  // 18: spaceone.api.inventory.v1.MetricExample.get:output_type -> spaceone.api.inventory.v1.MetricExampleInfo
	5,  // 19: spaceone.api.inventory.v1.MetricExample.list:output_type -> spaceone.api.inventory.v1.MetricExamplesInfo
	7,  // 20: spaceone.api.inventory.v1.MetricExample.stat:output_type -> google.protobuf.Struct
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_spaceone_api_inventory_v1_metric_example_proto_init() }
func file_spaceone_api_inventory_v1_metric_example_proto_init() {
	if File_spaceone_api_inventory_v1_metric_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMetricExampleRequest); i {
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
		file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMetricExampleRequest); i {
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
		file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricExampleRequest); i {
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
		file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricExampleQuery); i {
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
		file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricExampleInfo); i {
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
		file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricExamplesInfo); i {
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
		file_spaceone_api_inventory_v1_metric_example_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricExampleStatQuery); i {
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
			RawDescriptor: file_spaceone_api_inventory_v1_metric_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_inventory_v1_metric_example_proto_goTypes,
		DependencyIndexes: file_spaceone_api_inventory_v1_metric_example_proto_depIdxs,
		MessageInfos:      file_spaceone_api_inventory_v1_metric_example_proto_msgTypes,
	}.Build()
	File_spaceone_api_inventory_v1_metric_example_proto = out.File
	file_spaceone_api_inventory_v1_metric_example_proto_rawDesc = nil
	file_spaceone_api_inventory_v1_metric_example_proto_goTypes = nil
	file_spaceone_api_inventory_v1_metric_example_proto_depIdxs = nil
}
