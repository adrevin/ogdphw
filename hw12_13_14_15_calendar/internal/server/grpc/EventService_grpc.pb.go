// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/EventService.proto

package internalgrpc

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EvensClient is the client API for Evens service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EvensClient interface {
	CreateEvent(ctx context.Context, in *NewEventRequest, opts ...grpc.CallOption) (*EventIdResponse, error)
	UpdateEvent(ctx context.Context, in *ChangeEventRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteEvent(ctx context.Context, in *EventIdRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DayEvens(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*EventsResponse, error)
	WeekEvens(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*EventsResponse, error)
	MonthEvens(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*EventsResponse, error)
}

type evensClient struct {
	cc grpc.ClientConnInterface
}

func NewEvensClient(cc grpc.ClientConnInterface) EvensClient {
	return &evensClient{cc}
}

func (c *evensClient) CreateEvent(ctx context.Context, in *NewEventRequest, opts ...grpc.CallOption) (*EventIdResponse, error) {
	out := new(EventIdResponse)
	err := c.cc.Invoke(ctx, "/event.Evens/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evensClient) UpdateEvent(ctx context.Context, in *ChangeEventRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/event.Evens/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evensClient) DeleteEvent(ctx context.Context, in *EventIdRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/event.Evens/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evensClient) DayEvens(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/event.Evens/DayEvens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evensClient) WeekEvens(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/event.Evens/WeekEvens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evensClient) MonthEvens(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/event.Evens/MonthEvens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EvensServer is the server API for Evens service.
// All implementations must embed UnimplementedEvensServer
// for forward compatibility
type EvensServer interface {
	CreateEvent(context.Context, *NewEventRequest) (*EventIdResponse, error)
	UpdateEvent(context.Context, *ChangeEventRequest) (*empty.Empty, error)
	DeleteEvent(context.Context, *EventIdRequest) (*empty.Empty, error)
	DayEvens(context.Context, *TimeRequest) (*EventsResponse, error)
	WeekEvens(context.Context, *TimeRequest) (*EventsResponse, error)
	MonthEvens(context.Context, *TimeRequest) (*EventsResponse, error)
	mustEmbedUnimplementedEvensServer()
}

// UnimplementedEvensServer must be embedded to have forward compatible implementations.
type UnimplementedEvensServer struct {
}

func (UnimplementedEvensServer) CreateEvent(context.Context, *NewEventRequest) (*EventIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedEvensServer) UpdateEvent(context.Context, *ChangeEventRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (UnimplementedEvensServer) DeleteEvent(context.Context, *EventIdRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedEvensServer) DayEvens(context.Context, *TimeRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DayEvens not implemented")
}
func (UnimplementedEvensServer) WeekEvens(context.Context, *TimeRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WeekEvens not implemented")
}
func (UnimplementedEvensServer) MonthEvens(context.Context, *TimeRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MonthEvens not implemented")
}
func (UnimplementedEvensServer) mustEmbedUnimplementedEvensServer() {}

// UnsafeEvensServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EvensServer will
// result in compilation errors.
type UnsafeEvensServer interface {
	mustEmbedUnimplementedEvensServer()
}

func RegisterEvensServer(s grpc.ServiceRegistrar, srv EvensServer) {
	s.RegisterService(&Evens_ServiceDesc, srv)
}

func _Evens_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvensServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.Evens/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvensServer).CreateEvent(ctx, req.(*NewEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Evens_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvensServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.Evens/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvensServer).UpdateEvent(ctx, req.(*ChangeEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Evens_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvensServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.Evens/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvensServer).DeleteEvent(ctx, req.(*EventIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Evens_DayEvens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvensServer).DayEvens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.Evens/DayEvens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvensServer).DayEvens(ctx, req.(*TimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Evens_WeekEvens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvensServer).WeekEvens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.Evens/WeekEvens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvensServer).WeekEvens(ctx, req.(*TimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Evens_MonthEvens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvensServer).MonthEvens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.Evens/MonthEvens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvensServer).MonthEvens(ctx, req.(*TimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Evens_ServiceDesc is the grpc.ServiceDesc for Evens service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Evens_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "event.Evens",
	HandlerType: (*EvensServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _Evens_CreateEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _Evens_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _Evens_DeleteEvent_Handler,
		},
		{
			MethodName: "DayEvens",
			Handler:    _Evens_DayEvens_Handler,
		},
		{
			MethodName: "WeekEvens",
			Handler:    _Evens_WeekEvens_Handler,
		},
		{
			MethodName: "MonthEvens",
			Handler:    _Evens_MonthEvens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/EventService.proto",
}