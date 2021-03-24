package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"

	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type HealthCheck struct {
	Version string `json:"version"`
	Status  string `json:"status"`
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var healthStauts = HealthCheck{Version: "1.0", Status: "ok"}

var db *gorm.DB

func main() {
	r := mux.NewRouter()
	port := ":3000"

	r.HandleFunc("/", checkApiHealth).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")

	fmt.Printf(`Server running and listening on port %v`, port)

	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(port, handler))

}

func checkApiHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(healthStauts)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	json.NewEncoder(w).Encode(user)

}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params["id"])
}
