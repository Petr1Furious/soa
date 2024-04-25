package statistics_listener

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/ClickHouse/clickhouse-go"
)

type Clickhouse struct {
	db *sql.DB
}

func NewClickhouse() (*Clickhouse, error) {
	host := os.Getenv("STATS_DB_HOST")
	username := os.Getenv("STATS_DB_USERNAME")
	password := os.Getenv("STATS_DB_PASSWORD")

	dsn := fmt.Sprintf("tcp://%s:9000?username=%s&password=%s",
		host, username, password)

	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to clickhouse: %w", err)
	}

	clickhouse := &Clickhouse{db: db}
	err = clickhouse.CreateTables()
	if err != nil {
		return nil, fmt.Errorf("failed to create clickhouse tables: %w", err)
	}

	return clickhouse, nil
}

func (c *Clickhouse) CreateTables() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS views (user_id String, post_id String) ENGINE = Memory`,
		`CREATE TABLE IF NOT EXISTS likes (user_id String, post_id String) ENGINE = Memory`,
	}

	for _, query := range queries {
		if _, err := c.db.Exec(query); err != nil {
			return fmt.Errorf("failed to create table: %w", err)
		}
	}

	return nil
}

func (c *Clickhouse) SaveView(userId, postId string) error {
	query := `INSERT INTO views (user_id, post_id) VALUES (?, ?)`
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, userId, postId)
	tx.Commit()
	return err
}

func (c *Clickhouse) SaveLike(userId, postId string) error {
	query := `INSERT INTO likes (user_id, post_id) VALUES (?, ?)`
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, userId, postId)
	tx.Commit()
	return err
}
