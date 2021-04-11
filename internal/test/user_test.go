package test

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jotajay/go-quiz-app/internal/user"
)

func TestUsers(t *testing.T) {
	tCases := []struct {
		name               string
		method             string
		resource           string
		data               string
		expectedStatusCode int
		expectedResponse   user.User
	}{{
		name:               "create_new_user",
		method:             "POST",
		resource:           "/users",
		data:               `{"Name": "user","Email" : "user@mail.com", "password": "password"}`,
		expectedStatusCode: 201,
		expectedResponse:   user.User{Name: "name", Email: "email@email.com"},
	}, {
		name:               "get_user_by_id",
		method:             "GET",
		resource:           (`/users/` + firstUserId),
		data:               firstUserId,
		expectedStatusCode: 200,
		expectedResponse:   firstUser,
	}, {
		name:               "get_users",
		method:             "GET",
		resource:           "/users",
		data:               "",
		expectedStatusCode: 200,
		expectedResponse:   firstUser,
	}}

	for _, test := range tCases {
		t.Run(test.name, func(t *testing.T) {
			var jsonStr = []byte(test.data)

			log.Printf("started")
			req, err := http.NewRequest(test.method, test.resource, bytes.NewBuffer(jsonStr))
			if err != nil {
				log.Fatal("failed")
			}

			response := httptest.NewRecorder()

			a.Router.ServeHTTP(response, req)

			if response.Code != test.expectedStatusCode {
				t.Errorf("response code expected: %v. Got: %v", test.expectedStatusCode, response.Code)
			}

			log.Printf("finished")
		})
	}

}
