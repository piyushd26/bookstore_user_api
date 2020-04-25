package app

import (
	"bookstore/bookstore_user_api/src/controllers/ping"
	"bookstore/bookstore_user_api/src/controllers/user"

	"github.com/gin-gonic/gin"
)

func StartApplication() {
	r := gin.Default()

	// mappings
	r.GET("/ping", ping.Ping)
	r.POST("/users", user.CreateUser)
	r.GET("/users/:user_id", user.GetUser)

	r.Run(":8080")
}
