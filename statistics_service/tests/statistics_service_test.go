package statistics_service_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	pb "statistics_service/pkg/pb"
	"testing"
	"time"

	kpb "statistics_service/pkg/kafka_pb"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

func testGRPC(ctx context.Context) error {
	conn, err := grpc.NewClient("localhost:58158", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewStatisticsServiceClient(conn)

	req := &pb.TopPostsRequest{
		Type: "views",
	}

	resp, err := client.GetTopPosts(ctx, req)
	if err != nil {
		return err
	}

	if len(resp.Posts) != 0 {
		return fmt.Errorf("Expected empty list of posts")
	}

	req2 := &pb.TopUsersRequest{}

	resp2, err := client.GetTopUsers(ctx, req2)
	if err != nil {
		return err
	}

	if len(resp2.Users) != 0 {
		return fmt.Errorf("Expected empty list of users")
	}

	return nil
}

func testKafka(ctx context.Context) error {
	kafkaAddr := "192.168.1.2:9092"
	kafkaProducer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{kafkaAddr},
		Topic:   "my-topic",
		Logger:  log.New(os.Stdout, "kafka writer: ", 0),
	})

	db, err := sql.Open("clickhouse", "tcp://localhost:9228?username=test&password=test")
	if err != nil {
		return err
	}

	UserID := int64(0)
	PostID := int64(1)
	AuthorID := int64(2)
	viewEvent := &kpb.Event_ViewEvent{
		ViewEvent: &kpb.ViewEvent{
			UserId:   UserID,
			PostId:   PostID,
			AuthorId: AuthorID,
		},
	}
	event := &kpb.Event{
		EventType: viewEvent,
	}

	msg, err := proto.Marshal(event)
	if err != nil {
		return err
	}

	err = kafkaProducer.WriteMessages(ctx, kafka.Message{
		Value: msg,
	})
	if err != nil {
		return err
	}

	time.Sleep(10 * time.Second)

	query := "SELECT user_id, post_id, author_id FROM views"
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	var userId, postId, authorId int64
	if !rows.Next() {
		return fmt.Errorf("No rows in views table")
	}
	err = rows.Scan(&userId, &postId, &authorId)
	if err != nil {
		return err
	}
	if rows.Next() {
		return err
	}

	if userId != 0 || postId != 1 || authorId != 2 {
		return fmt.Errorf("Expected (0, 1, 2), got (%d, %d, %d)", userId, postId, authorId)
	}

	return nil
}

func setUpDockerCompose(ctx context.Context) error {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	dockerComposeFilePath := filepath.Join(dir, "../docker-compose-test.yml")

	cmd := exec.CommandContext(ctx, "docker", "compose", "-f", dockerComposeFilePath, "up", "-d", "--build")
	return cmd.Run()
}

func tearDownDockerCompose(ctx context.Context) {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	dockerComposeFilePath := filepath.Join(dir, "../docker-compose-test.yml")

	cmd := exec.CommandContext(ctx, "docker", "compose", "-f", dockerComposeFilePath, "down")
	cmd.Run()
}

func TestPostsServiceGRPC(t *testing.T) {
	ctx := context.Background()

	err := setUpDockerCompose(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tearDownDockerCompose(ctx)

	time.Sleep(10 * time.Second)

	err = testGRPC(ctx)
	if err != nil {
		t.Fatal(err)
	}

	err = testKafka(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
