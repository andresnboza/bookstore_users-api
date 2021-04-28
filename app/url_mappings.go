package app

import (
	"github.com/andresnboza/bookstore_users-api/controllers/ping"
	"github.com/andresnboza/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/", ping.ServerHello)

	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", users.SearchUser)
	router.POST("/users/", users.CreateUser)
}
