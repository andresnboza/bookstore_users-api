package users

import (
	"bookstore_users-api/utils/errors"
	"strings"
)

const (
	StatusActive = "active"
)

type User struct {
	Id 			int64  `json:"id"`
	FirstName 	string `json:"first_name"`
	LastName	string `json:"last_name"`
	Email		string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"-"`
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
	
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	if user.FirstName == "" {
		return errors.NewBadRequestError("invalid FirstName address")
	}

	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	if user.LastName == "" {
		return errors.NewBadRequestError("invalid LastName address")
	}
	
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password address")
	}
	
	return nil
}
