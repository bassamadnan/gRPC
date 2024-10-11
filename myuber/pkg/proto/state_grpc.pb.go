// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: pkg/proto/state.proto

package spb

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
	StateService_GetState_FullMethodName = "/myuber.StateService/GetState"
	StateService_SetState_FullMethodName = "/myuber.StateService/SetState"
)

// StateServiceClient is the client API for StateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StateServiceClient interface {
	GetState(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*State, error)
	SetState(ctx context.Context, in *State, opts ...grpc.CallOption) (*Empty, error)
}

type stateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStateServiceClient(cc grpc.ClientConnInterface) StateServiceClient {
	return &stateServiceClient{cc}
}

func (c *stateServiceClient) GetState(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*State, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(State)
	err := c.cc.Invoke(ctx, StateService_GetState_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateServiceClient) SetState(ctx context.Context, in *State, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, StateService_SetState_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StateServiceServer is the server API for StateService service.
// All implementations must embed UnimplementedStateServiceServer
// for forward compatibility.
type StateServiceServer interface {
	GetState(context.Context, *Empty) (*State, error)
	SetState(context.Context, *State) (*Empty, error)
	mustEmbedUnimplementedStateServiceServer()
}

// UnimplementedStateServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStateServiceServer struct{}

func (UnimplementedStateServiceServer) GetState(context.Context, *Empty) (*State, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}
func (UnimplementedStateServiceServer) SetState(context.Context, *State) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetState not implemented")
}
func (UnimplementedStateServiceServer) mustEmbedUnimplementedStateServiceServer() {}
func (UnimplementedStateServiceServer) testEmbeddedByValue()                      {}

// UnsafeStateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StateServiceServer will
// result in compilation errors.
type UnsafeStateServiceServer interface {
	mustEmbedUnimplementedStateServiceServer()
}

func RegisterStateServiceServer(s grpc.ServiceRegistrar, srv StateServiceServer) {
	// If the following call pancis, it indicates UnimplementedStateServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StateService_ServiceDesc, srv)
}

func _StateService_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateServiceServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StateService_GetState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateServiceServer).GetState(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StateService_SetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateServiceServer).SetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StateService_SetState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateServiceServer).SetState(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

// StateService_ServiceDesc is the grpc.ServiceDesc for StateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "myuber.StateService",
	HandlerType: (*StateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetState",
			Handler:    _StateService_GetState_Handler,
		},
		{
			MethodName: "SetState",
			Handler:    _StateService_SetState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/state.proto",
}
