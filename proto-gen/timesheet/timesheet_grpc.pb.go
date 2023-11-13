// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: proto/timesheet/timesheet.proto

package timesheet

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

const (
	TimesheetService_CreateTimesheet_FullMethodName                    = "/timesheet_app.TimesheetService/CreateTimesheet"
	TimesheetService_GetTimesheetsByContractor_FullMethodName          = "/timesheet_app.TimesheetService/GetTimesheetsByContractor"
	TimesheetService_GetTimesheetsByContractorAndClient_FullMethodName = "/timesheet_app.TimesheetService/GetTimesheetsByContractorAndClient"
)

// TimesheetServiceClient is the client API for TimesheetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimesheetServiceClient interface {
	CreateTimesheet(ctx context.Context, in *CreateTimesheetRequest, opts ...grpc.CallOption) (*Timesheet, error)
	GetTimesheetsByContractor(ctx context.Context, in *GetTimesheetsByContractorRequest, opts ...grpc.CallOption) (*GetTimesheetsResponse, error)
	GetTimesheetsByContractorAndClient(ctx context.Context, in *GetTimesheetsByContractorAndClientRequest, opts ...grpc.CallOption) (*GetTimesheetsResponse, error)
}

type timesheetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimesheetServiceClient(cc grpc.ClientConnInterface) TimesheetServiceClient {
	return &timesheetServiceClient{cc}
}

func (c *timesheetServiceClient) CreateTimesheet(ctx context.Context, in *CreateTimesheetRequest, opts ...grpc.CallOption) (*Timesheet, error) {
	out := new(Timesheet)
	err := c.cc.Invoke(ctx, TimesheetService_CreateTimesheet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timesheetServiceClient) GetTimesheetsByContractor(ctx context.Context, in *GetTimesheetsByContractorRequest, opts ...grpc.CallOption) (*GetTimesheetsResponse, error) {
	out := new(GetTimesheetsResponse)
	err := c.cc.Invoke(ctx, TimesheetService_GetTimesheetsByContractor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timesheetServiceClient) GetTimesheetsByContractorAndClient(ctx context.Context, in *GetTimesheetsByContractorAndClientRequest, opts ...grpc.CallOption) (*GetTimesheetsResponse, error) {
	out := new(GetTimesheetsResponse)
	err := c.cc.Invoke(ctx, TimesheetService_GetTimesheetsByContractorAndClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimesheetServiceServer is the server API for TimesheetService service.
// All implementations must embed UnimplementedTimesheetServiceServer
// for forward compatibility
type TimesheetServiceServer interface {
	CreateTimesheet(context.Context, *CreateTimesheetRequest) (*Timesheet, error)
	GetTimesheetsByContractor(context.Context, *GetTimesheetsByContractorRequest) (*GetTimesheetsResponse, error)
	GetTimesheetsByContractorAndClient(context.Context, *GetTimesheetsByContractorAndClientRequest) (*GetTimesheetsResponse, error)
	mustEmbedUnimplementedTimesheetServiceServer()
}

// UnimplementedTimesheetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTimesheetServiceServer struct {
}

func (UnimplementedTimesheetServiceServer) CreateTimesheet(context.Context, *CreateTimesheetRequest) (*Timesheet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTimesheet not implemented")
}
func (UnimplementedTimesheetServiceServer) GetTimesheetsByContractor(context.Context, *GetTimesheetsByContractorRequest) (*GetTimesheetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimesheetsByContractor not implemented")
}
func (UnimplementedTimesheetServiceServer) GetTimesheetsByContractorAndClient(context.Context, *GetTimesheetsByContractorAndClientRequest) (*GetTimesheetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimesheetsByContractorAndClient not implemented")
}
func (UnimplementedTimesheetServiceServer) mustEmbedUnimplementedTimesheetServiceServer() {}

// UnsafeTimesheetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimesheetServiceServer will
// result in compilation errors.
type UnsafeTimesheetServiceServer interface {
	mustEmbedUnimplementedTimesheetServiceServer()
}

func RegisterTimesheetServiceServer(s grpc.ServiceRegistrar, srv TimesheetServiceServer) {
	s.RegisterService(&TimesheetService_ServiceDesc, srv)
}

func _TimesheetService_CreateTimesheet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTimesheetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimesheetServiceServer).CreateTimesheet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimesheetService_CreateTimesheet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimesheetServiceServer).CreateTimesheet(ctx, req.(*CreateTimesheetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimesheetService_GetTimesheetsByContractor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimesheetsByContractorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimesheetServiceServer).GetTimesheetsByContractor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimesheetService_GetTimesheetsByContractor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimesheetServiceServer).GetTimesheetsByContractor(ctx, req.(*GetTimesheetsByContractorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimesheetService_GetTimesheetsByContractorAndClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimesheetsByContractorAndClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimesheetServiceServer).GetTimesheetsByContractorAndClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimesheetService_GetTimesheetsByContractorAndClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimesheetServiceServer).GetTimesheetsByContractorAndClient(ctx, req.(*GetTimesheetsByContractorAndClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TimesheetService_ServiceDesc is the grpc.ServiceDesc for TimesheetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimesheetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "timesheet_app.TimesheetService",
	HandlerType: (*TimesheetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTimesheet",
			Handler:    _TimesheetService_CreateTimesheet_Handler,
		},
		{
			MethodName: "GetTimesheetsByContractor",
			Handler:    _TimesheetService_GetTimesheetsByContractor_Handler,
		},
		{
			MethodName: "GetTimesheetsByContractorAndClient",
			Handler:    _TimesheetService_GetTimesheetsByContractorAndClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/timesheet/timesheet.proto",
}
