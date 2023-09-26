package postgr

import (
	"News/pkg/storage"
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

// AddNews добавляет новости из news в базу даннх
func (p *Postgres) AddNews(news []storage.Post) error {
	ctx := context.Background()
	for _, post := range news {
		_, err := p.db.Exec(ctx, `
			INSERT INTO news(title, content, pubtime, link)
			VALUES ($1, $2, $3, $4)
			ON CONFLICT (link) DO NOTHING;
		`, post.Title, post.Content, post.PubTime, post.Link)
		if err != nil {
			return fmt.Errorf("postgrsql: %w", err)
		}
	}

	return nil
}

// News возвращает последние n новостенй из базы данных
func (p *Postgres) NewsList(limit, offset int) ([]storage.Post, error) {
	rows, err := p.db.Query(context.Background(), `
		SELECT id, title 
		FROM news
		ORDER BY pubtime DESC
		LIMIT $1
		OFFSET $2;
	`,
		limit, offset,
	)
	if err != nil {
		return nil, err
	}

	var news []storage.Post
	var post storage.Post
	for rows.Next() {
		err := rows.Scan(&post.ID, &post.Title)
		if err != nil {
			return nil, err
		}
		news = append(news, post)
	}

	return news, rows.Err()
}

// Search возвращяет список новостей в которых есть подстрока s
func (p *Postgres) Search(s string) ([]storage.Post, error) {
	sql := fmt.Sprintf(`SELECT id, title FROM news WHERE title ILIKE '%%%v%%' ORDER BY pubtime DESC;`, s)
	rows, err := p.db.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	var post storage.Post
	var news []storage.Post
	for rows.Next() {
		err := rows.Scan(&post.ID, &post.Title)
		if err != nil {
			return nil, err
		}
		news = append(news, post)
	}
	return news, rows.Err()
}

// Count возвращает количество записей в базе новостей
func (p *Postgres) Count() int {
	row := p.db.QueryRow(context.Background(), `SELECT count(id) FROM news;`)
	var res int
	row.Scan(&res)
	return res
}

// PostId возвращает новость по ее id
func (p *Postgres) PostId(id int) (storage.Post, error) {
	row := p.db.QueryRow(context.Background(), `
		SELECT id, title,  content, pubtime, link
		FROM news
		WHERE id = $1;
	`, id)

	var res storage.Post
	err := row.Scan(&res.ID, &res.Title, &res.Content, &res.PubTime, &res.Link)
	if err != nil {
		return storage.Post{}, err
	}

	return res, nil
}

// Close закрывает все соединения с базой
func (p *Postgres) Close() {
	p.db.Close()
}
