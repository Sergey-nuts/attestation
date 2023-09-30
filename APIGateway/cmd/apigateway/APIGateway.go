package main

import (
	"apigateway/pkg/api"
	"log"
	"net/http"
	"os"
)

func main() {
	news := os.Getenv("news")
	if news != "" {
		api.NewsHost = news
	}
	comments := os.Getenv("comments")
	if comments != "" {
		api.CommentsHost = comments
	}
	api := api.New()

	// запуск веб-сервера с APIGateway
	log.Println("starting APIGateway server")
	err := http.ListenAndServe(":8080", api.Router())
	if err != nil {
		log.Fatal(err)
	}
}
