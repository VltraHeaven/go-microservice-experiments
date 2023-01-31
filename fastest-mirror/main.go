package main

import (
	"encoding/json"
	"fmt"
	"github.com/vltraheaven/go-microservice-experiments/fastest-mirror/client"
	"github.com/vltraheaven/go-microservice-experiments/fastest-mirror/mirrors"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/fastest-mirror", func(writer http.ResponseWriter, request *http.Request) {
		response := client.FindFastest(mirrors.MirrorList)
		respJSON, _ := json.Marshal(response)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(respJSON)
	})
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
