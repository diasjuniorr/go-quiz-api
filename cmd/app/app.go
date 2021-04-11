package app

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

type AppStruct struct {
	Db     *gorm.DB
	Router *mux.Router
}

func (a *AppStruct) Initialize(connStr string) error {
	a.Db, err = db.InitializeDB(connStr)
	if err != nil {
		return err
	}

	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/", metadata.CheckApiHealth).Methods("GET")
	a.Router.HandleFunc("/users", user.CreateUser(a.Db)).Methods("POST")
	a.Router.HandleFunc("/users", user.GetAllUsers(a.Db)).Methods("GET")
	a.Router.HandleFunc("/users/{id}", user.GetUser(a.Db)).Methods("GET")
	a.Router.HandleFunc("/quiz", quiz.CreateQuiz(a.Db)).Methods("POST")

	return nil

}

func (a *AppStruct) Run(port string) {
	fmt.Printf(`Server running and listening on port %v`, port)
	handler := cors.Default().Handler(a.Router)
	log.Fatal(http.ListenAndServe(port, handler))
}
