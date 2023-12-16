// Trusted secret is a service that stores and manages credentials.
// Trusted secret is merged with linked secret and used to access data in other microservices.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/secret/v1/trusted_secret.proto

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
	TrustedSecret_Create_FullMethodName     = "/spaceone.api.secret.v1.TrustedSecret/create"
	TrustedSecret_Update_FullMethodName     = "/spaceone.api.secret.v1.TrustedSecret/update"
	TrustedSecret_Delete_FullMethodName     = "/spaceone.api.secret.v1.TrustedSecret/delete"
	TrustedSecret_UpdateData_FullMethodName = "/spaceone.api.secret.v1.TrustedSecret/update_data"
	TrustedSecret_Get_FullMethodName        = "/spaceone.api.secret.v1.TrustedSecret/get"
	TrustedSecret_List_FullMethodName       = "/spaceone.api.secret.v1.TrustedSecret/list"
	TrustedSecret_Stat_FullMethodName       = "/spaceone.api.secret.v1.TrustedSecret/stat"
)

// TrustedSecretClient is the client API for TrustedSecret service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrustedSecretClient interface {
	// Create a new trusted secret.
	// Created trusted secret is encrypted and stored securely.
	Create(ctx context.Context, in *CreateTrustedSecretRequest, opts ...grpc.CallOption) (*TrustedSecretInfo, error)
	// Updates a specific trusted secret's information.
	// You can only change the 'name' and 'tags', and to change the data you must use the update_data API.
	Update(ctx context.Context, in *UpdateTrustedSecretRequest, opts ...grpc.CallOption) (*TrustedSecretInfo, error)
	// Deletes a specific trusted secret.
	// If a trusted secret is attached to a Secret, it cannot be deleted.
	Delete(ctx context.Context, in *TrustedSecretRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Updates a specific trusted secret's data.
	// Updated trusted secret is encrypted and stored securely.
	UpdateData(ctx context.Context, in *UpdateTrustedSecretDataRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Get a specific trusted secret's information.
	Get(ctx context.Context, in *TrustedSecretRequest, opts ...grpc.CallOption) (*TrustedSecretInfo, error)
	// Queries a list of trusted secrets.
	// You can use a query to get a filtered list of trusted secrets.
	List(ctx context.Context, in *TrustedSecretQuery, opts ...grpc.CallOption) (*TrustedSecretsInfo, error)
	Stat(ctx context.Context, in *TrustedSecretStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type trustedSecretClient struct {
	cc grpc.ClientConnInterface
}

func NewTrustedSecretClient(cc grpc.ClientConnInterface) TrustedSecretClient {
	return &trustedSecretClient{cc}
}

func (c *trustedSecretClient) Create(ctx context.Context, in *CreateTrustedSecretRequest, opts ...grpc.CallOption) (*TrustedSecretInfo, error) {
	out := new(TrustedSecretInfo)
	err := c.cc.Invoke(ctx, TrustedSecret_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustedSecretClient) Update(ctx context.Context, in *UpdateTrustedSecretRequest, opts ...grpc.CallOption) (*TrustedSecretInfo, error) {
	out := new(TrustedSecretInfo)
	err := c.cc.Invoke(ctx, TrustedSecret_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustedSecretClient) Delete(ctx context.Context, in *TrustedSecretRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, TrustedSecret_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustedSecretClient) UpdateData(ctx context.Context, in *UpdateTrustedSecretDataRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, TrustedSecret_UpdateData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustedSecretClient) Get(ctx context.Context, in *TrustedSecretRequest, opts ...grpc.CallOption) (*TrustedSecretInfo, error) {
	out := new(TrustedSecretInfo)
	err := c.cc.Invoke(ctx, TrustedSecret_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustedSecretClient) List(ctx context.Context, in *TrustedSecretQuery, opts ...grpc.CallOption) (*TrustedSecretsInfo, error) {
	out := new(TrustedSecretsInfo)
	err := c.cc.Invoke(ctx, TrustedSecret_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustedSecretClient) Stat(ctx context.Context, in *TrustedSecretStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, TrustedSecret_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrustedSecretServer is the server API for TrustedSecret service.
// All implementations must embed UnimplementedTrustedSecretServer
// for forward compatibility
type TrustedSecretServer interface {
	// Create a new trusted secret.
	// Created trusted secret is encrypted and stored securely.
	Create(context.Context, *CreateTrustedSecretRequest) (*TrustedSecretInfo, error)
	// Updates a specific trusted secret's information.
	// You can only change the 'name' and 'tags', and to change the data you must use the update_data API.
	Update(context.Context, *UpdateTrustedSecretRequest) (*TrustedSecretInfo, error)
	// Deletes a specific trusted secret.
	// If a trusted secret is attached to a Secret, it cannot be deleted.
	Delete(context.Context, *TrustedSecretRequest) (*empty.Empty, error)
	// Updates a specific trusted secret's data.
	// Updated trusted secret is encrypted and stored securely.
	UpdateData(context.Context, *UpdateTrustedSecretDataRequest) (*empty.Empty, error)
	// Get a specific trusted secret's information.
	Get(context.Context, *TrustedSecretRequest) (*TrustedSecretInfo, error)
	// Queries a list of trusted secrets.
	// You can use a query to get a filtered list of trusted secrets.
	List(context.Context, *TrustedSecretQuery) (*TrustedSecretsInfo, error)
	Stat(context.Context, *TrustedSecretStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedTrustedSecretServer()
}

// UnimplementedTrustedSecretServer must be embedded to have forward compatible implementations.
type UnimplementedTrustedSecretServer struct {
}

func (UnimplementedTrustedSecretServer) Create(context.Context, *CreateTrustedSecretRequest) (*TrustedSecretInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTrustedSecretServer) Update(context.Context, *UpdateTrustedSecretRequest) (*TrustedSecretInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTrustedSecretServer) Delete(context.Context, *TrustedSecretRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTrustedSecretServer) UpdateData(context.Context, *UpdateTrustedSecretDataRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateData not implemented")
}
func (UnimplementedTrustedSecretServer) Get(context.Context, *TrustedSecretRequest) (*TrustedSecretInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTrustedSecretServer) List(context.Context, *TrustedSecretQuery) (*TrustedSecretsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTrustedSecretServer) Stat(context.Context, *TrustedSecretStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedTrustedSecretServer) mustEmbedUnimplementedTrustedSecretServer() {}

// UnsafeTrustedSecretServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrustedSecretServer will
// result in compilation errors.
type UnsafeTrustedSecretServer interface {
	mustEmbedUnimplementedTrustedSecretServer()
}

func RegisterTrustedSecretServer(s grpc.ServiceRegistrar, srv TrustedSecretServer) {
	s.RegisterService(&TrustedSecret_ServiceDesc, srv)
}

func _TrustedSecret_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTrustedSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustedSecretServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustedSecret_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustedSecretServer).Create(ctx, req.(*CreateTrustedSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustedSecret_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTrustedSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustedSecretServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustedSecret_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustedSecretServer).Update(ctx, req.(*UpdateTrustedSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustedSecret_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrustedSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustedSecretServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustedSecret_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustedSecretServer).Delete(ctx, req.(*TrustedSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustedSecret_UpdateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTrustedSecretDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustedSecretServer).UpdateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustedSecret_UpdateData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustedSecretServer).UpdateData(ctx, req.(*UpdateTrustedSecretDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustedSecret_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrustedSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustedSecretServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustedSecret_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustedSecretServer).Get(ctx, req.(*TrustedSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustedSecret_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrustedSecretQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustedSecretServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustedSecret_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustedSecretServer).List(ctx, req.(*TrustedSecretQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustedSecret_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrustedSecretStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustedSecretServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustedSecret_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustedSecretServer).Stat(ctx, req.(*TrustedSecretStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// TrustedSecret_ServiceDesc is the grpc.ServiceDesc for TrustedSecret service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrustedSecret_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.secret.v1.TrustedSecret",
	HandlerType: (*TrustedSecretServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _TrustedSecret_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _TrustedSecret_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _TrustedSecret_Delete_Handler,
		},
		{
			MethodName: "update_data",
			Handler:    _TrustedSecret_UpdateData_Handler,
		},
		{
			MethodName: "get",
			Handler:    _TrustedSecret_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _TrustedSecret_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _TrustedSecret_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/secret/v1/trusted_secret.proto",
}
