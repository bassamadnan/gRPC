// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: pkg/proto/docs.proto

package dpb

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
	DocsService_SendMessage_FullMethodName = "/docs.DocsService/SendMessage"
)

// DocsServiceClient is the client API for DocsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DocsServiceClient interface {
	SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageResponse, error)
}

type docsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDocsServiceClient(cc grpc.ClientConnInterface) DocsServiceClient {
	return &docsServiceClient{cc}
}

func (c *docsServiceClient) SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, DocsService_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocsServiceServer is the server API for DocsService service.
// All implementations must embed UnimplementedDocsServiceServer
// for forward compatibility.
type DocsServiceServer interface {
	SendMessage(context.Context, *Message) (*MessageResponse, error)
	mustEmbedUnimplementedDocsServiceServer()
}

// UnimplementedDocsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDocsServiceServer struct{}

func (UnimplementedDocsServiceServer) SendMessage(context.Context, *Message) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedDocsServiceServer) mustEmbedUnimplementedDocsServiceServer() {}
func (UnimplementedDocsServiceServer) testEmbeddedByValue()                     {}

// UnsafeDocsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocsServiceServer will
// result in compilation errors.
type UnsafeDocsServiceServer interface {
	mustEmbedUnimplementedDocsServiceServer()
}

func RegisterDocsServiceServer(s grpc.ServiceRegistrar, srv DocsServiceServer) {
	// If the following call pancis, it indicates UnimplementedDocsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DocsService_ServiceDesc, srv)
}

func _DocsService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocsServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocsService_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocsServiceServer).SendMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// DocsService_ServiceDesc is the grpc.ServiceDesc for DocsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DocsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "docs.DocsService",
	HandlerType: (*DocsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _DocsService_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/docs.proto",
}
