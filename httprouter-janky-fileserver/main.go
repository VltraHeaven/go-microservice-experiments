package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/vltraheaven/go-microservice-experiments/httprouter-janky-fileserver/handlers"
	"log"
	"net/http"
)

func main() {
	// Create router
	router := httprouter.New()
	// Set routes
	router.GET("/api/v1/go-version", handlers.GetGoVersion)
	router.GET("/api/v1/show-file/:name", handlers.GetFileContent)
	log.Fatal(http.ListenAndServe(":8080", router))
}
