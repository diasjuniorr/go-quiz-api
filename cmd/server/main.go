package main

import (
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

type Hello struct {
	Msg string `json:"msg"`
}

var hello = Hello{Msg: "Hello World!"}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/products", HomeHandler)
	r.HandleFunc("/articles", HomeHandler)
	http.ListenAndServe(":3000", r)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hello)
}
