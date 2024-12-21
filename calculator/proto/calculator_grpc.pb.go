// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: calculator.proto

package proto

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
	CalculatorService_Add_FullMethodName       = "/calculator.CalculatorService/Add"
	CalculatorService_AddSeries_FullMethodName = "/calculator.CalculatorService/AddSeries"
	CalculatorService_Avg_FullMethodName       = "/calculator.CalculatorService/Avg"
	CalculatorService_Primes_FullMethodName    = "/calculator.CalculatorService/Primes"
)

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The Calculator service service definition
type CalculatorServiceClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	AddSeries(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[AddSeriesRequest, AddResponse], error)
	Avg(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[AvgRequest, AvgResponse], error)
	Primes(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PrimeResponse], error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, CalculatorService_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) AddSeries(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[AddSeriesRequest, AddResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[0], CalculatorService_AddSeries_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[AddSeriesRequest, AddResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_AddSeriesClient = grpc.BidiStreamingClient[AddSeriesRequest, AddResponse]

func (c *calculatorServiceClient) Avg(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[AvgRequest, AvgResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[1], CalculatorService_Avg_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[AvgRequest, AvgResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_AvgClient = grpc.ClientStreamingClient[AvgRequest, AvgResponse]

func (c *calculatorServiceClient) Primes(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PrimeResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[2], CalculatorService_Primes_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PrimeRequest, PrimeResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_PrimesClient = grpc.ServerStreamingClient[PrimeResponse]

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations must embed UnimplementedCalculatorServiceServer
// for forward compatibility.
//
// The Calculator service service definition
type CalculatorServiceServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	AddSeries(grpc.BidiStreamingServer[AddSeriesRequest, AddResponse]) error
	Avg(grpc.ClientStreamingServer[AvgRequest, AvgResponse]) error
	Primes(*PrimeRequest, grpc.ServerStreamingServer[PrimeResponse]) error
	mustEmbedUnimplementedCalculatorServiceServer()
}

// UnimplementedCalculatorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCalculatorServiceServer struct{}

func (UnimplementedCalculatorServiceServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCalculatorServiceServer) AddSeries(grpc.BidiStreamingServer[AddSeriesRequest, AddResponse]) error {
	return status.Errorf(codes.Unimplemented, "method AddSeries not implemented")
}
func (UnimplementedCalculatorServiceServer) Avg(grpc.ClientStreamingServer[AvgRequest, AvgResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Avg not implemented")
}
func (UnimplementedCalculatorServiceServer) Primes(*PrimeRequest, grpc.ServerStreamingServer[PrimeResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Primes not implemented")
}
func (UnimplementedCalculatorServiceServer) mustEmbedUnimplementedCalculatorServiceServer() {}
func (UnimplementedCalculatorServiceServer) testEmbeddedByValue()                           {}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	// If the following call pancis, it indicates UnimplementedCalculatorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalculatorService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_AddSeries_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).AddSeries(&grpc.GenericServerStream[AddSeriesRequest, AddResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_AddSeriesServer = grpc.BidiStreamingServer[AddSeriesRequest, AddResponse]

func _CalculatorService_Avg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).Avg(&grpc.GenericServerStream[AvgRequest, AvgResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_AvgServer = grpc.ClientStreamingServer[AvgRequest, AvgResponse]

func _CalculatorService_Primes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).Primes(m, &grpc.GenericServerStream[PrimeRequest, PrimeResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_PrimesServer = grpc.ServerStreamingServer[PrimeResponse]

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CalculatorService_Add_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddSeries",
			Handler:       _CalculatorService_AddSeries_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Avg",
			Handler:       _CalculatorService_Avg_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Primes",
			Handler:       _CalculatorService_Primes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
