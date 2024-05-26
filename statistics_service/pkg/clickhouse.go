package statistics_listener

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/ClickHouse/clickhouse-go"

	pb "statistics_service/pkg/pb"
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
		`CREATE TABLE IF NOT EXISTS views (user_id String, post_id String, author_id String) ENGINE = MergeTree() ORDER BY user_id`,
		`CREATE TABLE IF NOT EXISTS likes (user_id String, post_id String, author_id String) ENGINE = MergeTree() ORDER BY user_id`,
	}

	for _, query := range queries {
		if _, err := c.db.Exec(query); err != nil {
			return fmt.Errorf("failed to create table: %w", err)
		}
	}

	return nil
}

func (c *Clickhouse) SaveView(userId, postId, authorID string) error {
	query := `SELECT COUNT(*) FROM views WHERE user_id = ? AND post_id = ? AND author_id = ?`
	var count int
	if err := c.db.QueryRow(query, userId, postId, authorID).Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	query = `INSERT INTO views (user_id, post_id, author_id) VALUES (?, ?, ?)`
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, userId, postId, authorID)
	tx.Commit()
	return err
}

func (c *Clickhouse) SaveLike(userId, postId, authorID string) error {
	query := `SELECT COUNT(*) FROM likes WHERE user_id = ? AND post_id = ? AND author_id = ?`
	var count int
	if err := c.db.QueryRow(query, userId, postId, authorID).Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	query = `INSERT INTO likes (user_id, post_id, author_id) VALUES (?, ?, ?)`
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, userId, postId, authorID)
	tx.Commit()
	return err
}

func (c *Clickhouse) GetPostStats(postId string) (int, int, error) {
	var views, likes int

	query := `SELECT COUNT(DISTINCT user_id) as count FROM views WHERE post_id = ?`
	row := c.db.QueryRow(query, postId)
	err := row.Scan(&views)
	if err != nil {
		return 0, 0, err
	}

	query = `SELECT COUNT(DISTINCT user_id) as count FROM likes WHERE post_id = ?`
	row = c.db.QueryRow(query, postId)
	err = row.Scan(&likes)
	if err != nil {
		return 0, 0, err
	}

	return views, likes, nil
}

func (c *Clickhouse) GetTopPosts(sortBy string) ([]*pb.Post, error) {
	const topPostsCount = 5
	var query string
	if sortBy == "views" || sortBy == "likes" {
		query = fmt.Sprintf(`SELECT post_id, author_id, COUNT(*) as count FROM %s GROUP BY post_id, author_id ORDER BY count DESC LIMIT %d`, sortBy, topPostsCount)
	} else {
		return nil, fmt.Errorf("invalid sort type")
	}

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*pb.Post
	for rows.Next() {
		var post pb.Post
		err := rows.Scan(&post.Id, &post.AuthorId, &post.Count)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (c *Clickhouse) GetTopUsers() ([]*pb.User, error) {
	const topUsersCount = 3
	query := `SELECT author_id, COUNT(*) as total_likes FROM likes GROUP BY author_id ORDER BY total_likes DESC LIMIT ` + fmt.Sprint(topUsersCount)
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*pb.User
	for rows.Next() {
		var user pb.User
		err := rows.Scan(&user.Id, &user.LikesCount)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
