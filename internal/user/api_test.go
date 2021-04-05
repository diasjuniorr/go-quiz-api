package user_test

import (
	"log"

	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/jotajay/go-quiz-app/internal/user"
)

func TestCreateUser(t *testing.T) {
	var newUser user.User = user.User{}
	newUser = user.MakeNewUser("user", "user@mail.com", "password")

	req, err := http.NewRequest("POST", "/users")
	if err != nil {
		log.Fatal("failed")
	}

	response := httptest.NewRecorder()
}
