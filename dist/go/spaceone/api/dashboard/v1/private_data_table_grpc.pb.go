// description of data table

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/dashboard/v1/private_data_table.proto

package v1

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
	PrivateDataTable_Create_FullMethodName = "/spaceone.api.data_table.v1.PrivateDataTable/create"
	PrivateDataTable_Update_FullMethodName = "/spaceone.api.data_table.v1.PrivateDataTable/update"
	PrivateDataTable_Delete_FullMethodName = "/spaceone.api.data_table.v1.PrivateDataTable/delete"
	PrivateDataTable_Load_FullMethodName   = "/spaceone.api.data_table.v1.PrivateDataTable/load"
	PrivateDataTable_Get_FullMethodName    = "/spaceone.api.data_table.v1.PrivateDataTable/get"
	PrivateDataTable_List_FullMethodName   = "/spaceone.api.data_table.v1.PrivateDataTable/list"
)

// PrivateDataTableClient is the client API for PrivateDataTable service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrivateDataTableClient interface {
	Create(ctx context.Context, in *CreatePrivateDataTableRequest, opts ...grpc.CallOption) (*PrivateDataTableInfo, error)
	Update(ctx context.Context, in *UpdatePrivateDataTableRequest, opts ...grpc.CallOption) (*PrivateDataTableInfo, error)
	Delete(ctx context.Context, in *PrivateDataTableRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Load(ctx context.Context, in *LoadPrivateDataTableRequest, opts ...grpc.CallOption) (*PrivateDataTableInfo, error)
	Get(ctx context.Context, in *PrivateDataTableRequest, opts ...grpc.CallOption) (*PrivateDataTableInfo, error)
	List(ctx context.Context, in *PrivateDataTableQuery, opts ...grpc.CallOption) (*PrivateDataTablesInfo, error)
}

type privateDataTableClient struct {
	cc grpc.ClientConnInterface
}

func NewPrivateDataTableClient(cc grpc.ClientConnInterface) PrivateDataTableClient {
	return &privateDataTableClient{cc}
}

func (c *privateDataTableClient) Create(ctx context.Context, in *CreatePrivateDataTableRequest, opts ...grpc.CallOption) (*PrivateDataTableInfo, error) {
	out := new(PrivateDataTableInfo)
	err := c.cc.Invoke(ctx, PrivateDataTable_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateDataTableClient) Update(ctx context.Context, in *UpdatePrivateDataTableRequest, opts ...grpc.CallOption) (*PrivateDataTableInfo, error) {
	out := new(PrivateDataTableInfo)
	err := c.cc.Invoke(ctx, PrivateDataTable_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateDataTableClient) Delete(ctx context.Context, in *PrivateDataTableRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, PrivateDataTable_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateDataTableClient) Load(ctx context.Context, in *LoadPrivateDataTableRequest, opts ...grpc.CallOption) (*PrivateDataTableInfo, error) {
	out := new(PrivateDataTableInfo)
	err := c.cc.Invoke(ctx, PrivateDataTable_Load_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateDataTableClient) Get(ctx context.Context, in *PrivateDataTableRequest, opts ...grpc.CallOption) (*PrivateDataTableInfo, error) {
	out := new(PrivateDataTableInfo)
	err := c.cc.Invoke(ctx, PrivateDataTable_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateDataTableClient) List(ctx context.Context, in *PrivateDataTableQuery, opts ...grpc.CallOption) (*PrivateDataTablesInfo, error) {
	out := new(PrivateDataTablesInfo)
	err := c.cc.Invoke(ctx, PrivateDataTable_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivateDataTableServer is the server API for PrivateDataTable service.
// All implementations must embed UnimplementedPrivateDataTableServer
// for forward compatibility
type PrivateDataTableServer interface {
	Create(context.Context, *CreatePrivateDataTableRequest) (*PrivateDataTableInfo, error)
	Update(context.Context, *UpdatePrivateDataTableRequest) (*PrivateDataTableInfo, error)
	Delete(context.Context, *PrivateDataTableRequest) (*empty.Empty, error)
	Load(context.Context, *LoadPrivateDataTableRequest) (*PrivateDataTableInfo, error)
	Get(context.Context, *PrivateDataTableRequest) (*PrivateDataTableInfo, error)
	List(context.Context, *PrivateDataTableQuery) (*PrivateDataTablesInfo, error)
	mustEmbedUnimplementedPrivateDataTableServer()
}

// UnimplementedPrivateDataTableServer must be embedded to have forward compatible implementations.
type UnimplementedPrivateDataTableServer struct {
}

func (UnimplementedPrivateDataTableServer) Create(context.Context, *CreatePrivateDataTableRequest) (*PrivateDataTableInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPrivateDataTableServer) Update(context.Context, *UpdatePrivateDataTableRequest) (*PrivateDataTableInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPrivateDataTableServer) Delete(context.Context, *PrivateDataTableRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPrivateDataTableServer) Load(context.Context, *LoadPrivateDataTableRequest) (*PrivateDataTableInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Load not implemented")
}
func (UnimplementedPrivateDataTableServer) Get(context.Context, *PrivateDataTableRequest) (*PrivateDataTableInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPrivateDataTableServer) List(context.Context, *PrivateDataTableQuery) (*PrivateDataTablesInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPrivateDataTableServer) mustEmbedUnimplementedPrivateDataTableServer() {}

// UnsafePrivateDataTableServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrivateDataTableServer will
// result in compilation errors.
type UnsafePrivateDataTableServer interface {
	mustEmbedUnimplementedPrivateDataTableServer()
}

func RegisterPrivateDataTableServer(s grpc.ServiceRegistrar, srv PrivateDataTableServer) {
	s.RegisterService(&PrivateDataTable_ServiceDesc, srv)
}

func _PrivateDataTable_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePrivateDataTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateDataTableServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateDataTable_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateDataTableServer).Create(ctx, req.(*CreatePrivateDataTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateDataTable_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePrivateDataTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateDataTableServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateDataTable_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateDataTableServer).Update(ctx, req.(*UpdatePrivateDataTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateDataTable_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivateDataTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateDataTableServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateDataTable_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateDataTableServer).Delete(ctx, req.(*PrivateDataTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateDataTable_Load_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadPrivateDataTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateDataTableServer).Load(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateDataTable_Load_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateDataTableServer).Load(ctx, req.(*LoadPrivateDataTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateDataTable_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivateDataTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateDataTableServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateDataTable_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateDataTableServer).Get(ctx, req.(*PrivateDataTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateDataTable_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivateDataTableQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateDataTableServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateDataTable_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateDataTableServer).List(ctx, req.(*PrivateDataTableQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// PrivateDataTable_ServiceDesc is the grpc.ServiceDesc for PrivateDataTable service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrivateDataTable_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.data_table.v1.PrivateDataTable",
	HandlerType: (*PrivateDataTableServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _PrivateDataTable_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _PrivateDataTable_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _PrivateDataTable_Delete_Handler,
		},
		{
			MethodName: "load",
			Handler:    _PrivateDataTable_Load_Handler,
		},
		{
			MethodName: "get",
			Handler:    _PrivateDataTable_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _PrivateDataTable_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/dashboard/v1/private_data_table.proto",
}