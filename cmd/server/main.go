package main

import (
	"fmt"
	"os"

	"github.com/jotajay/go-quiz-app/cmd/app"
)

func main() {
	a := app.AppStruct{}

	//TODO use fmt.Printf for dynamic variables
	connStr := os.Getenv("DATABASE_URL")
	// connStr := `port=5432 user=postgres dbname=postgres sslmode=disable password=superpass@123`

	err := a.Initialize(connStr)
	if err != nil {
		fmt.Printf("error initializing: %v\n", err)
		panic("application failed to initialize")
	}
	defer a.Db.Close()

	a.Run(":3000")
}
