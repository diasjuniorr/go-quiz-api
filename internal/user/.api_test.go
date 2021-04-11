// package user

// import (
// 	"bytes"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestCreateUser(t *testing.T) {
// 	// tCases := []struct {
// 	// 	name string
// 	// 	request
// 	// }
// 	var jsonStr = []byte(`{"Name": "user","Email" : "user@mail.com", "password": "password"}`)

// 	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
// 	if err != nil {
// 		log.Fatal("failed")
// 	}

// 	response := httptest.NewRecorder()

// 	a.Router.ServeHTTP(response, req)

// 	if response.Code != http.StatusCreated {
// 		t.Errorf("response code expected: %v. Got: %v", http.StatusCreated, response.Code)
// 	}

// }
