package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/andresnboza/bookstore_users-api/domain/users"
	"io/ioutil"
	"fmt"
	"github.com/andresnboza/bookstore_users-api/services"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func CreateUser(c *gin.Context) {
	var user users.User

	// Getting the json representation of the user
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		// TODO: return bad request to the caller
		return
	}

	// Saving the user
	result, saveErr = services.CreateUser(user)
	if saveErr != nil {
		//TODO: handle user creation error
		return
	}

	// Returning the user recently created
	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
