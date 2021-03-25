package api 

import (

	"net/http"

	"github.com/gorilla/mux"

	"fmt"

	"log"

	"github.com/rs/cors"



)

func RegisterHandlers(services *services.Handlers) {

	r := mux.NewRouter()
	port := ":3000"

	r.HandleFunc("/", checkApiHealth).Methods("GET")
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")

	fmt.Printf(`Server running and listening on port %v`, port)
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(port, handler))
}