package handlers

import (
	"encoding/json"
	"github.com/vltraheaven/go-microservice-experiments/fastest-mirror/mirrors"
	"github.com/vltraheaven/go-microservice-experiments/fastest-mirror/requests"
	"net/http"
)

func FindFastestHandler(writer http.ResponseWriter, request *http.Request) {
	response := requests.FindFastest(mirrors.MirrorList)
	respJSON, _ := json.Marshal(response)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(respJSON)
}
