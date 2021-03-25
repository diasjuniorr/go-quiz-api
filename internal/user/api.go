package user

import (
	"encoding/json"

	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
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

func getUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var user User

	result := db.First(&user, params["id"])

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
