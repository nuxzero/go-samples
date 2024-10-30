package foo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Foo struct {
	ID  int    `json:"id"`
	Bar string `json:"bar"`
}

type FooHandler struct {
}

func (h *FooHandler) Initialize() {
	// TODO: Initialize
}

func (h *FooHandler) GetFoo(c *gin.Context) {
	f := Foo{
		ID:  1,
		Bar: "Hello, World!",
	}
	c.JSON(http.StatusOK, f)
}

func (h *FooHandler) PostFoo(c *gin.Context) {
	var f Foo
	if err := c.ShouldBindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, f)
}
