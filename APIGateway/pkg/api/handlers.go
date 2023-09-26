package api

import (
	"apigateway/pkg/storage"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

// for json
type Comment struct {
	ID      int    `json:"id,omitempty"` // номер записи
	PostID  int    `json:"postid"`       // номер новости
	Text    string `json:"text"`         // текст комментария
	Author  string `json:"author"`       // автор комментария
	PubTime string `json:"pubtime"`      // время публикации
}

// for json
type Post struct {
	ID      int    `json:"id"`      // номер записи
	Title   string `json:"title"`   // заголовок публикации
	Content string `json:"content"` // содержание публикации
	PubTime string `json:"pubtime"` // время публикации
	Link    string `json:"link"`    // ссылка на источник
}

type PaginationData struct {
	CurrentPage int `json:"currentPage"`
	TotalPages  int `lson:"totalpages"`
}
type list struct {
	Posts      []storage.Post `json:"news"`
	Pagination PaginationData `json:"pagination,omitempty"`
}

// for json
type Result struct {
	News     Post      `json:"news,omitempty"`
	Comments []Comment `json:"comments,omitempty"`
}

// Получения списка новостей
func (api *API) newsShortDetailedHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}

	// подготовка запроса к сервису новостей
	client := newClient()
	url := "http://localhost:2000/news"
	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, url, nil) // ctx
	if err != nil {
		// log.Printf("newsShortHandler err: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	requestId := req.Context().Value(requestIdHeader)
	req.Header.Add(requestIdHeader, fmt.Sprint(requestId))
	q := req.URL.Query()
	q.Add("page", fmt.Sprintf("%v", page))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		// log.Printf("newsShortHandler err: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// расщифровка ответа
	dec := json.NewDecoder(resp.Body)
	var data list
	err = dec.Decode(&data)
	if err != nil {
		// log.Printf("newsShortHandler err: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}

// Получения детальной информации о новости с коммнетариями к ней
func (api *API) newsFullDetailedHandler(w http.ResponseWriter, r *http.Request) {
	s := mux.Vars(r)["postid"]
	postid, _ := strconv.Atoi(s)

	ch := make(chan res, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go postReq(&wg, r.Context(), postid, ch)
	go commentsReq(&wg, r.Context(), postid, ch)

	wg.Wait()
	close(ch)
	var err error
	var post storage.Post
	var comments []storage.Comment
	for c := range ch {
		err = c.err
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if c.comment != nil {
			comments = c.comment
			continue
		}
		post = c.post

	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var result Result

	result.News.ID, result.News.Title, result.News.Content, result.News.Link = post.ID, post.Title, post.Content, post.Link
	result.News.PubTime = time.Unix(post.PubTime, 0).String()
	for _, c := range comments {
		var comment Comment
		comment.ID = c.ID
		comment.PostID = c.ID
		comment.Text = c.Text
		comment.Author = c.Author
		comment.PubTime = time.Unix(c.PubTime, 0).String()
		result.Comments = append(result.Comments, comment)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// добавленик комментария к новости
func (api *API) commentsHandler(w http.ResponseWriter, r *http.Request) {
	// подготовка запроса в сервис комментариев
	url := "http://localhost:2010/comments"
	requestId := r.Context().Value(requestIdHeader)
	req, err := http.NewRequestWithContext(r.Context(), http.MethodPost, url, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("commentsHandler err: %v\n", err)
		return
	}
	req.Header.Add(requestIdHeader, fmt.Sprint(requestId))
	req.Header.Add("Content-Type", "application/json")
	client := newClient()

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		// log.Printf("commentsHandler err: %v\n", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		str := buf.String()
		http.Error(w, str, resp.StatusCode)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// поис новостей по названию
func (api *API) searchHandler(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("value")
	url := "http://localhost:2000/news/search"
	requestId := r.Header.Get(requestIdHeader)

	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, url, nil)
	if err != nil {
		log.Printf("searchHeader err: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Add(requestIdHeader, fmt.Sprint(requestId))
	q := req.URL.Query()
	q.Add("value", fmt.Sprintf("%v", search))
	req.URL.RawQuery = q.Encode()
	client := newClient()

	resp, err := client.Do(req)
	if err != nil {
		// log.Printf("searchHeader err: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// расшифровка ответа
	dec := json.NewDecoder(resp.Body)
	var data []storage.Post
	err = dec.Decode(&data)
	if err != nil {
		// log.Printf("newsShortHandler err: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}
