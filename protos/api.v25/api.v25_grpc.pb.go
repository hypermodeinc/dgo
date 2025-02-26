//
// SPDX-FileCopyrightText: © Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

// Style guide for Protocol Buffer 3.
// Use CamelCase (with an initial capital) for message names – for example,
// SongServerRequest. Use underscore_separated_names for field names – for
// example, song_name.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: api.v25.proto

package api_v25

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DgraphHM_Ping_FullMethodName            = "/api.v25.DgraphHM/Ping"
	DgraphHM_LoginUser_FullMethodName       = "/api.v25.DgraphHM/LoginUser"
	DgraphHM_CreateNamespace_FullMethodName = "/api.v25.DgraphHM/CreateNamespace"
	DgraphHM_DropNamespace_FullMethodName   = "/api.v25.DgraphHM/DropNamespace"
	DgraphHM_UpdateNamespace_FullMethodName = "/api.v25.DgraphHM/UpdateNamespace"
	DgraphHM_ListNamespaces_FullMethodName  = "/api.v25.DgraphHM/ListNamespaces"
)

// DgraphHMClient is the client API for DgraphHM service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DgraphHMClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*CreateNamespaceResponse, error)
	DropNamespace(ctx context.Context, in *DropNamespaceRequest, opts ...grpc.CallOption) (*DropNamespaceResponse, error)
	UpdateNamespace(ctx context.Context, in *UpdateNamespaceRequest, opts ...grpc.CallOption) (*UpdateNamespaceResponse, error)
	ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error)
}

type dgraphHMClient struct {
	cc grpc.ClientConnInterface
}

func NewDgraphHMClient(cc grpc.ClientConnInterface) DgraphHMClient {
	return &dgraphHMClient{cc}
}

func (c *dgraphHMClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, DgraphHM_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphHMClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, DgraphHM_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphHMClient) CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*CreateNamespaceResponse, error) {
	out := new(CreateNamespaceResponse)
	err := c.cc.Invoke(ctx, DgraphHM_CreateNamespace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphHMClient) DropNamespace(ctx context.Context, in *DropNamespaceRequest, opts ...grpc.CallOption) (*DropNamespaceResponse, error) {
	out := new(DropNamespaceResponse)
	err := c.cc.Invoke(ctx, DgraphHM_DropNamespace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphHMClient) UpdateNamespace(ctx context.Context, in *UpdateNamespaceRequest, opts ...grpc.CallOption) (*UpdateNamespaceResponse, error) {
	out := new(UpdateNamespaceResponse)
	err := c.cc.Invoke(ctx, DgraphHM_UpdateNamespace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphHMClient) ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error) {
	out := new(ListNamespacesResponse)
	err := c.cc.Invoke(ctx, DgraphHM_ListNamespaces_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DgraphHMServer is the server API for DgraphHM service.
// All implementations must embed UnimplementedDgraphHMServer
// for forward compatibility
type DgraphHMServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	CreateNamespace(context.Context, *CreateNamespaceRequest) (*CreateNamespaceResponse, error)
	DropNamespace(context.Context, *DropNamespaceRequest) (*DropNamespaceResponse, error)
	UpdateNamespace(context.Context, *UpdateNamespaceRequest) (*UpdateNamespaceResponse, error)
	ListNamespaces(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error)
	mustEmbedUnimplementedDgraphHMServer()
}

// UnimplementedDgraphHMServer must be embedded to have forward compatible implementations.
type UnimplementedDgraphHMServer struct {
}

func (UnimplementedDgraphHMServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedDgraphHMServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedDgraphHMServer) CreateNamespace(context.Context, *CreateNamespaceRequest) (*CreateNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNamespace not implemented")
}
func (UnimplementedDgraphHMServer) DropNamespace(context.Context, *DropNamespaceRequest) (*DropNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropNamespace not implemented")
}
func (UnimplementedDgraphHMServer) UpdateNamespace(context.Context, *UpdateNamespaceRequest) (*UpdateNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNamespace not implemented")
}
func (UnimplementedDgraphHMServer) ListNamespaces(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNamespaces not implemented")
}
func (UnimplementedDgraphHMServer) mustEmbedUnimplementedDgraphHMServer() {}

// UnsafeDgraphHMServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DgraphHMServer will
// result in compilation errors.
type UnsafeDgraphHMServer interface {
	mustEmbedUnimplementedDgraphHMServer()
}

func RegisterDgraphHMServer(s grpc.ServiceRegistrar, srv DgraphHMServer) {
	s.RegisterService(&DgraphHM_ServiceDesc, srv)
}

func _DgraphHM_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphHMServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DgraphHM_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphHMServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphHM_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphHMServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DgraphHM_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphHMServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphHM_CreateNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphHMServer).CreateNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DgraphHM_CreateNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphHMServer).CreateNamespace(ctx, req.(*CreateNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphHM_DropNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphHMServer).DropNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DgraphHM_DropNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphHMServer).DropNamespace(ctx, req.(*DropNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphHM_UpdateNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphHMServer).UpdateNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DgraphHM_UpdateNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphHMServer).UpdateNamespace(ctx, req.(*UpdateNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphHM_ListNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNamespacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphHMServer).ListNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DgraphHM_ListNamespaces_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphHMServer).ListNamespaces(ctx, req.(*ListNamespacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DgraphHM_ServiceDesc is the grpc.ServiceDesc for DgraphHM service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DgraphHM_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v25.DgraphHM",
	HandlerType: (*DgraphHMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _DgraphHM_Ping_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _DgraphHM_LoginUser_Handler,
		},
		{
			MethodName: "CreateNamespace",
			Handler:    _DgraphHM_CreateNamespace_Handler,
		},
		{
			MethodName: "DropNamespace",
			Handler:    _DgraphHM_DropNamespace_Handler,
		},
		{
			MethodName: "UpdateNamespace",
			Handler:    _DgraphHM_UpdateNamespace_Handler,
		},
		{
			MethodName: "ListNamespaces",
			Handler:    _DgraphHM_ListNamespaces_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.v25.proto",
}
