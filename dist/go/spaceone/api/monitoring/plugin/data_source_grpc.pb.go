// A DataSource is a plugin instance receiving Metric and Log data from cloud services.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/monitoring/plugin/data_source.proto

package plugin

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DataSource_Init_FullMethodName   = "/spaceone.api.monitoring.plugin.DataSource/init"
	DataSource_Verify_FullMethodName = "/spaceone.api.monitoring.plugin.DataSource/verify"
)

// DataSourceClient is the client API for DataSource service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataSourceClient interface {
	// Initializes a specific DataSource. During initialization, the DataSource information to be passed to the DataSource user is delivered as `metadata`. DataSource information includes its name and version.
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*PluginInfo, error)
	// Verifies a specific DataSource. You must specify the parameter `secret_data`, encrypted account data of the DataSource to validate.
	Verify(ctx context.Context, in *PluginVerifyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type dataSourceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataSourceClient(cc grpc.ClientConnInterface) DataSourceClient {
	return &dataSourceClient{cc}
}

func (c *dataSourceClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*PluginInfo, error) {
	out := new(PluginInfo)
	err := c.cc.Invoke(ctx, DataSource_Init_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSourceClient) Verify(ctx context.Context, in *PluginVerifyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, DataSource_Verify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataSourceServer is the server API for DataSource service.
// All implementations must embed UnimplementedDataSourceServer
// for forward compatibility
type DataSourceServer interface {
	// Initializes a specific DataSource. During initialization, the DataSource information to be passed to the DataSource user is delivered as `metadata`. DataSource information includes its name and version.
	Init(context.Context, *InitRequest) (*PluginInfo, error)
	// Verifies a specific DataSource. You must specify the parameter `secret_data`, encrypted account data of the DataSource to validate.
	Verify(context.Context, *PluginVerifyRequest) (*empty.Empty, error)
	mustEmbedUnimplementedDataSourceServer()
}

// UnimplementedDataSourceServer must be embedded to have forward compatible implementations.
type UnimplementedDataSourceServer struct {
}

func (UnimplementedDataSourceServer) Init(context.Context, *InitRequest) (*PluginInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedDataSourceServer) Verify(context.Context, *PluginVerifyRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
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

func _DataSource_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataSource_Init_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSource_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataSource_Verify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceServer).Verify(ctx, req.(*PluginVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataSource_ServiceDesc is the grpc.ServiceDesc for DataSource service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataSource_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.monitoring.plugin.DataSource",
	HandlerType: (*DataSourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "init",
			Handler:    _DataSource_Init_Handler,
		},
		{
			MethodName: "verify",
			Handler:    _DataSource_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/monitoring/plugin/data_source.proto",
}
