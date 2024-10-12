// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: pkg/proto/labyrinth.proto

package lrpb

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
	KNN_GetLabyrinthInfo_FullMethodName = "/labyrinth.KNN/GetLabyrinthInfo"
	KNN_GetPlayerStatus_FullMethodName  = "/labyrinth.KNN/GetPlayerStatus"
	KNN_RegisterMove_FullMethodName     = "/labyrinth.KNN/RegisterMove"
	KNN_Revelio_FullMethodName          = "/labyrinth.KNN/Revelio"
	KNN_Bombarda_FullMethodName         = "/labyrinth.KNN/Bombarda"
)

// KNNClient is the client API for KNN service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KNNClient interface {
	GetLabyrinthInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*LabyrinthInfo, error)
	GetPlayerStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlayerStatus, error)
	RegisterMove(ctx context.Context, in *Move, opts ...grpc.CallOption) (*MoveResponse, error)
	Revelio(ctx context.Context, in *RevelioRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[RevelioResponse], error)
	Bombarda(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[BombardaRequest, BombardaResponse], error)
}

type kNNClient struct {
	cc grpc.ClientConnInterface
}

func NewKNNClient(cc grpc.ClientConnInterface) KNNClient {
	return &kNNClient{cc}
}

func (c *kNNClient) GetLabyrinthInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*LabyrinthInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LabyrinthInfo)
	err := c.cc.Invoke(ctx, KNN_GetLabyrinthInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kNNClient) GetPlayerStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlayerStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlayerStatus)
	err := c.cc.Invoke(ctx, KNN_GetPlayerStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kNNClient) RegisterMove(ctx context.Context, in *Move, opts ...grpc.CallOption) (*MoveResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MoveResponse)
	err := c.cc.Invoke(ctx, KNN_RegisterMove_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kNNClient) Revelio(ctx context.Context, in *RevelioRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[RevelioResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &KNN_ServiceDesc.Streams[0], KNN_Revelio_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[RevelioRequest, RevelioResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type KNN_RevelioClient = grpc.ServerStreamingClient[RevelioResponse]

func (c *kNNClient) Bombarda(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[BombardaRequest, BombardaResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &KNN_ServiceDesc.Streams[1], KNN_Bombarda_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[BombardaRequest, BombardaResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type KNN_BombardaClient = grpc.ClientStreamingClient[BombardaRequest, BombardaResponse]

// KNNServer is the server API for KNN service.
// All implementations must embed UnimplementedKNNServer
// for forward compatibility.
type KNNServer interface {
	GetLabyrinthInfo(context.Context, *Empty) (*LabyrinthInfo, error)
	GetPlayerStatus(context.Context, *Empty) (*PlayerStatus, error)
	RegisterMove(context.Context, *Move) (*MoveResponse, error)
	Revelio(*RevelioRequest, grpc.ServerStreamingServer[RevelioResponse]) error
	Bombarda(grpc.ClientStreamingServer[BombardaRequest, BombardaResponse]) error
	mustEmbedUnimplementedKNNServer()
}

// UnimplementedKNNServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKNNServer struct{}

func (UnimplementedKNNServer) GetLabyrinthInfo(context.Context, *Empty) (*LabyrinthInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabyrinthInfo not implemented")
}
func (UnimplementedKNNServer) GetPlayerStatus(context.Context, *Empty) (*PlayerStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerStatus not implemented")
}
func (UnimplementedKNNServer) RegisterMove(context.Context, *Move) (*MoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterMove not implemented")
}
func (UnimplementedKNNServer) Revelio(*RevelioRequest, grpc.ServerStreamingServer[RevelioResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Revelio not implemented")
}
func (UnimplementedKNNServer) Bombarda(grpc.ClientStreamingServer[BombardaRequest, BombardaResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Bombarda not implemented")
}
func (UnimplementedKNNServer) mustEmbedUnimplementedKNNServer() {}
func (UnimplementedKNNServer) testEmbeddedByValue()             {}

// UnsafeKNNServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KNNServer will
// result in compilation errors.
type UnsafeKNNServer interface {
	mustEmbedUnimplementedKNNServer()
}

func RegisterKNNServer(s grpc.ServiceRegistrar, srv KNNServer) {
	// If the following call pancis, it indicates UnimplementedKNNServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KNN_ServiceDesc, srv)
}

func _KNN_GetLabyrinthInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KNNServer).GetLabyrinthInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KNN_GetLabyrinthInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KNNServer).GetLabyrinthInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KNN_GetPlayerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KNNServer).GetPlayerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KNN_GetPlayerStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KNNServer).GetPlayerStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KNN_RegisterMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Move)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KNNServer).RegisterMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KNN_RegisterMove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KNNServer).RegisterMove(ctx, req.(*Move))
	}
	return interceptor(ctx, in, info, handler)
}

func _KNN_Revelio_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RevelioRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KNNServer).Revelio(m, &grpc.GenericServerStream[RevelioRequest, RevelioResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type KNN_RevelioServer = grpc.ServerStreamingServer[RevelioResponse]

func _KNN_Bombarda_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KNNServer).Bombarda(&grpc.GenericServerStream[BombardaRequest, BombardaResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type KNN_BombardaServer = grpc.ClientStreamingServer[BombardaRequest, BombardaResponse]

// KNN_ServiceDesc is the grpc.ServiceDesc for KNN service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KNN_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "labyrinth.KNN",
	HandlerType: (*KNNServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLabyrinthInfo",
			Handler:    _KNN_GetLabyrinthInfo_Handler,
		},
		{
			MethodName: "GetPlayerStatus",
			Handler:    _KNN_GetPlayerStatus_Handler,
		},
		{
			MethodName: "RegisterMove",
			Handler:    _KNN_RegisterMove_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Revelio",
			Handler:       _KNN_Revelio_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Bombarda",
			Handler:       _KNN_Bombarda_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/proto/labyrinth.proto",
}
