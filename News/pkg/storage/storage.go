package storage

// Публикация, получаемая из RSS.
type Post struct {
	ID      int    `json:"id"`                // номер записи
	Title   string `json:"title"`             // заголовок публикации
	Content string `json:"content,omitempty"` // содержание публикации
	PubTime int64  `json:"pubtime,omitempty"` // время публикации
	Link    string `json:"link,omitempty"`    // ссылка на источник
}

type Interfase interface {
	// добавляет новость
	AddNews([]Post) error
	// возвращает список новостей
	NewsList(int, int) ([]Post, error)
	// поиск новостей по названию
	Search(string) ([]Post, error)
	// возвращает количество запичсей в базе
	Count() int
	// возвращает новость по ее id
	PostId(int) (Post, error)
}
