package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"github.com/user/bookstore_users-api/domain/users"
	"github.com/user/bookstore_users-api/services"
	"github.com/user/bookstore_users-api/utils/errors"
)

func GetUser(c *gin.Context) {
	user_id, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		err := errors.NewBadRequestError("user_id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(user_id)
	
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User

	// Getting the json representation of the user
	// and validation the json representation
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	// Saving the user
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	// Returning the recently created succesfull user
	c.JSON(http.StatusCreated, result)
	return
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
