package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	NewsHost     = "localhost"
	CommentsHost = "localhost"
)

type API struct {
	r *mux.Router
}

// Конструктор API
func New() *API {
	api := API{}
	api.r = mux.NewRouter()
	api.endpoints()
	return &api
}

// Router возвращает маршрутизатор запросов
func (api *API) Router() *mux.Router {
	return api.r
}

// Регистрация методов API в маршрутизаторе запросов.
func (api *API) endpoints() {
	api.r.Use(requestIDMiddleware)
	api.r.Use(logMiddleware)

	// получить список новостей
	api.r.HandleFunc("/news", api.newsShortDetailedHandler).Methods(http.MethodGet, http.MethodOptions)

	// получить детальную информацию и комментарии к новости postid
	api.r.HandleFunc("/news/{postid:[0-9]+}/full", api.newsFullDetailedHandler).Methods(http.MethodGet, http.MethodOptions)

	// добавить комментарии к новости postid
	api.r.HandleFunc("/news/comments", api.commentsHandler).Methods(http.MethodPost, http.MethodOptions)

	// поиск новостей по названию
	api.r.HandleFunc("/news/search", api.searchHandler).Methods(http.MethodGet, http.MethodOptions)

}
