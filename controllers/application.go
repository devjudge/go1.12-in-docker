package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("hello world")
}

type MyRouter mux.Router

func Endpoint() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")
	return r
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r) // call original
	})
}
