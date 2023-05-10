// A DataSource is a plugin instance collecting `metric` and `log` data from Cloudforet.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/monitoring/v1/data_source.proto

package v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DataSource_Register_FullMethodName     = "/spaceone.api.monitoring.v1.DataSource/register"
	DataSource_Update_FullMethodName       = "/spaceone.api.monitoring.v1.DataSource/update"
	DataSource_Enable_FullMethodName       = "/spaceone.api.monitoring.v1.DataSource/enable"
	DataSource_Disable_FullMethodName      = "/spaceone.api.monitoring.v1.DataSource/disable"
	DataSource_Deregister_FullMethodName   = "/spaceone.api.monitoring.v1.DataSource/deregister"
	DataSource_UpdatePlugin_FullMethodName = "/spaceone.api.monitoring.v1.DataSource/update_plugin"
	DataSource_VerifyPlugin_FullMethodName = "/spaceone.api.monitoring.v1.DataSource/verify_plugin"
	DataSource_Get_FullMethodName          = "/spaceone.api.monitoring.v1.DataSource/get"
	DataSource_List_FullMethodName         = "/spaceone.api.monitoring.v1.DataSource/list"
	DataSource_Stat_FullMethodName         = "/spaceone.api.monitoring.v1.DataSource/stat"
)

// DataSourceClient is the client API for DataSource service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataSourceClient interface {
	// Registers a DataSource with information of the plugin to use. Information of the plugin includes `version`, `provider`, `upgrade_mode`.\
	Register(ctx context.Context, in *RegisterDataSourceRequest, opts ...grpc.CallOption) (*DataSourceInfo, error)
	// Updates a specific DataSource. You can make changes in DataSource settings, including `name` and `tags`.
	Update(ctx context.Context, in *UpdateDataSourceRequest, opts ...grpc.CallOption) (*DataSourceInfo, error)
	// Enables a specific DataSource. By enabling a DataSource, you can communicate with an external cloud service via the plugin.
	Enable(ctx context.Context, in *DataSourceRequest, opts ...grpc.CallOption) (*DataSourceInfo, error)
	// Disables a specific DataSource. By disabling a DataSource, you can block communication with an external cloud service via the plugin.
	Disable(ctx context.Context, in *DataSourceRequest, opts ...grpc.CallOption) (*DataSourceInfo, error)
	// Deregisters and deletes a specific DataSource. You must specify the `data_source_id` of the DataSource to deregister.
	Deregister(ctx context.Context, in *DataSourceRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Updates the plugin of a specific DataSource. This method resets the plugin data in the DataSource to update the `metadata`.
	UpdatePlugin(ctx context.Context, in *UpdateDataSourcePluginRequest, opts ...grpc.CallOption) (*DataSourceInfo, error)
	// Verifies the plugin of a specific DataSource. This method validates the plugin data, `version` and `endpoint`.
	VerifyPlugin(ctx context.Context, in *DataSourceRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Gets a specific DataSource. Prints detailed information about the DataSource, including `name`, `state`, and `plugin_info`.
	Get(ctx context.Context, in *GetDataSourceRequest, opts ...grpc.CallOption) (*DataSourceInfo, error)
	// Gets a list of all DataSources. You can use a query to get a filtered list of DataSources.
	List(ctx context.Context, in *DataSourceQuery, opts ...grpc.CallOption) (*DataSourcesInfo, error)
	Stat(ctx context.Context, in *DataSourceStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type dataSourceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataSourceClient(cc grpc.ClientConnInterface) DataSourceClient {
	return &dataSourceClient{cc}
}

func (c *dataSourceClient) Register(ctx context.Context, in *RegisterDataSourceRequest, opts ...grpc.CallOption) (*DataSourceInfo, error) {
	out := new(DataSourceInfo)
	err := c.cc.Invoke(ctx, DataSource_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSourceClient) Update(ctx context.Context, in *UpdateDataSourceRequest, opts ...grpc.CallOption) (*DataSourceInfo, error) {
	out := new(DataSourceInfo)
	err := c.cc.Invoke(ctx, DataSource_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSourceClient) Enable(ctx context.Context, in *DataSourceRequest, opts ...grpc.CallOption) (*DataSourceInfo, error) {
	out := new(DataSourceInfo)
	err := c.cc.Invoke(ctx, DataSource_Enable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSourceClient) Disable(ctx context.Context, in *DataSourceRequest, opts ...grpc.CallOption) (*DataSourceInfo, error) {
	out := new(DataSourceInfo)
	err := c.cc.Invoke(ctx, DataSource_Disable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSourceClient) Deregister(ctx context.Context, in *DataSourceRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, DataSource_Deregister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSourceClient) UpdatePlugin(ctx context.Context, in *UpdateDataSourcePluginRequest, opts ...grpc.CallOption) (*DataSourceInfo, error) {
	out := new(DataSourceInfo)
	err := c.cc.Invoke(ctx, DataSource_UpdatePlugin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSourceClient) VerifyPlugin(ctx context.Context, in *DataSourceRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, DataSource_VerifyPlugin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSourceClient) Get(ctx context.Context, in *GetDataSourceRequest, opts ...grpc.CallOption) (*DataSourceInfo, error) {
	out := new(DataSourceInfo)
	err := c.cc.Invoke(ctx, DataSource_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSourceClient) List(ctx context.Context, in *DataSourceQuery, opts ...grpc.CallOption) (*DataSourcesInfo, error) {
	out := new(DataSourcesInfo)
	err := c.cc.Invoke(ctx, DataSource_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSourceClient) Stat(ctx context.Context, in *DataSourceStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, DataSource_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataSourceServer is the server API for DataSource service.
// All implementations must embed UnimplementedDataSourceServer
// for forward compatibility
type DataSourceServer interface {
	// Registers a DataSource with information of the plugin to use. Information of the plugin includes `version`, `provider`, `upgrade_mode`.\
	Register(context.Context, *RegisterDataSourceRequest) (*DataSourceInfo, error)
	// Updates a specific DataSource. You can make changes in DataSource settings, including `name` and `tags`.
	Update(context.Context, *UpdateDataSourceRequest) (*DataSourceInfo, error)
	// Enables a specific DataSource. By enabling a DataSource, you can communicate with an external cloud service via the plugin.
	Enable(context.Context, *DataSourceRequest) (*DataSourceInfo, error)
	// Disables a specific DataSource. By disabling a DataSource, you can block communication with an external cloud service via the plugin.
	Disable(context.Context, *DataSourceRequest) (*DataSourceInfo, error)
	// Deregisters and deletes a specific DataSource. You must specify the `data_source_id` of the DataSource to deregister.
	Deregister(context.Context, *DataSourceRequest) (*empty.Empty, error)
	// Updates the plugin of a specific DataSource. This method resets the plugin data in the DataSource to update the `metadata`.
	UpdatePlugin(context.Context, *UpdateDataSourcePluginRequest) (*DataSourceInfo, error)
	// Verifies the plugin of a specific DataSource. This method validates the plugin data, `version` and `endpoint`.
	VerifyPlugin(context.Context, *DataSourceRequest) (*empty.Empty, error)
	// Gets a specific DataSource. Prints detailed information about the DataSource, including `name`, `state`, and `plugin_info`.
	Get(context.Context, *GetDataSourceRequest) (*DataSourceInfo, error)
	// Gets a list of all DataSources. You can use a query to get a filtered list of DataSources.
	List(context.Context, *DataSourceQuery) (*DataSourcesInfo, error)
	Stat(context.Context, *DataSourceStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedDataSourceServer()
}

// UnimplementedDataSourceServer must be embedded to have forward compatible implementations.
type UnimplementedDataSourceServer struct {
}

func (UnimplementedDataSourceServer) Register(context.Context, *RegisterDataSourceRequest) (*DataSourceInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedDataSourceServer) Update(context.Context, *UpdateDataSourceRequest) (*DataSourceInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDataSourceServer) Enable(context.Context, *DataSourceRequest) (*DataSourceInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedDataSourceServer) Disable(context.Context, *DataSourceRequest) (*DataSourceInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedDataSourceServer) Deregister(context.Context, *DataSourceRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deregister not implemented")
}
func (UnimplementedDataSourceServer) UpdatePlugin(context.Context, *UpdateDataSourcePluginRequest) (*DataSourceInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlugin not implemented")
}
func (UnimplementedDataSourceServer) VerifyPlugin(context.Context, *DataSourceRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPlugin not implemented")
}
func (UnimplementedDataSourceServer) Get(context.Context, *GetDataSourceRequest) (*DataSourceInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDataSourceServer) List(context.Context, *DataSourceQuery) (*DataSourcesInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDataSourceServer) Stat(context.Context, *DataSourceStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedDataSourceServer) mustEmbedUnimplementedDataSourceServer() {}

// UnsafeDataSourceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataSourceServer will
// result in compilation errors.
type UnsafeDataSourceServer interface {
	mustEmbedUnimplementedDataSourceServer()
}

func RegisterDataSourceServer(s grpc.ServiceRegistrar, srv DataSourceServer) {
	s.RegisterService(&DataSource_ServiceDesc, srv)
}

func _DataSource_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDataSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataSource_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceServer).Register(ctx, req.(*RegisterDataSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSource_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDataSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataSource_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceServer).Update(ctx, req.(*UpdateDataSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSource_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataSource_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceServer).Enable(ctx, req.(*DataSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSource_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataSource_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceServer).Disable(ctx, req.(*DataSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSource_Deregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceServer).Deregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataSource_Deregister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceServer).Deregister(ctx, req.(*DataSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSource_UpdatePlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDataSourcePluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceServer).UpdatePlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataSource_UpdatePlugin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceServer).UpdatePlugin(ctx, req.(*UpdateDataSourcePluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSource_VerifyPlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceServer).VerifyPlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataSource_VerifyPlugin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceServer).VerifyPlugin(ctx, req.(*DataSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSource_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataSource_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceServer).Get(ctx, req.(*GetDataSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSource_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSourceQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataSource_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceServer).List(ctx, req.(*DataSourceQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSource_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSourceStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataSource_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceServer).Stat(ctx, req.(*DataSourceStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// DataSource_ServiceDesc is the grpc.ServiceDesc for DataSource service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataSource_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.monitoring.v1.DataSource",
	HandlerType: (*DataSourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "register",
			Handler:    _DataSource_Register_Handler,
		},
		{
			MethodName: "update",
			Handler:    _DataSource_Update_Handler,
		},
		{
			MethodName: "enable",
			Handler:    _DataSource_Enable_Handler,
		},
		{
			MethodName: "disable",
			Handler:    _DataSource_Disable_Handler,
		},
		{
			MethodName: "deregister",
			Handler:    _DataSource_Deregister_Handler,
		},
		{
			MethodName: "update_plugin",
			Handler:    _DataSource_UpdatePlugin_Handler,
		},
		{
			MethodName: "verify_plugin",
			Handler:    _DataSource_VerifyPlugin_Handler,
		},
		{
			MethodName: "get",
			Handler:    _DataSource_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _DataSource_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _DataSource_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/monitoring/v1/data_source.proto",
}
