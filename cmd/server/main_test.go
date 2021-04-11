package main

import (
	"fmt"

	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
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

func TestCreateUser(t *testing.T) {
	// tCases := []struct {
	// 	name string
	// 	request
	// }
	var jsonStr = []byte(`{"Name": "user","Email" : "user@mail.com", "password": "password"}`)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal("failed")
	}

	response := httptest.NewRecorder()

	a.Router.ServeHTTP(response, req)

	if response.Code != http.StatusCreated {
		t.Errorf("response code expected: %v. Got: %v", http.StatusCreated, response.Code)
	}

}
