package db

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jotajay/go-quiz-app/cmd/db/repository"
)

func initializeDB() *gorm.DB {

	var db *gorm.DB
	var err error

	//todo create DATABASE_CONNSTR
	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=postgres sslmode=disable password=superpass@123")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(&repository.User{})

	return db

}
