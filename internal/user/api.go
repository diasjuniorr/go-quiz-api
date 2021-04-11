package user

import (
	"encoding/json"

	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"
)

type RequestError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func CreateUser(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		userIsAllowed := db.Where("email = ?", user.Email).First(&user).RecordNotFound()

		// u := MakeNewUser(user)
		if !userIsAllowed {
			w.WriteHeader(http.StatusConflict)
			err := RequestError{Code: http.StatusConflict, Msg: "User already exists"}
			json.NewEncoder(w).Encode(err)
			return
		}

		db.Create(&user)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}

}

func GetAllUsers(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		var users []User
		result := db.Find(&users)

		if result.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(result.Error)
			return
		}

		json.NewEncoder(w).Encode(result.Value)
	}
}

func GetUser(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)

		var user User

		result := db.Where("ID = ?", params["id"]).First(&user)

		if result.Error == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(result.Error)
			return
		}

		if result.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(result.Error)
			return
		}

		json.NewEncoder(w).Encode(result.Value)
	}
}
