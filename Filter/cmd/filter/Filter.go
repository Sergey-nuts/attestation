package main

import (
	"filter/pkg/api"
	"log"
	"net/http"
)

func main() {
	api := api.New()

	// // запуск веб-сервера с API
	log.Println("starting Filter server")
	err := http.ListenAndServe(":2020", api.Router())
	if err != nil {
		log.Fatal(err)
	}
}
