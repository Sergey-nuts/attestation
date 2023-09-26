package api

import (
	"encoding/json"
	"filter/pkg/filter"
	"net/http"

	"github.com/gorilla/mux"
)

type comment struct {
	Text string `json:"text"`
}

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

// Регистрация методов API в маршрутизаторе запросов.
func (api *API) endpoints() {
	api.r.Use(requestIDMiddleware)
	api.r.Use(logMiddleware)

	// цензура комментария
	api.r.HandleFunc("/filter", api.filterHandler).Methods(http.MethodPost, http.MethodOptions)
}

// Router возвращает маршрутизатор запросов
func (api *API) Router() *mux.Router {
	return api.r
}

func (api *API) filterHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var item comment
	err := dec.Decode(&item)
	if err != nil {
		// log.Printf("filterHandler err: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if filter.Censorship(item.Text) {
		http.Error(w, "You are bad boy!", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
