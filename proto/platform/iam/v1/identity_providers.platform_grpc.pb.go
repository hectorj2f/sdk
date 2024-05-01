// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: identity_providers.platform.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	IdentityProviders_Create_FullMethodName = "/chainguard.platform.iam.IdentityProviders/Create"
	IdentityProviders_Update_FullMethodName = "/chainguard.platform.iam.IdentityProviders/Update"
	IdentityProviders_List_FullMethodName   = "/chainguard.platform.iam.IdentityProviders/List"
	IdentityProviders_Delete_FullMethodName = "/chainguard.platform.iam.IdentityProviders/Delete"
)

// IdentityProvidersClient is the client API for IdentityProviders service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityProvidersClient interface {
	Create(ctx context.Context, in *CreateIdentityProviderRequest, opts ...grpc.CallOption) (*IdentityProvider, error)
	Update(ctx context.Context, in *IdentityProvider, opts ...grpc.CallOption) (*IdentityProvider, error)
	List(ctx context.Context, in *IdentityProviderFilter, opts ...grpc.CallOption) (*IdentityProviderList, error)
	Delete(ctx context.Context, in *DeleteIdentityProviderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type identityProvidersClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityProvidersClient(cc grpc.ClientConnInterface) IdentityProvidersClient {
	return &identityProvidersClient{cc}
}

func (c *identityProvidersClient) Create(ctx context.Context, in *CreateIdentityProviderRequest, opts ...grpc.CallOption) (*IdentityProvider, error) {
	out := new(IdentityProvider)
	err := c.cc.Invoke(ctx, IdentityProviders_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityProvidersClient) Update(ctx context.Context, in *IdentityProvider, opts ...grpc.CallOption) (*IdentityProvider, error) {
	out := new(IdentityProvider)
	err := c.cc.Invoke(ctx, IdentityProviders_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityProvidersClient) List(ctx context.Context, in *IdentityProviderFilter, opts ...grpc.CallOption) (*IdentityProviderList, error) {
	out := new(IdentityProviderList)
	err := c.cc.Invoke(ctx, IdentityProviders_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityProvidersClient) Delete(ctx context.Context, in *DeleteIdentityProviderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, IdentityProviders_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityProvidersServer is the server API for IdentityProviders service.
// All implementations must embed UnimplementedIdentityProvidersServer
// for forward compatibility
type IdentityProvidersServer interface {
	Create(context.Context, *CreateIdentityProviderRequest) (*IdentityProvider, error)
	Update(context.Context, *IdentityProvider) (*IdentityProvider, error)
	List(context.Context, *IdentityProviderFilter) (*IdentityProviderList, error)
	Delete(context.Context, *DeleteIdentityProviderRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedIdentityProvidersServer()
}

// UnimplementedIdentityProvidersServer must be embedded to have forward compatible implementations.
type UnimplementedIdentityProvidersServer struct {
}

func (UnimplementedIdentityProvidersServer) Create(context.Context, *CreateIdentityProviderRequest) (*IdentityProvider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedIdentityProvidersServer) Update(context.Context, *IdentityProvider) (*IdentityProvider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedIdentityProvidersServer) List(context.Context, *IdentityProviderFilter) (*IdentityProviderList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedIdentityProvidersServer) Delete(context.Context, *DeleteIdentityProviderRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedIdentityProvidersServer) mustEmbedUnimplementedIdentityProvidersServer() {}

// UnsafeIdentityProvidersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityProvidersServer will
// result in compilation errors.
type UnsafeIdentityProvidersServer interface {
	mustEmbedUnimplementedIdentityProvidersServer()
}

func RegisterIdentityProvidersServer(s grpc.ServiceRegistrar, srv IdentityProvidersServer) {
	s.RegisterService(&IdentityProviders_ServiceDesc, srv)
}

func _IdentityProviders_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIdentityProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityProvidersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityProviders_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityProvidersServer).Create(ctx, req.(*CreateIdentityProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityProviders_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityProvider)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityProvidersServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityProviders_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityProvidersServer).Update(ctx, req.(*IdentityProvider))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityProviders_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityProviderFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityProvidersServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityProviders_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityProvidersServer).List(ctx, req.(*IdentityProviderFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityProviders_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIdentityProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityProvidersServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityProviders_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityProvidersServer).Delete(ctx, req.(*DeleteIdentityProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IdentityProviders_ServiceDesc is the grpc.ServiceDesc for IdentityProviders service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdentityProviders_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chainguard.platform.iam.IdentityProviders",
	HandlerType: (*IdentityProvidersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _IdentityProviders_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _IdentityProviders_Update_Handler,
		},
		{
			MethodName: "List",
			Handler:    _IdentityProviders_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _IdentityProviders_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity_providers.platform.proto",
}
