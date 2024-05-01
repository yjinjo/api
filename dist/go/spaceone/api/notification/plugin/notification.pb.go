// A Notification is a resource delivering data from a Protocol plugin instance to a nofitication tool of an external user.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v3.6.1
// source: spaceone/api/notification/plugin/notification.proto

package plugin

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type PluginDispatchRequest_NotificationType int32

const (
	PluginDispatchRequest_NOTIFICATION_TYPE_NONE PluginDispatchRequest_NotificationType = 0
	PluginDispatchRequest_INFO                   PluginDispatchRequest_NotificationType = 1
	PluginDispatchRequest_ERROR                  PluginDispatchRequest_NotificationType = 2
	PluginDispatchRequest_SUCCESS                PluginDispatchRequest_NotificationType = 3
	PluginDispatchRequest_WARNING                PluginDispatchRequest_NotificationType = 4
)

// Enum value maps for PluginDispatchRequest_NotificationType.
var (
	PluginDispatchRequest_NotificationType_name = map[int32]string{
		0: "NOTIFICATION_TYPE_NONE",
		1: "INFO",
		2: "ERROR",
		3: "SUCCESS",
		4: "WARNING",
	}
	PluginDispatchRequest_NotificationType_value = map[string]int32{
		"NOTIFICATION_TYPE_NONE": 0,
		"INFO":                   1,
		"ERROR":                  2,
		"SUCCESS":                3,
		"WARNING":                4,
	}
)

func (x PluginDispatchRequest_NotificationType) Enum() *PluginDispatchRequest_NotificationType {
	p := new(PluginDispatchRequest_NotificationType)
	*p = x
	return p
}

func (x PluginDispatchRequest_NotificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PluginDispatchRequest_NotificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_notification_plugin_notification_proto_enumTypes[0].Descriptor()
}

func (PluginDispatchRequest_NotificationType) Type() protoreflect.EnumType {
	return &file_spaceone_api_notification_plugin_notification_proto_enumTypes[0]
}

func (x PluginDispatchRequest_NotificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PluginDispatchRequest_NotificationType.Descriptor instead.
func (PluginDispatchRequest_NotificationType) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_notification_plugin_notification_proto_rawDescGZIP(), []int{0, 0}
}

//	{
//	   "options": {},
//	   "message": {
//	   "tags": [
//	       {
//	           "key": "Alert Number",
//	           "options": {"short": true},
//	           "value": "#108664"
//	       },
//	       {
//	           "options": {"short": true},
//	           "key": "State",
//	           "value": "TRIGGERED"
//	       },
//	       {
//	           "value": "LOW",
//	           "options": {"short": true},
//	           "key": "Urgency"
//	       },
//	       {
//	           "value": "kubectl-webhook",
//	           "key": "Triggered by",
//	           "options": {"short": true}
//	       },
//	       {
//	           "value": "SpaceONE > Project1",
//	           "key": "Project"
//	       },
//	       {
//	           "value": "spaceone-api",
//	           "key": "Resource"
//	       }
//	   ],
//	   "occurred_at": "2022-06-27T09:22:57.967Z",
//	   "callbacks": [{
//	       "url": "https://monitoring-webhook.dev.spaceone.dev/monitoring/v1/alert/alert-x1v2c3v456/8f2ede36213dqw4d7d5awe07ds32d883/ACKNOWLEDGED",
//	       "label": "Acknowledge Alerts"}],
//	   "link": "https://spaceone.console.dev.spaceone.dev/alert-manager/alert/alert-x1v2c3v456",
//	   "title": "[Alerting] Notification of access to the SpaceONE",
//	   "description": "SSH Access to spaceone-api from admin"},
//	   "notification_type": "INFO",
//	   "secret_data": "********",
//	   "channel_data": {"email": "test5@test.com"}
//	}
type PluginDispatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Option value required for notification delivery.
	Options *_struct.Struct `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	// Message containing notification information
	Message *_struct.Struct `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// The type of Notification
	NotificationType PluginDispatchRequest_NotificationType `protobuf:"varint,3,opt,name=notification_type,json=notificationType,proto3,enum=spaceone.api.notification.plugin.PluginDispatchRequest_NotificationType" json:"notification_type,omitempty"`
	// Secret value required for notification delivery.
	// The secret data usually includes the credential information required for the plugin to access the external system.
	SecretData *_struct.Struct `protobuf:"bytes,4,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	// Channel data required for notification delivery.
	ChannelData *_struct.Struct `protobuf:"bytes,5,opt,name=channel_data,json=channelData,proto3" json:"channel_data,omitempty"`
}

func (x *PluginDispatchRequest) Reset() {
	*x = PluginDispatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_notification_plugin_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginDispatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginDispatchRequest) ProtoMessage() {}

func (x *PluginDispatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_notification_plugin_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginDispatchRequest.ProtoReflect.Descriptor instead.
func (*PluginDispatchRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_notification_plugin_notification_proto_rawDescGZIP(), []int{0}
}

func (x *PluginDispatchRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *PluginDispatchRequest) GetMessage() *_struct.Struct {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *PluginDispatchRequest) GetNotificationType() PluginDispatchRequest_NotificationType {
	if x != nil {
		return x.NotificationType
	}
	return PluginDispatchRequest_NOTIFICATION_TYPE_NONE
}

func (x *PluginDispatchRequest) GetSecretData() *_struct.Struct {
	if x != nil {
		return x.SecretData
	}
	return nil
}

func (x *PluginDispatchRequest) GetChannelData() *_struct.Struct {
	if x != nil {
		return x.ChannelData
	}
	return nil
}

var File_spaceone_api_notification_plugin_notification_proto protoreflect.FileDescriptor

var file_spaceone_api_notification_plugin_notification_proto_rawDesc = []byte{
	0x0a, 0x33, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc9, 0x03, 0x0a, 0x15, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x31, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x75, 0x0a, 0x11, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x48, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x10, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x5d, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10,
	0x03, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x32, 0x6d,
	0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5d,
	0x0a, 0x08, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x37, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x47, 0x5a,
	0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69,
	0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_notification_plugin_notification_proto_rawDescOnce sync.Once
	file_spaceone_api_notification_plugin_notification_proto_rawDescData = file_spaceone_api_notification_plugin_notification_proto_rawDesc
)

func file_spaceone_api_notification_plugin_notification_proto_rawDescGZIP() []byte {
	file_spaceone_api_notification_plugin_notification_proto_rawDescOnce.Do(func() {
		file_spaceone_api_notification_plugin_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_notification_plugin_notification_proto_rawDescData)
	})
	return file_spaceone_api_notification_plugin_notification_proto_rawDescData
}

var file_spaceone_api_notification_plugin_notification_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spaceone_api_notification_plugin_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_spaceone_api_notification_plugin_notification_proto_goTypes = []interface{}{
	(PluginDispatchRequest_NotificationType)(0), // 0: spaceone.api.notification.plugin.PluginDispatchRequest.NotificationType
	(*PluginDispatchRequest)(nil),               // 1: spaceone.api.notification.plugin.PluginDispatchRequest
	(*_struct.Struct)(nil),                      // 2: google.protobuf.Struct
	(*empty.Empty)(nil),                         // 3: google.protobuf.Empty
}
var file_spaceone_api_notification_plugin_notification_proto_depIdxs = []int32{
	2, // 0: spaceone.api.notification.plugin.PluginDispatchRequest.options:type_name -> google.protobuf.Struct
	2, // 1: spaceone.api.notification.plugin.PluginDispatchRequest.message:type_name -> google.protobuf.Struct
	0, // 2: spaceone.api.notification.plugin.PluginDispatchRequest.notification_type:type_name -> spaceone.api.notification.plugin.PluginDispatchRequest.NotificationType
	2, // 3: spaceone.api.notification.plugin.PluginDispatchRequest.secret_data:type_name -> google.protobuf.Struct
	2, // 4: spaceone.api.notification.plugin.PluginDispatchRequest.channel_data:type_name -> google.protobuf.Struct
	1, // 5: spaceone.api.notification.plugin.Notification.dispatch:input_type -> spaceone.api.notification.plugin.PluginDispatchRequest
	3, // 6: spaceone.api.notification.plugin.Notification.dispatch:output_type -> google.protobuf.Empty
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_spaceone_api_notification_plugin_notification_proto_init() }
func file_spaceone_api_notification_plugin_notification_proto_init() {
	if File_spaceone_api_notification_plugin_notification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_notification_plugin_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginDispatchRequest); i {
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
			RawDescriptor: file_spaceone_api_notification_plugin_notification_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_notification_plugin_notification_proto_goTypes,
		DependencyIndexes: file_spaceone_api_notification_plugin_notification_proto_depIdxs,
		EnumInfos:         file_spaceone_api_notification_plugin_notification_proto_enumTypes,
		MessageInfos:      file_spaceone_api_notification_plugin_notification_proto_msgTypes,
	}.Build()
	File_spaceone_api_notification_plugin_notification_proto = out.File
	file_spaceone_api_notification_plugin_notification_proto_rawDesc = nil
	file_spaceone_api_notification_plugin_notification_proto_goTypes = nil
	file_spaceone_api_notification_plugin_notification_proto_depIdxs = nil
}
