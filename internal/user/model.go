package user

import (
	"github.com/jinzhu/gorm"
	"github.com/jotajay/go-quiz-app/internal/entity"
	"github.com/jotajay/go-quiz-app/pkg/util"
)

type User struct {
	entity.Base
	Name     string `json:"name"`
	Email    string `json:"email"`
	password string
}

// func MakeNewUser(user PostUserReq) User {
// 	return User{
// 		Name:     user.Name,
// 		Email:    user.Email,
// 		password: user.Password,
// 	}
// }

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hash, err := util.HashPassword(u.password)
	if err != nil {
		panic("pass encryption went wrong")
	}

	u.password = hash
	return
}
