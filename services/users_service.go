package services

import "github.com/andresnboza/bookstore_users-api/domin/users"

func CreateUser(user users.User) (*users.User, error){
	return &user, nil
}