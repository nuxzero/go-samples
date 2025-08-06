package foo

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Foo struct {
	ID  int    `json:"id"`
	Bar string `json:"bar"`
}

func SetupFooRoutes(app *fiber.App) {
	foo := app.Group("/foo")

	foo.Get("/", GetFoo)
	foo.Post("/", CreateFoo)
}

func GetFoo(c *fiber.Ctx) error {
	return c.JSON(Foo{
		ID:  1,
		Bar: "bar",
	})
}

func CreateFoo(c *fiber.Ctx) error {
	var foo Foo
	err := c.BodyParser(&foo)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(foo)
}
