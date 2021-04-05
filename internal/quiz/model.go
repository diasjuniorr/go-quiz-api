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
	Title       string     `json:"title"`
	Level       int        `json:"level"`
	Questions   []Question `json:"questions" `
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

func validateSubject(s string, minLen int) error {
	if len(s) < minLen {
		return fmt.Errorf("subject should be longer than %v characters", minLen)
	}
	return nil
}

func validateCategory(s string, minLen int) error {
	if len(s) < minLen {
		return fmt.Errorf("category should be longer than %v characters", minLen)
	}
	return nil
}

func validateSubCategory(s string, minLen int) error {
	if len(s) < minLen {
		return fmt.Errorf("subcategory should be longer than %v characters", minLen)
	}
	return nil
}

func validateTitle(s string, minLen int) error {
	if len(s) < minLen {
		return fmt.Errorf("title should be longer than %v characters", minLen)
	}
	return nil
}

func validateLevel(level int, minLevel int) error {
	if level < minLevel {
		return fmt.Errorf("level should be an int greater than %v", minLevel)
	}
	return nil
}

func validateQuestions(q []Question, minLen int) error {
	if len(q) > minLen {
		return fmt.Errorf("questions array should contain more than %v question(s)", minLen)
	}
	return nil
}
