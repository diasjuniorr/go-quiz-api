package quiz

import (
	"github.com/jotajay/go-quiz-app/internal/entity"
)

type Quiz struct {
	entity.Base
	Subject     string     `json:"subject"`
	Category    string     `json:"category"`
	Subcategory string     `json:"subcategory"`
	Title       string     `json:"title"`
	Level       int        `json:"level"`
	Questions   []Question `json:"questions"`
}

type Question struct {
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct"`
	IncorrectAnswers []string `json:"incorrects"`
}
