// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: v1/api/service.proto

package api

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

// HelloWorldClient is the client API for HelloWorld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloWorldClient interface {
	// This is just a simple POST request example. The
	// request parameters are all expected to be included
	// in the body of the request.
	//
	// POST http://localhost:8081/greet
	Greet(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Goodbye, error)
	// This request is a example of path parameters.
	// Path parameters are evaluated before the body
	// option, so any fields not included in the path
	// are expected in the body instead.
	//
	// POST http://localhost:8081/greetother/en
	GreetOther(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Goodbye, error)
	// This is an example of a GET request that uses query
	// parameters. GET requests cannot have a request body,
	// and must provide all parameters via the URL.
	// The codegen evaluates the path parameters first. Any
	// other parameters are expected as query params.
	//
	// GET http://localhost:8081/v1/greet/hello?language=en
	GreetOther2(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Goodbye, error)
	// This is another simple GET example.
	//
	// GET http://localhost:8081/v1/greet/reverse/helloworld
	GreetReverse(ctx context.Context, in *Goodbye, opts ...grpc.CallOption) (*Goodbye, error)
}

type helloWorldClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloWorldClient(cc grpc.ClientConnInterface) HelloWorldClient {
	return &helloWorldClient{cc}
}

func (c *helloWorldClient) Greet(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Goodbye, error) {
	out := new(Goodbye)
	err := c.cc.Invoke(ctx, "/v1.api.HelloWorld/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloWorldClient) GreetOther(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Goodbye, error) {
	out := new(Goodbye)
	err := c.cc.Invoke(ctx, "/v1.api.HelloWorld/GreetOther", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloWorldClient) GreetOther2(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Goodbye, error) {
	out := new(Goodbye)
	err := c.cc.Invoke(ctx, "/v1.api.HelloWorld/GreetOther2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloWorldClient) GreetReverse(ctx context.Context, in *Goodbye, opts ...grpc.CallOption) (*Goodbye, error) {
	out := new(Goodbye)
	err := c.cc.Invoke(ctx, "/v1.api.HelloWorld/GreetReverse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloWorldServer is the server API for HelloWorld service.
// All implementations should embed UnimplementedHelloWorldServer
// for forward compatibility
type HelloWorldServer interface {
	// This is just a simple POST request example. The
	// request parameters are all expected to be included
	// in the body of the request.
	//
	// POST http://localhost:8081/greet
	Greet(context.Context, *Hello) (*Goodbye, error)
	// This request is a example of path parameters.
	// Path parameters are evaluated before the body
	// option, so any fields not included in the path
	// are expected in the body instead.
	//
	// POST http://localhost:8081/greetother/en
	GreetOther(context.Context, *Hello) (*Goodbye, error)
	// This is an example of a GET request that uses query
	// parameters. GET requests cannot have a request body,
	// and must provide all parameters via the URL.
	// The codegen evaluates the path parameters first. Any
	// other parameters are expected as query params.
	//
	// GET http://localhost:8081/v1/greet/hello?language=en
	GreetOther2(context.Context, *Hello) (*Goodbye, error)
	// This is another simple GET example.
	//
	// GET http://localhost:8081/v1/greet/reverse/helloworld
	GreetReverse(context.Context, *Goodbye) (*Goodbye, error)
}

// UnimplementedHelloWorldServer should be embedded to have forward compatible implementations.
type UnimplementedHelloWorldServer struct {
}

func (UnimplementedHelloWorldServer) Greet(context.Context, *Hello) (*Goodbye, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (UnimplementedHelloWorldServer) GreetOther(context.Context, *Hello) (*Goodbye, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreetOther not implemented")
}
func (UnimplementedHelloWorldServer) GreetOther2(context.Context, *Hello) (*Goodbye, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreetOther2 not implemented")
}
func (UnimplementedHelloWorldServer) GreetReverse(context.Context, *Goodbye) (*Goodbye, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreetReverse not implemented")
}

// UnsafeHelloWorldServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloWorldServer will
// result in compilation errors.
type UnsafeHelloWorldServer interface {
	mustEmbedUnimplementedHelloWorldServer()
}

func RegisterHelloWorldServer(s grpc.ServiceRegistrar, srv HelloWorldServer) {
	s.RegisterService(&HelloWorld_ServiceDesc, srv)
}

func _HelloWorld_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hello)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.api.HelloWorld/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServer).Greet(ctx, req.(*Hello))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloWorld_GreetOther_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hello)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServer).GreetOther(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.api.HelloWorld/GreetOther",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServer).GreetOther(ctx, req.(*Hello))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloWorld_GreetOther2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hello)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServer).GreetOther2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.api.HelloWorld/GreetOther2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServer).GreetOther2(ctx, req.(*Hello))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloWorld_GreetReverse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Goodbye)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServer).GreetReverse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.api.HelloWorld/GreetReverse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServer).GreetReverse(ctx, req.(*Goodbye))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloWorld_ServiceDesc is the grpc.ServiceDesc for HelloWorld service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloWorld_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.api.HelloWorld",
	HandlerType: (*HelloWorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _HelloWorld_Greet_Handler,
		},
		{
			MethodName: "GreetOther",
			Handler:    _HelloWorld_GreetOther_Handler,
		},
		{
			MethodName: "GreetOther2",
			Handler:    _HelloWorld_GreetOther2_Handler,
		},
		{
			MethodName: "GreetReverse",
			Handler:    _HelloWorld_GreetReverse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/api/service.proto",
}