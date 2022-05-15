package users

import (
	"github.com/sharkx018/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id          int64  `json:"id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Email       string `json:"email,omitempty"`
	DateCreated string `json:"date_created,omitempty"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func Validate(user *User) *errors.RestErr {
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
