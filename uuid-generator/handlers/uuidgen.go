package handlers

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

type UUID struct {
	UUID string `json:"uuid"`
}

func (p *UUID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/generate" {
		p.UUID = generateRandomUUID()
		respJSON, _ := json.Marshal(p)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(respJSON)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func generateRandomUUID() string {
	groups := []int{4, 2, 2, 2, 6}
	var u []string
	for _, g := range groups {
		b := make([]byte, g)
		_, err := rand.Read(b)
		if err != nil {
			log.Fatal(err)
		}
		u = append(u, hex.EncodeToString(b))
	}
	return strings.Join(u, "-")
}
