// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.0
// source: grpc/v1/grpc.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Grpc_GetUserInfo_FullMethodName = "/api.grpc.v1.Grpc/GetUserInfo"
)

// GrpcClient is the client API for Grpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Grpc服务
type GrpcClient interface {
	// 获取用户信息
	GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
}

type grpcClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpcClient(cc grpc.ClientConnInterface) GrpcClient {
	return &grpcClient{cc}
}

func (c *grpcClient) GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, Grpc_GetUserInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcServer is the server API for Grpc service.
// All implementations must embed UnimplementedGrpcServer
// for forward compatibility.
//
// Grpc服务
type GrpcServer interface {
	// 获取用户信息
	GetUserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
	mustEmbedUnimplementedGrpcServer()
}

// UnimplementedGrpcServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGrpcServer struct{}

func (UnimplementedGrpcServer) GetUserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedGrpcServer) mustEmbedUnimplementedGrpcServer() {}
func (UnimplementedGrpcServer) testEmbeddedByValue()              {}

// UnsafeGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GrpcServer will
// result in compilation errors.
type UnsafeGrpcServer interface {
	mustEmbedUnimplementedGrpcServer()
}

func RegisterGrpcServer(s grpc.ServiceRegistrar, srv GrpcServer) {
	// If the following call pancis, it indicates UnimplementedGrpcServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Grpc_ServiceDesc, srv)
}

func _Grpc_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Grpc_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).GetUserInfo(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Grpc_ServiceDesc is the grpc.ServiceDesc for Grpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Grpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.grpc.v1.Grpc",
	HandlerType: (*GrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _Grpc_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/v1/grpc.proto",
}
