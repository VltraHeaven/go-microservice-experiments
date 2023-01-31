package main

import (
	"fmt"
	"github.com/vltraheaven/go-microservice-experiments/fastest-mirror/handlers"
	"log"
	"net/http"
	"time"
)

func main() {
	// Instantiate handler for "fastest-mirror" route
	http.HandleFunc("/fastest-mirror", handlers.FindFastestHandler)

	// Configure and initialize API server
	port := ":9090"
	server := &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}
