package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is:%v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got parameter ID:%s\n", queryParams["id"][0])
	fmt.Fprintf(w, "Got parameter category:%s\n", queryParams["category"][0])
}
