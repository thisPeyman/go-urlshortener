// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: id_generator_service.proto

package api

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
	IDGeneratorService_GenerateID_FullMethodName = "/api.IDGeneratorService/GenerateID"
)

// IDGeneratorServiceClient is the client API for IDGeneratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IDGeneratorServiceClient interface {
	GenerateID(ctx context.Context, in *GenerateIDRequest, opts ...grpc.CallOption) (*GenerateIDResponse, error)
}

type iDGeneratorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIDGeneratorServiceClient(cc grpc.ClientConnInterface) IDGeneratorServiceClient {
	return &iDGeneratorServiceClient{cc}
}

func (c *iDGeneratorServiceClient) GenerateID(ctx context.Context, in *GenerateIDRequest, opts ...grpc.CallOption) (*GenerateIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateIDResponse)
	err := c.cc.Invoke(ctx, IDGeneratorService_GenerateID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IDGeneratorServiceServer is the server API for IDGeneratorService service.
// All implementations must embed UnimplementedIDGeneratorServiceServer
// for forward compatibility.
type IDGeneratorServiceServer interface {
	GenerateID(context.Context, *GenerateIDRequest) (*GenerateIDResponse, error)
	mustEmbedUnimplementedIDGeneratorServiceServer()
}

// UnimplementedIDGeneratorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedIDGeneratorServiceServer struct{}

func (UnimplementedIDGeneratorServiceServer) GenerateID(context.Context, *GenerateIDRequest) (*GenerateIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateID not implemented")
}
func (UnimplementedIDGeneratorServiceServer) mustEmbedUnimplementedIDGeneratorServiceServer() {}
func (UnimplementedIDGeneratorServiceServer) testEmbeddedByValue()                            {}

// UnsafeIDGeneratorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IDGeneratorServiceServer will
// result in compilation errors.
type UnsafeIDGeneratorServiceServer interface {
	mustEmbedUnimplementedIDGeneratorServiceServer()
}

func RegisterIDGeneratorServiceServer(s grpc.ServiceRegistrar, srv IDGeneratorServiceServer) {
	// If the following call pancis, it indicates UnimplementedIDGeneratorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&IDGeneratorService_ServiceDesc, srv)
}

func _IDGeneratorService_GenerateID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IDGeneratorServiceServer).GenerateID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IDGeneratorService_GenerateID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IDGeneratorServiceServer).GenerateID(ctx, req.(*GenerateIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IDGeneratorService_ServiceDesc is the grpc.ServiceDesc for IDGeneratorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IDGeneratorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.IDGeneratorService",
	HandlerType: (*IDGeneratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateID",
			Handler:    _IDGeneratorService_GenerateID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "id_generator_service.proto",
}
