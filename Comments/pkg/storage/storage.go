package storage

type Comment struct {
	ID      int    `json:"id"`
	PostID  int    `json:"postid"`
	Text    string `json:"text"`
	Author  string `json:"author"`
	PubTime int64  `json:"pubtime"`
}

type Interface interface {
	// Добавление комментария в базу
	AddComment(Comment) error
	// Получение всех комментариев к новости по ее ID
	Comments(int) ([]Comment, error)
}
