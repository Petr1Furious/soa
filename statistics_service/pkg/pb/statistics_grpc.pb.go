// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: statistics_service/pkg/pb/statistics.proto

package pb

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
	StatisticsService_GetPostStats_FullMethodName = "/statistics.StatisticsService/GetPostStats"
	StatisticsService_GetTopPosts_FullMethodName  = "/statistics.StatisticsService/GetTopPosts"
	StatisticsService_GetTopUsers_FullMethodName  = "/statistics.StatisticsService/GetTopUsers"
)

// StatisticsServiceClient is the client API for StatisticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatisticsServiceClient interface {
	GetPostStats(ctx context.Context, in *PostStatsRequest, opts ...grpc.CallOption) (*PostStatsResponse, error)
	GetTopPosts(ctx context.Context, in *TopPostsRequest, opts ...grpc.CallOption) (*TopPostsResponse, error)
	GetTopUsers(ctx context.Context, in *TopUsersRequest, opts ...grpc.CallOption) (*TopUsersResponse, error)
}

type statisticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatisticsServiceClient(cc grpc.ClientConnInterface) StatisticsServiceClient {
	return &statisticsServiceClient{cc}
}

func (c *statisticsServiceClient) GetPostStats(ctx context.Context, in *PostStatsRequest, opts ...grpc.CallOption) (*PostStatsResponse, error) {
	out := new(PostStatsResponse)
	err := c.cc.Invoke(ctx, StatisticsService_GetPostStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statisticsServiceClient) GetTopPosts(ctx context.Context, in *TopPostsRequest, opts ...grpc.CallOption) (*TopPostsResponse, error) {
	out := new(TopPostsResponse)
	err := c.cc.Invoke(ctx, StatisticsService_GetTopPosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statisticsServiceClient) GetTopUsers(ctx context.Context, in *TopUsersRequest, opts ...grpc.CallOption) (*TopUsersResponse, error) {
	out := new(TopUsersResponse)
	err := c.cc.Invoke(ctx, StatisticsService_GetTopUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatisticsServiceServer is the server API for StatisticsService service.
// All implementations must embed UnimplementedStatisticsServiceServer
// for forward compatibility
type StatisticsServiceServer interface {
	GetPostStats(context.Context, *PostStatsRequest) (*PostStatsResponse, error)
	GetTopPosts(context.Context, *TopPostsRequest) (*TopPostsResponse, error)
	GetTopUsers(context.Context, *TopUsersRequest) (*TopUsersResponse, error)
	mustEmbedUnimplementedStatisticsServiceServer()
}

// UnimplementedStatisticsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStatisticsServiceServer struct {
}

func (UnimplementedStatisticsServiceServer) GetPostStats(context.Context, *PostStatsRequest) (*PostStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostStats not implemented")
}
func (UnimplementedStatisticsServiceServer) GetTopPosts(context.Context, *TopPostsRequest) (*TopPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopPosts not implemented")
}
func (UnimplementedStatisticsServiceServer) GetTopUsers(context.Context, *TopUsersRequest) (*TopUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopUsers not implemented")
}
func (UnimplementedStatisticsServiceServer) mustEmbedUnimplementedStatisticsServiceServer() {}

// UnsafeStatisticsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatisticsServiceServer will
// result in compilation errors.
type UnsafeStatisticsServiceServer interface {
	mustEmbedUnimplementedStatisticsServiceServer()
}

func RegisterStatisticsServiceServer(s grpc.ServiceRegistrar, srv StatisticsServiceServer) {
	s.RegisterService(&StatisticsService_ServiceDesc, srv)
}

func _StatisticsService_GetPostStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatisticsServiceServer).GetPostStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatisticsService_GetPostStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatisticsServiceServer).GetPostStats(ctx, req.(*PostStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatisticsService_GetTopPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatisticsServiceServer).GetTopPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatisticsService_GetTopPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatisticsServiceServer).GetTopPosts(ctx, req.(*TopPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatisticsService_GetTopUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatisticsServiceServer).GetTopUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatisticsService_GetTopUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatisticsServiceServer).GetTopUsers(ctx, req.(*TopUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StatisticsService_ServiceDesc is the grpc.ServiceDesc for StatisticsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatisticsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "statistics.StatisticsService",
	HandlerType: (*StatisticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPostStats",
			Handler:    _StatisticsService_GetPostStats_Handler,
		},
		{
			MethodName: "GetTopPosts",
			Handler:    _StatisticsService_GetTopPosts_Handler,
		},
		{
			MethodName: "GetTopUsers",
			Handler:    _StatisticsService_GetTopUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "statistics_service/pkg/pb/statistics.proto",
}
