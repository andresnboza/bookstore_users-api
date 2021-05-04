package app

import (
	"github.com/andresnboza/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.POST("/users/", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.updateUser)
	router.PATCH("/users/:user_id", users.updateUser)
}
