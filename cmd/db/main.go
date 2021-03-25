package db

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jotajay/go-quiz-app/internal/user"
)

func InitializeDB() (*gorm.DB, error) {

	var db *gorm.DB
	var err error

	//todo create DATABASE_CONNSTR
	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=postgres sslmode=disable password=superpass@123")
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&user.User{})

	return db, nil

}
