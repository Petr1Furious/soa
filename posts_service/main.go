package main

import (
	"log"
	"net"

	posts "posts_service/pkg"
	pb "posts_service/pkg/pb"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server, err := posts.NewServer()
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, server)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
