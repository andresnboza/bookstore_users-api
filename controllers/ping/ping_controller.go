package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ServerHello(c *gin.Context) {
	c.String(http.StatusOK, "The server is working!! :)")
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}