package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nuxzero/go-samples/foo"
)

func main() {
	app := fiber.New()
	foo.SetupFooRoutes(app)
	app.Listen(":8080")
}
