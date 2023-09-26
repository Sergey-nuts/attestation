package main

import (
	"Comments/pkg/api"
	"Comments/pkg/storage/postgr"
	"log"
	"net/http"
	"os"
)

func main() {
	postgrUser := os.Getenv("dbuser")
	postgrPwd := os.Getenv("dbpass")
	dbhost := os.Getenv("dbhost")
	db, err := postgr.New("postgres://" + postgrUser + ":" + postgrPwd + "@" + dbhost + "/Comments")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("postgresql connected.\n")
	defer db.Close()

	api := api.New(db)
	// запуск веб-сервера с API и приложением
	log.Println("starting Comments server")
	err = http.ListenAndServe(":2010", api.Router())
	if err != nil {
		log.Fatal(err)
	}
}
