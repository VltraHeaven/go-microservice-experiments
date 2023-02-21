package main

import (
	"github.com/gorilla/mux"
	"github.com/vltraheaven/go-microservice-experiments/gorilla-mux-matching/handlers"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles", handlers.QueryHandler)
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", handlers.ArticleHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
