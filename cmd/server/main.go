package main

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

type Hello struct {
	Msg string `json:"msg"`
}

type Account struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var hello = Hello{Msg: "Hello World!"}

func main() {
	r := mux.NewRouter()
	port := ":3000"

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/articles", HomeHandler)

	fmt.Printf(`Server running and listening on port %v`, port)
	http.ListenAndServe(port, r)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hello)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Printf("params %v", params)
}
