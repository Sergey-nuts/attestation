package api

import (
	"Comments/pkg/storage"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var FilterHost = "localhost"

type API struct {
	r  *mux.Router
	db storage.Interface
}

// Конструктор API
func New(db storage.Interface) *API {
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

	// получить комментарии к новости postid
	api.r.HandleFunc("/comments/{postid:[0-9]+}", api.commentsHandler).Methods(http.MethodGet, http.MethodOptions)

	// добавить комментарий к новости postid
	api.r.HandleFunc("/comments", api.addCommentsHandler).Methods(http.MethodPost, http.MethodOptions)
}

// Router возвращает маршрутизатор запросов
func (api *API) Router() *mux.Router {
	return api.r
}

// возвращает настроеного клиента для запроса
func newClient() *http.Client {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}
	return client
}
