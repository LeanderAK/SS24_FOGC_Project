// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: fog.proto

package fogpb

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

// SensorServiceClient is the client API for SensorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SensorServiceClient interface {
	// StreamData will be called by the EdgeService
	// to initiate a stream of sensor data.
	StreamData(ctx context.Context, in *StreamDataRequest, opts ...grpc.CallOption) (SensorService_StreamDataClient, error)
}

type sensorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSensorServiceClient(cc grpc.ClientConnInterface) SensorServiceClient {
	return &sensorServiceClient{cc}
}

func (c *sensorServiceClient) StreamData(ctx context.Context, in *StreamDataRequest, opts ...grpc.CallOption) (SensorService_StreamDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &SensorService_ServiceDesc.Streams[0], "/fogpb.SensorService/StreamData", opts...)
	if err != nil {
		return nil, err
	}
	x := &sensorServiceStreamDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SensorService_StreamDataClient interface {
	Recv() (*StreamDataResponse, error)
	grpc.ClientStream
}

type sensorServiceStreamDataClient struct {
	grpc.ClientStream
}

func (x *sensorServiceStreamDataClient) Recv() (*StreamDataResponse, error) {
	m := new(StreamDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SensorServiceServer is the server API for SensorService service.
// All implementations must embed UnimplementedSensorServiceServer
// for forward compatibility
type SensorServiceServer interface {
	// StreamData will be called by the EdgeService
	// to initiate a stream of sensor data.
	StreamData(*StreamDataRequest, SensorService_StreamDataServer) error
	mustEmbedUnimplementedSensorServiceServer()
}

// UnimplementedSensorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSensorServiceServer struct {
}

func (UnimplementedSensorServiceServer) StreamData(*StreamDataRequest, SensorService_StreamDataServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamData not implemented")
}
func (UnimplementedSensorServiceServer) mustEmbedUnimplementedSensorServiceServer() {}

// UnsafeSensorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SensorServiceServer will
// result in compilation errors.
type UnsafeSensorServiceServer interface {
	mustEmbedUnimplementedSensorServiceServer()
}

func RegisterSensorServiceServer(s grpc.ServiceRegistrar, srv SensorServiceServer) {
	s.RegisterService(&SensorService_ServiceDesc, srv)
}

func _SensorService_StreamData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamDataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SensorServiceServer).StreamData(m, &sensorServiceStreamDataServer{stream})
}

type SensorService_StreamDataServer interface {
	Send(*StreamDataResponse) error
	grpc.ServerStream
}

type sensorServiceStreamDataServer struct {
	grpc.ServerStream
}

func (x *sensorServiceStreamDataServer) Send(m *StreamDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SensorService_ServiceDesc is the grpc.ServiceDesc for SensorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SensorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fogpb.SensorService",
	HandlerType: (*SensorServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamData",
			Handler:       _SensorService_StreamData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "fog.proto",
}

// EdgeServiceClient is the client API for EdgeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EdgeServiceClient interface {
	// UpdatePosition will be called by the SensorService
	// to update the position of the edge device.
	// This endpoint returns nothing,
	// no error means position was updated successfully.
	UpdatePosition(ctx context.Context, in *UpdatePositionRequest, opts ...grpc.CallOption) (*UpdatePositionResponse, error)
}

type edgeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEdgeServiceClient(cc grpc.ClientConnInterface) EdgeServiceClient {
	return &edgeServiceClient{cc}
}

func (c *edgeServiceClient) UpdatePosition(ctx context.Context, in *UpdatePositionRequest, opts ...grpc.CallOption) (*UpdatePositionResponse, error) {
	out := new(UpdatePositionResponse)
	err := c.cc.Invoke(ctx, "/fogpb.EdgeService/UpdatePosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EdgeServiceServer is the server API for EdgeService service.
// All implementations must embed UnimplementedEdgeServiceServer
// for forward compatibility
type EdgeServiceServer interface {
	// UpdatePosition will be called by the SensorService
	// to update the position of the edge device.
	// This endpoint returns nothing,
	// no error means position was updated successfully.
	UpdatePosition(context.Context, *UpdatePositionRequest) (*UpdatePositionResponse, error)
	mustEmbedUnimplementedEdgeServiceServer()
}

// UnimplementedEdgeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEdgeServiceServer struct {
}

func (UnimplementedEdgeServiceServer) UpdatePosition(context.Context, *UpdatePositionRequest) (*UpdatePositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePosition not implemented")
}
func (UnimplementedEdgeServiceServer) mustEmbedUnimplementedEdgeServiceServer() {}

// UnsafeEdgeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EdgeServiceServer will
// result in compilation errors.
type UnsafeEdgeServiceServer interface {
	mustEmbedUnimplementedEdgeServiceServer()
}

func RegisterEdgeServiceServer(s grpc.ServiceRegistrar, srv EdgeServiceServer) {
	s.RegisterService(&EdgeService_ServiceDesc, srv)
}

func _EdgeService_UpdatePosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdgeServiceServer).UpdatePosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fogpb.EdgeService/UpdatePosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdgeServiceServer).UpdatePosition(ctx, req.(*UpdatePositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EdgeService_ServiceDesc is the grpc.ServiceDesc for EdgeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EdgeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fogpb.EdgeService",
	HandlerType: (*EdgeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdatePosition",
			Handler:    _EdgeService_UpdatePosition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fog.proto",
}

// CloudServiceClient is the client API for CloudService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudServiceClient interface {
	// ProcessData will be called by the EdgeService
	// to process aggregated sensor data.
	// This endpoint returns nothing,
	// no error means data was received successfully.
	// Results will be returned to the UpdatePosition endpoint
	// in the EdgeService.
	ProcessData(ctx context.Context, in *ProcessDataRequest, opts ...grpc.CallOption) (*ProcessDataResponse, error)
}

type cloudServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudServiceClient(cc grpc.ClientConnInterface) CloudServiceClient {
	return &cloudServiceClient{cc}
}

func (c *cloudServiceClient) ProcessData(ctx context.Context, in *ProcessDataRequest, opts ...grpc.CallOption) (*ProcessDataResponse, error) {
	out := new(ProcessDataResponse)
	err := c.cc.Invoke(ctx, "/fogpb.CloudService/ProcessData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudServiceServer is the server API for CloudService service.
// All implementations must embed UnimplementedCloudServiceServer
// for forward compatibility
type CloudServiceServer interface {
	// ProcessData will be called by the EdgeService
	// to process aggregated sensor data.
	// This endpoint returns nothing,
	// no error means data was received successfully.
	// Results will be returned to the UpdatePosition endpoint
	// in the EdgeService.
	ProcessData(context.Context, *ProcessDataRequest) (*ProcessDataResponse, error)
	mustEmbedUnimplementedCloudServiceServer()
}

// UnimplementedCloudServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCloudServiceServer struct {
}

func (UnimplementedCloudServiceServer) ProcessData(context.Context, *ProcessDataRequest) (*ProcessDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessData not implemented")
}
func (UnimplementedCloudServiceServer) mustEmbedUnimplementedCloudServiceServer() {}

// UnsafeCloudServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudServiceServer will
// result in compilation errors.
type UnsafeCloudServiceServer interface {
	mustEmbedUnimplementedCloudServiceServer()
}

func RegisterCloudServiceServer(s grpc.ServiceRegistrar, srv CloudServiceServer) {
	s.RegisterService(&CloudService_ServiceDesc, srv)
}

func _CloudService_ProcessData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).ProcessData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fogpb.CloudService/ProcessData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).ProcessData(ctx, req.(*ProcessDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudService_ServiceDesc is the grpc.ServiceDesc for CloudService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fogpb.CloudService",
	HandlerType: (*CloudServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessData",
			Handler:    _CloudService_ProcessData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fog.proto",
}
