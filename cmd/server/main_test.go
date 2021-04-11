package main

import (
	"fmt"

	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var app App

func TestMain(m *testing.M) {
	app = App{}
	connStr := `port=5432 user=postgres dbname=postgres sslmode=disable password=superpass@123`

	err := app.Initialize(connStr)
	if err != nil {
		fmt.Printf("error initializing: %v\n", err)
		panic("application failed to initialize")
	}
	defer app.Db.Close()

	code := m.Run()
	os.Exit(code)

	app.Run(":3000")
}

func TestCreateUser(t *testing.T) {
	var jsonStr = []byte(`{"Name": "user","Email" : "user@mail.com", "password": "password"}`)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal("failed")
	}

	response := httptest.NewRecorder()

	app.Router.ServeHTTP(response, req)

	if response.Code != http.StatusCreated {
		t.Errorf("response code expected: %v. Got %v", http.StatusCreated, response.Code)
	}

}
