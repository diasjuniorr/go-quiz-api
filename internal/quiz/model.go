package quiz

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jotajay/go-quiz-app/internal/entity"
)

type Quiz struct {
	entity.Base
	Subject     string     `json:"subject"`
	Category    string     `json:"category"`
	Subcategory string     `json:"subcategory"`
	Title       string     `json:"title" gorm:"unique;not null"`
	Level       int        `json:"level"`
	Questions   []Question `gorm:"not null" json:"questions" `
}

type Question struct {
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct"`
	IncorrectAnswers []string `json:"incorrects"`
}

func (q *Quiz) BeforeCreate(tx *gorm.DB) error {
	var err error
	if len(q.Subject) < 3 {
		err = fmt.Errorf("subject: must be at least 3 char long")
	}
	if err != nil {
		return err
	}

	if len(q.Category) < 3 {
		err = fmt.Errorf("category: must be at least 3 char long")
	}
	if err != nil {
		return err
	}

	if len(q.Subcategory) < 3 {
		err = fmt.Errorf("subcategory: must be at least 3 char long")
	}
	if err != nil {
		return err
	}

	if len(q.Title) < 3 {
		err = fmt.Errorf("title: must be at least 3 char long")
	}
	if err != nil {
		return err
	}

	if (q.Level) < 100 {
		err = fmt.Errorf("level: must be a number between 100 and 999")
	}
	if err != nil {
		return err
	}

	if len(q.Questions) < 1 {
		err = fmt.Errorf("questions: must have at least one question")
	}
	if err != nil {
		return err
	}

	return nil
}
