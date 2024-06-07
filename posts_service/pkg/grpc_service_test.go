package grpcServer

import (
	"context"
	"testing"

	pb "posts_service/pkg/pb"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDatabase struct {
	mock.Mock
}

type MockRow struct {
}

func (r MockRow) Scan(dest ...interface{}) error {
	id := dest[0].(*int64)
	*id = 1
	return nil
}

func (db *MockDatabase) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	db.Called(ctx, sql, args)
	return MockRow{}
}

func (db *MockDatabase) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	db.Called(ctx, sql, args)
	return pgconn.CommandTag{}, nil
}

func (db *MockDatabase) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	db.Called(ctx, sql, args)
	return nil, nil
}

func TestCreatePost(t *testing.T) {
	db := new(MockDatabase)
	server := &Server{db: db}

	db.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(MockRow{})

	postResponse, err := server.CreatePost(context.Background(), &pb.CreatePostRequest{Content: "test content", UserId: 1})

	assert.NoError(t, err)
	db.AssertExpectations(t)

	assert.Equal(t, postResponse, &pb.PostResponse{PostId: 1, Content: "test content", UserId: 1})
}

func TestUpdatePost(t *testing.T) {
	db := new(MockDatabase)
	server := &Server{db: db}

	db.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(MockRow{})
	db.On("Exec", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)

	req := &pb.UpdatePostRequest{PostId: 1, Content: "updated content", UserId: 1}
	resp, err := server.UpdatePost(context.Background(), req)

	assert.NoError(t, err)
	db.AssertExpectations(t)

	expectedResponse := &pb.PostResponse{PostId: req.PostId, Content: req.Content, UserId: req.UserId}
	assert.Equal(t, expectedResponse, resp)
}
