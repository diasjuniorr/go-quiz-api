package quiz

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

type RequestError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func CreateQuiz(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var quiz Quiz
		json.NewDecoder(r.Body).Decode(&quiz)

		quizIsAllowed := db.Where("title = ?", quiz.Title).First(&quiz).RecordNotFound()

		if !quizIsAllowed {
			w.WriteHeader(http.StatusConflict)
			err := RequestError{Code: http.StatusConflict, Msg: "A quiz with specified title already exists"}
			json.NewEncoder(w).Encode(err)
			return
		}

		db.Create(&quiz)

		json.NewEncoder(w).Encode(quiz)
	}

}
