package posts_service_test

import (
	"context"
	pb "posts_service/pkg/pb"
	"testing"
	"time"

	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestPostsService(t *testing.T) {
	ctx := context.Background()

	compose, err := tc.NewDockerCompose("../docker-compose-test.yml")
	if err != nil {
		t.Fatal(err)
	}

	compose.WithEnv(map[string]string{
		"POSTS_DB_HOST":     "posts_db",
		"POSTS_DB_USERNAME": "test",
		"POSTS_DB_PASSWORD": "test",
		"POSTS_DB_NAME":     "test",
	}).Up(ctx)

	if err != nil {
		t.Fatal(err)
	}
	defer compose.Down(ctx)

	time.Sleep(10 * time.Second)

	conn, err := grpc.NewClient("localhost:58157", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewPostServiceClient(conn)

	req := &pb.CreatePostRequest{
		Content: "Hello, world!",
		UserId:  1,
	}

	resp, err := client.CreatePost(ctx, req)
	if err != nil {
		t.Fatalf("Failed to call CreatePost: %v", err)
	}

	if resp.UserId != req.UserId {
		t.Errorf("Expected user ID %q, got %q", req.UserId, resp.UserId)
	}
	if resp.Content != req.Content {
		t.Errorf("Expected content %q, got %q", req.Content, resp.Content)
	}

	req2 := &pb.GetPostRequest{
		PostId: resp.PostId,
	}

	resp2, err := client.GetPost(ctx, req2)
	if err != nil {
		t.Fatalf("Failed to call GetPost: %v", err)
	}

	if resp2.UserId != req.UserId {
		t.Errorf("Expected user ID %q, got %q", req.UserId, resp2.UserId)
	}
	if resp2.Content != req.Content {
		t.Errorf("Expected content %q, got %q", req.Content, resp2.Content)
	}
}
