package main

import (
	"encoding/json"
	"fmt"
	"github.com/vltraheaven/go-microservice-experiments/fastest-mirror/mirrors"
	"github.com/vltraheaven/go-microservice-experiments/fastest-mirror/requests"
	"log"
	"net/http"
	"time"
)

func main() {
	// Instantiate handler for "fastest-mirror" route
	http.HandleFunc("/fastest-mirror", func(writer http.ResponseWriter, request *http.Request) {
		response := requests.FindFastest(mirrors.MirrorList)
		respJSON, _ := json.Marshal(response)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(respJSON)
	})

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
