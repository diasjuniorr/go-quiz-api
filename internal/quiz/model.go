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

	err = validateSubject(q.Subject, 3)
	if err != nil {
		return err
	}

	err = validateCategory(q.Category, 3)
	if err != nil {
		return err
	}

	err = validateSubCategory(q.Subcategory, 3)
	if err != nil {
		return err
	}

	err = validateTitle(q.Title, 3)
	if err != nil {
		return err
	}

	err = validateLevel(q.Level, 100)
	if err != nil {
		return err
	}

	err = validateQuestions(q.Questions, 1)
	if err != nil {
		return err
	}

	return nil
}

func validateSubject(s string, l int) error {
	if len(s) < l {
		return fmt.Errorf("should be longer than %v characters", l)
	}
	return nil
}

func validateCategory(s string, l int) error {
	if len(s) < l {
		return fmt.Errorf("should be longer than %v characters", l)
	}
	return nil
}

func validateSubCategory(s string, l int) error {
	if len(s) < l {
		return fmt.Errorf("should be longer than %v characters", l)
	}
	return nil
}

func validateTitle(s string, l int) error {
	if len(s) < l {
		return fmt.Errorf("should be longer than %v characters", l)
	}
	return nil
}

func validateLevel(level int, limit int) error {
	if level > limit {
		return fmt.Errorf("should be an int greater than %v", limit)
	}
	return nil
}

func validateQuestions(q []Question, l int) error {
	if len(q) > 1 {
		return fmt.Errorf("array should contain more than %v question(s)", l)
	}
	return nil
}
