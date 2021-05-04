package app

import (
	"bookstore_users-api/controllers/ping"

	"github.com/andresnboza/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/", ping.Ping)

	router.POST("/users/", users.Create)
	router.GET("/users/:user_id", users.Getr)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
}
