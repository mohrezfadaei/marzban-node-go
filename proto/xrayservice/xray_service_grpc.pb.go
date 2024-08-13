// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.57.0
// - protoc             v3.12.4
// source: proto/xray_service.proto

package xrayservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.57.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	XrayService_Start_FullMethodName            = "/xrayservice.XrayService/Start"
	XrayService_Stop_FullMethodName             = "/xrayservice.XrayService/Stop"
	XrayService_Restart_FullMethodName          = "/xrayservice.XrayService/Restart"
	XrayService_FetchXrayVersion_FullMethodName = "/xrayservice.XrayService/FetchXrayVersion"
	XrayService_FetchLogs_FullMethodName        = "/xrayservice.XrayService/FetchLogs"
)

// XrayServiceClient is the client API for XrayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service definition for XrayService
type XrayServiceClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
	Restart(ctx context.Context, in *RestartRequest, opts ...grpc.CallOption) (*RestartResponse, error)
	FetchXrayVersion(ctx context.Context, in *FetchXrayVersionRequest, opts ...grpc.CallOption) (*FetchXrayVersionResponse, error)
	FetchLogs(ctx context.Context, in *FetchLogsRequest, opts ...grpc.CallOption) (XrayService_FetchLogsClient, error)
}

type xrayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewXrayServiceClient(cc grpc.ClientConnInterface) XrayServiceClient {
	return &xrayServiceClient{cc}
}

func (c *xrayServiceClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, XrayService_Start_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xrayServiceClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, XrayService_Stop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xrayServiceClient) Restart(ctx context.Context, in *RestartRequest, opts ...grpc.CallOption) (*RestartResponse, error) {
	out := new(RestartResponse)
	err := c.cc.Invoke(ctx, XrayService_Restart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xrayServiceClient) FetchXrayVersion(ctx context.Context, in *FetchXrayVersionRequest, opts ...grpc.CallOption) (*FetchXrayVersionResponse, error) {
	out := new(FetchXrayVersionResponse)
	err := c.cc.Invoke(ctx, XrayService_FetchXrayVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xrayServiceClient) FetchLogs(ctx context.Context, in *FetchLogsRequest, opts ...grpc.CallOption) (XrayService_FetchLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &XrayService_ServiceDesc.Streams[0], XrayService_FetchLogs_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &xrayServiceFetchLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type XrayService_FetchLogsClient interface {
	Recv() (*LogMessage, error)
	grpc.ClientStream
}

type xrayServiceFetchLogsClient struct {
	grpc.ClientStream
}

func (x *xrayServiceFetchLogsClient) Recv() (*LogMessage, error) {
	m := new(LogMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// XrayServiceServer is the server API for XrayService service.
// All implementations must embed UnimplementedXrayServiceServer
// for forward compatibility
type XrayServiceServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Stop(context.Context, *StopRequest) (*StopResponse, error)
	Restart(context.Context, *RestartRequest) (*RestartResponse, error)
	FetchXrayVersion(context.Context, *FetchXrayVersionRequest) (*FetchXrayVersionResponse, error)
	FetchLogs(*FetchLogsRequest, XrayService_FetchLogsServer) error
	mustEmbedUnimplementedXrayServiceServer()
}

// UnimplementedXrayServiceServer must be embedded to have forward compatible implementations
type UnimplementedXrayServiceServer struct{}

func (UnimplementedXrayServiceServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedXrayServiceServer) Stop(context.Context, *StopRequest) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedXrayServiceServer) Restart(context.Context, *RestartRequest) (*RestartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedXrayServiceServer) FetchXrayVersion(context.Context, *FetchXrayVersionRequest) (*FetchXrayVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchXrayVersion not implemented")
}
func (UnimplementedXrayServiceServer) FetchLogs(*FetchLogsRequest, XrayService_FetchLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchLogs not implemented")
}
func (UnimplementedXrayServiceServer) mustEmbedUnimplementedXrayServiceServer() {}

// UnsafeXrayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to XrayServiceServer will
// result in compilation errors.
type UnsafeXrayServiceServer interface {
	mustEmbedUnimplementedXrayServiceServer()
}

func RegisterXrayServiceServer(s grpc.ServiceRegistrar, srv XrayServiceServer) {
	s.RegisterService(&XrayService_ServiceDesc, srv)
}

func _XrayService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XrayServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XrayService_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XrayServiceServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XrayService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XrayServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XrayService_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XrayServiceServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XrayService_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XrayServiceServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XrayService_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XrayServiceServer).Restart(ctx, req.(*RestartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XrayService_FetchXrayVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchXrayVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XrayServiceServer).FetchXrayVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XrayService_FetchXrayVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XrayServiceServer).FetchXrayVersion(ctx, req.(*FetchXrayVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XrayService_FetchLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FetchLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(XrayServiceServer).FetchLogs(m, &xrayServiceFetchLogsServer{stream})
}

type XrayService_FetchLogsServer interface {
	Send(*LogMessage) error
	grpc.ServerStream
}

type xrayServiceFetchLogsServer struct {
	grpc.ServerStream
}

func (x *xrayServiceFetchLogsServer) Send(m *LogMessage) error {
	return x.ServerStream.SendMsg(m)
}

// XrayService_ServiceDesc is the grpc.ServiceDesc for XrayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var XrayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xrayservice.XrayService",
	HandlerType: (*XrayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _XrayService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _XrayService_Stop_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _XrayService_Restart_Handler,
		},
		{
			MethodName: "FetchXrayVersion",
			Handler:    _XrayService_FetchXrayVersion_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchLogs",
			Handler:       _XrayService_FetchLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/xray_service.proto",
}