package test

import (
	"fmt"

	"os"
	"testing"

	"github.com/jotajay/go-quiz-app/cmd/app"
	"github.com/jotajay/go-quiz-app/internal/user"
)

var a app.AppStruct
var firstUser user.User
var firstUserId string

func TestMain(m *testing.M) {
	a = app.AppStruct{}
	connStr := `port=5432 user=postgres dbname=quiztestdb sslmode=disable password=superpass@123`

	err := a.Initialize(connStr)
	if err != nil {
		fmt.Printf("error initializing: %v\n", err)
		panic("application failed to initialize")
	}
	defer a.Db.Close()

	createNewUser()

	code := m.Run()
	os.Exit(code)

	a.Run(":3000")
}

func createNewUser() {
	firstUser = user.User{Name: "first user", Email: "fu@email.com"}
	result := a.Db.Create(&firstUser)
	if result.Error != nil {
		panic("first user creationg failed")
	}
	firstUserId = firstUser.ID.String()
}
