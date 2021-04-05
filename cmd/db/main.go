package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jotajay/go-quiz-app/internal/quiz"
	"github.com/jotajay/go-quiz-app/internal/user"
)

func InitializeDB(conn string) (*gorm.DB, error) {

	var db *gorm.DB
	var err error

	//todo create DATABASE_CONNSTR
	db, err = gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&quiz.Quiz{})

	return db, nil

}
