package test

import (
	"fmt"

	"os"
	"testing"

	"github.com/jotajay/go-quiz-app/cmd/app"
)

var a app.AppStruct

func TestMain(m *testing.M) {
	a = app.AppStruct{}
	connStr := `port=5432 user=postgres dbname=quiztestdb sslmode=disable password=superpass@123`

	err := a.Initialize(connStr)
	if err != nil {
		fmt.Printf("error initializing: %v\n", err)
		panic("application failed to initialize")
	}
	defer a.Db.Close()

	code := m.Run()
	os.Exit(code)

	a.Run(":3000")
}
