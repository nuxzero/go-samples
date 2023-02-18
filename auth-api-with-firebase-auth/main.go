package main

import (
	"gofirebase/api"
	"gofirebase/config"
	"gofirebase/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := config.CreateDatabase()
	firebaseAuth := config.SetupFirebase()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("firebaseAuth", firebaseAuth)
	})

	r.Use(middleware.AuthMiddleware)

	r.GET("/users", api.FindUsers)
	r.POST("/users", api.CreateUser)

	r.Run(":3000")
}
