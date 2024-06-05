package statistics_listener

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "statistics_service/pkg/pb"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedStatisticsServiceServer
	clickhouse *Clickhouse
}

func NewServer() (*Server, error) {
	clickhouse, err := NewClickhouse()
	if err != nil {
		return nil, fmt.Errorf("failed to create clickhouse: %w", err)
	}
	return &Server{clickhouse: clickhouse}, nil
}

func (s *Server) Listen() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStatisticsServiceServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) GetPostStats(ctx context.Context, req *pb.PostStatsRequest) (*pb.PostStatsResponse, error) {
	views, likes, err := s.clickhouse.GetPostStats(req.PostId)
	if err != nil {
		return nil, err
	}
	return &pb.PostStatsResponse{ViewCount: views, LikeCount: likes}, nil
}

func (s *Server) GetTopPosts(ctx context.Context, req *pb.TopPostsRequest) (*pb.TopPostsResponse, error) {
	posts, err := s.clickhouse.GetTopPosts(req.Type)
	if err != nil {
		return nil, err
	}
	return &pb.TopPostsResponse{Posts: posts}, nil
}

func (s *Server) GetTopUsers(ctx context.Context, req *pb.TopUsersRequest) (*pb.TopUsersResponse, error) {
	users, err := s.clickhouse.GetTopUsers()
	if err != nil {
		return nil, err
	}
	return &pb.TopUsersResponse{Users: users}, nil
}
