package api

import (
	"Comments/pkg/storage"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Item struct {
	PostID  int       `json:"postid"`
	Text    string    `json:"text"`
	Author  string    `json:"author"`
	PubTime time.Time `json:"pubtime"`
}

// получение комментариев к новости с postid
func (api *API) commentsHandler(w http.ResponseWriter, r *http.Request) {
	postid, _ := strconv.Atoi(mux.Vars(r)["postid"])
	comments, err := api.db.Comments(postid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		// log.Printf("commentsHandler err: %v\n", err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

// добавление комментария в базу
func (api *API) addCommentsHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var item Item
	err := dec.Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		// log.Printf("addCommentHandler decode err: %v\n", err.Error())
		return
	}

	// подготовка запроса в сервис цензурирования
	type massage struct {
		Text string `json:"text"`
	}
	m := massage{Text: item.Text}
	// url := "http://localhost:2020/filter"
	url := fmt.Sprintf("http://%s:2020/filter", FilterHost)
	requestId := r.Context().Value(requestIdHeader)
	body, err := json.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("addCommentsHandler err: %v\n", err)
		return
	}
	req, err := http.NewRequestWithContext(r.Context(), http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("addCommentsHandler err: %v\n", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(requestIdHeader, fmt.Sprint(requestId))

	client := newClient()
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		// log.Printf("addCommentsHandler err: %v\n", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		str := buf.String()
		http.Error(w, str, resp.StatusCode)
		// log.Printf("addCommentsHandler response status from filter: %v\n", resp.Status)
		return

	}

	// запись комментария в базу
	var c storage.Comment
	c.PostID, c.Text, c.Author, c.PubTime = item.PostID, item.Text, item.Author, item.PubTime.Unix()
	err = api.db.AddComment(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// ответ на первоначальный запрос
	w.WriteHeader(http.StatusOK)
}
