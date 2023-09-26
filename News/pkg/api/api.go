package api

import (
	"net/http"

	"News/pkg/storage"

	"github.com/gorilla/mux"
)

type API struct {
	r  *mux.Router
	db storage.Interfase
}

// Конструктор API
func New(db storage.Interfase) *API {
	api := API{}
	api.db = db
	api.r = mux.NewRouter()
	api.endpoints()
	return &api
}

// Регистрация методов API в маршрутизаторе запросов.
func (api *API) endpoints() {
	api.r.Use(requestIDMiddleware)
	api.r.Use(logMiddleware)

	// получить новость по ее id
	api.r.HandleFunc("/news/{postid:[0-9]+}", api.postIDHandler).Methods(http.MethodGet, http.MethodOptions)

	// получить списка новостей
	api.r.HandleFunc("/news", api.postsHandler).Methods(http.MethodGet, http.MethodOptions)

	// поиск по названию новостей
	api.r.HandleFunc("/news/search", api.searchHandler).Methods(http.MethodGet, http.MethodOptions)

}

// Router возвращает маршрутизатор запросов
func (api *API) Router() *mux.Router {
	return api.r
}
