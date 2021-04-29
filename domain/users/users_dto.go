package users

import (
	"github.com/andresnboza/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id 					int64	 `json:"id"`
	FirstName 	string `json:"firstName"`
	LastName		string `json:"lastName"`
	Email				string `json:"email"`
	DateCreated string `json:"dateCreated"`
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
