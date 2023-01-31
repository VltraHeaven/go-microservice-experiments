package main

import (
	"fmt"
	"github.com/vltraheaven/go-microservice-experiments/uuid-generator/handlers"
	"log"
	"net/http"
)

func main() {
	mux := &handlers.UUID{}
	port := ":9091"
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
