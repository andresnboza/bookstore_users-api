package app

import (
	"github.com/gin-gonic/gin"
	"github.com/andresnboza/bookstore_users-api/app"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
