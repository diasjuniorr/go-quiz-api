package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/jotajay/go-quiz-app/cmd/db"
	"github.com/jotajay/go-quiz-app/internal/user"
	"github.com/rs/cors"
)

type HealthCheck struct {
	Version string `json:"version"`
	Status  string `json:"status"`
}

// var healthStauts = HealthCheck{Version: "1.0", Status: "ok"}
var database *gorm.DB
var err error

func main() {

	//set api version and port
	port := ":3000"

	r := mux.NewRouter()

	database, err = db.InitializeDB()
	defer database.Close()
	if err != nil {
		panic("failed to initialize db")
	}

	r.HandleFunc("/users", user.CreateUser(database)).Methods("POST")
	r.HandleFunc("/users", user.GetAllUsers(database)).Methods("GET")
	r.HandleFunc("/users/{id}", user.GetUser(database)).Methods("GET")

	fmt.Printf(`Server running and listening on port %v`, port)
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(port, handler))

}

// func checkApiHealth(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(healthStauts)
// }
