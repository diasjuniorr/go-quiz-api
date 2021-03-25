package user

import (
	"github.com/jinzhu/gorm"

	"github.com/jotajay/go-quiz-app/pkg/util"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	password string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hash, err := util.HashPassword(u.password)
	if err != nil {
		panic("pass encryption went wrong")
	}

	u.password = hash
	return
}
