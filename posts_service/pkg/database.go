package grpcServer

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Database interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}

type RealDatabase struct {
	conn *pgx.Conn
}

func (db *RealDatabase) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return db.conn.QueryRow(ctx, sql, args...)
}

func (db *RealDatabase) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.conn.Exec(ctx, sql, args...)
}

func (db *RealDatabase) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return db.conn.Query(ctx, sql, args...)
}
