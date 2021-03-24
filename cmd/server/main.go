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

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	return
}

var healthStauts = HealthCheck{Version: "1.0", Status: "ok"}

var db *gorm.DB
var err error

func main() {
	r := mux.NewRouter()
	port := ":3000"

	r.HandleFunc("/", checkApiHealth).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")

	fmt.Printf(`Server running and listening on port %v`, port)

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=postgres sslmode=disable password=superpass@123")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(&User{})

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

	db.Create(&user)

	json.NewEncoder(w).Encode(user)

}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params["id"])
}
