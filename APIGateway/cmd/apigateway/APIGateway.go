package main

import (
	"apigateway/pkg/api"
	"log"
	"net/http"
)

func main() {

	api := api.New()

	// запуск веб-сервера с APIGateway
	log.Println("starting APIGateway server")
	err := http.ListenAndServe(":8080", api.Router())
	if err != nil {
		log.Fatal(err)
	}
}
