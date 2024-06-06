package grpcServer

import (
	"context"
	"log"
	pb "posts_service/pkg/pb"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestPostsService(t *testing.T) {
	ctx := context.Background()

	containerRequest := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context: "../posts_service",
		},
		WaitingFor: wait.ForLog("Posts service started on"),
	}
	postsServiceContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: containerRequest,
		Started:          true,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer postsServiceContainer.Terminate(ctx)
	log.Println("Docker container started")

	ip, err := postsServiceContainer.Host(ctx)
	if err != nil {
		t.Fatal(err)
	}
	port, err := postsServiceContainer.MappedPort(ctx, "8080")
	if err != nil {
		t.Fatal(err)
	}

	conn, err := grpc.NewClient(ip+":"+port.Port(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	log.Println("gRPC connection established")

	client := pb.NewPostServiceClient(conn)

	req := &pb.CreatePostRequest{
		Content: "Hello, world!",
		UserId:  2,
	}

	resp, err := client.CreatePost(ctx, req)
	if err != nil {
		t.Fatalf("Failed to call GetPost: %v", err)
	}
	log.Println("CreatePost RPC call returned")

	if resp.UserId != req.UserId {
		t.Errorf("Expected post ID %q, got %q", req.UserId, resp.UserId)
	}
	if resp.Content == "" {
		t.Error("Expected non-empty post title")
	}
}
