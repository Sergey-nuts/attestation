package storage

// Публикация, получаемая из RSS.
type Post struct {
	ID      int    `json:"id"`                // номер записи
	Title   string `json:"title"`             // заголовок публикации
	Content string `json:"content,omitempty"` // содержание публикации
	PubTime int64  `json:"pubtime,omitempty"` // время публикации
	Link    string `json:"link,omitempty"`    // ссылка на источник
}

// Комментарий к новости
type Comment struct {
	ID      int    // номер записи
	PostID  int    // номер новости
	Text    string // текст комментария
	Author  string // автор комментария
	PubTime int64  // время публикации
}
