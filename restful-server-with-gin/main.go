package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nuxzero/go-samples/foo"
)

func main() {
	r := setupRoute()
	r.Run(":8080")
}

func setupRoute() *gin.Engine {
	r := gin.Default()

	h := foo.FooHandler{}
	h.Initialize()

	r.GET("/foo", h.GetFoo)
	r.POST("/foo", h.PostFoo)

	return r
}
