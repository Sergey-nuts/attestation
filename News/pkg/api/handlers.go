package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	// limit news on one page
	limit = 10
)

type PaginationData struct {
	CurrentPage int
	TotalPages  int
}

// получить новость по ее id
func (api *API) postIDHandler(w http.ResponseWriter, r *http.Request) {
	s := mux.Vars(r)["postid"]
	postid, _ := strconv.Atoi(s)

	post, err := api.db.PostId(postid)
	if err != nil {
		// log.Printf("postIDHandler err: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// Получение списка новостей
func (api *API) postsHandler(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if page < 1 {
		page = 1
	}

	offset := limit * (page - 1)
	posts, err := api.db.NewsList(limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	records := api.db.Count()
	total := records / limit
	remain := records % limit
	if remain != 0 {
		total += 1
	}
	pg := PaginationData{TotalPages: total, CurrentPage: page}
	data := map[string]interface{}{}
	data["pagination"] = pg
	data["news"] = posts

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Поиск новостей с подстрокой в названии
func (api *API) searchHandler(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("value")
	if search == "" {
		http.Error(w, "not found searchValue", http.StatusBadRequest)
		// log.Printf("searchHandler: not found Value get:\"%s\"\n", search)
		return
	}

	posts, err := api.db.Search(search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
