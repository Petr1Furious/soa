package grpcServer

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "posts_service/pkg/pb"

	"github.com/jackc/pgx/v5"
)

type Server struct {
	pb.UnimplementedPostServiceServer
	db Database
}

func NewServer() (*Server, error) {
	var err error
	host, ok := os.LookupEnv("POSTS_DB_HOST")
	if !ok {
		return nil, fmt.Errorf("POSTS_DB_HOST not set")
	}
	username, ok := os.LookupEnv("POSTS_DB_USERNAME")
	if !ok {
		return nil, fmt.Errorf("POSTS_DB_USERNAME not set")
	}
	password, ok := os.LookupEnv("POSTS_DB_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("POSTS_DB_PASSWORD not set")
	}
	dbname, ok := os.LookupEnv("POSTS_DB_NAME")
	if !ok {
		return nil, fmt.Errorf("POSTS_DB_NAME not set")
	}

	var conn *pgx.Conn
	for i := 0; i < 5; i++ {
		connString := fmt.Sprintf("postgresql://%s:%s@%s/%s", username, password, host, dbname)
		conn, err = pgx.Connect(context.Background(), connString)
		if err == nil {
			break
		}

		if i < 4 {
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}

		return nil, fmt.Errorf("unable to connect to database after 5 attempts: %v", err)
	}

	_, err = conn.Exec(context.Background(), `
        CREATE TABLE IF NOT EXISTS posts (
            id SERIAL PRIMARY KEY,
            content TEXT NOT NULL,
            user_id BIGINT NOT NULL
        )
    `)
	if err != nil {
		return nil, fmt.Errorf("unable to create table: %v", err)
	}

	return &Server{db: &RealDatabase{conn}}, nil
}

func (s *Server) checkUserPermission(ctx context.Context, userId int64, postId int64) (bool, error) {
	var dbUserId int64
	err := s.db.QueryRow(ctx, `SELECT user_id FROM posts WHERE id = $1`, postId).Scan(&dbUserId)
	if err != nil {
		return false, err
	}

	return dbUserId == userId, nil
}

func (s *Server) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.PostResponse, error) {
	var id int64
	err := s.db.QueryRow(ctx, `INSERT INTO posts (content, user_id) VALUES ($1, $2) RETURNING id`, req.Content, req.UserId).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &pb.PostResponse{PostId: id, Content: req.Content, UserId: req.UserId}, nil
}

func (s *Server) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.PostResponse, error) {
	if ok, err := s.checkUserPermission(ctx, req.UserId, req.PostId); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("user does not have permission to update post")
	}

	_, err := s.db.Exec(ctx, `UPDATE posts SET content = $1 WHERE id = $2 AND user_id = $3`, req.Content, req.PostId, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.PostResponse{PostId: req.PostId, Content: req.Content, UserId: req.UserId}, nil
}

func (s *Server) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	if ok, err := s.checkUserPermission(ctx, req.UserId, req.PostId); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("user does not have permission to update post")
	}

	_, err := s.db.Exec(ctx, `DELETE FROM posts WHERE id = $1 AND user_id = $2`, req.PostId, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.DeletePostResponse{PostId: req.PostId, Success: true}, nil
}

func (s *Server) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.PostResponse, error) {
	var content string
	var userId int64
	err := s.db.QueryRow(ctx, `SELECT content, user_id FROM posts WHERE id = $1`, req.PostId).Scan(&content, &userId)
	if err != nil {
		return nil, err
	}

	return &pb.PostResponse{PostId: req.PostId, Content: content, UserId: userId}, nil
}

func (s *Server) ListPosts(ctx context.Context, req *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	rows, err := s.db.Query(ctx, `SELECT id, content, user_id FROM posts WHERE user_id = $1 LIMIT $2 OFFSET $3`, req.UserId, req.PageSize, req.Page*req.PageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*pb.PostResponse
	for rows.Next() {
		var id, userId int64
		var content string
		err := rows.Scan(&id, &content, &userId)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &pb.PostResponse{PostId: id, Content: content, UserId: userId})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &pb.ListPostsResponse{Posts: posts}, nil
}
