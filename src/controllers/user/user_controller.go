package user

import (
	"bookstore/bookstore_user_api/src/domain/users"
	"bookstore/bookstore_user_api/src/services"
	"bookstore/bookstore_user_api/utils/errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(string(err.Error()))
		c.JSON(http.StatusBadRequest, errors.BadRequestError(err.Error()))
		return
	}

	fmt.Println(user)

	result, restErr := services.CreateUser(&user)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	fmt.Println(result)

	c.JSON(http.StatusCreated, result)
	return
}

func GetUser(c *gin.Context) {
	// value from the path param
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		c.JSON(400, errors.BadRequestError("invalid user id"))
	}

	result, restErr := services.GetUser(userId)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
