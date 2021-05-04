package users

import (
	"github.com/andresnboza/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id 			int64  `json:"id"`
	FirstName 	string `json:"first_name"`
	LastName	string `json:"last_name"`
	Email		string `json:"email"`
	DateCreated string `json:"date_created"`
}

// Function 
// Called using: users.ValidateUser($user)
func Validate(user *User) *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "asd" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}

// Map over the user instance
// Called using: user.Validate()
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
