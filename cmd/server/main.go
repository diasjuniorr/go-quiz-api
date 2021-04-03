package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/jotajay/go-quiz-app/cmd/db"
	"github.com/jotajay/go-quiz-app/internal/metadata"
	"github.com/jotajay/go-quiz-app/internal/quiz"
	"github.com/jotajay/go-quiz-app/internal/user"
	"github.com/rs/cors"
)

var database *gorm.DB
var err error

func main() {

	port := ":3000"

	r := mux.NewRouter()

	database, err = db.InitializeDB()
	defer database.Close()
	if err != nil {
		panic("failed to initialize db")
	}

	r.HandleFunc("/", metadata.CheckApiHealth).Methods("GET")
	r.HandleFunc("/users", user.CreateUser(database)).Methods("POST")
	r.HandleFunc("/users", user.GetAllUsers(database)).Methods("GET")
	r.HandleFunc("/users/{id}", user.GetUser(database)).Methods("GET")
	r.HandleFunc("/quiz", quiz.CreateQuiz(database)).Methods("POST")

	fmt.Printf(`Server running and listening on port %v`, port)
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(port, handler))

}
