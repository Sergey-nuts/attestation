package postgr

import (
	"Comments/pkg/storage"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Postgres struct {
	db *pgxpool.Pool
}

// конструктор БД
func New(conf string) (*Postgres, error) {
	db, err := pgxpool.Connect(context.Background(), conf)
	if err != nil {
		return nil, err
	}
	return &Postgres{db}, nil
}

// AddComment добавляет комментарий в базу данных
func (p *Postgres) AddComment(c storage.Comment) error {
	ctx := context.Background()

	_, err := p.db.Exec(ctx, `
		INSERT INTO comments(postid, content, author, pubtime)
		VALUES ($1, $2, $3, $4)
	`, c.PostID, c.Text, c.Author, c.PubTime)

	if err != nil {
		return fmt.Errorf("postgrsql: %w", err)
	}
	return nil
}

// Comments возвращает список комментариев из базы данных
// к новости с postid
func (p *Postgres) Comments(postid int) ([]storage.Comment, error) {
	ctx := context.Background()

	rows, err := p.db.Query(ctx, `
		SELECT id, postid, content, author, pubtime
		FROM comments
		WHERE postid = $1
		ORDER BY pubtime DESC
	`, postid)
	if err != nil {
		return nil, err
	}

	var comments []storage.Comment
	var c storage.Comment
	for rows.Next() {
		err := rows.Scan(&c.ID, &c.PostID, &c.Text, &c.Author, &c.PubTime)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, rows.Err()
}

// Close закрывает все соединения с базой
func (p *Postgres) Close() {
	p.db.Close()
}
