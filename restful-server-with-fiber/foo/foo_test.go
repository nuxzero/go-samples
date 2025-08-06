package foo

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetFoo(t *testing.T) {
	app := fiber.New()

	app.Get("/foo", GetFoo)

	req, _ := http.NewRequest("GET", "/foo", nil)
	resp, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var foo Foo
	err = json.NewDecoder(resp.Body).Decode(&foo)
	assert.Nil(t, err)
	assert.Equal(t, Foo{ID: 1, Bar: "bar"}, foo)
}
func TestCreateFoo(t *testing.T) {
	app := fiber.New()

	app.Post("/foo", CreateFoo)

	t.Run("valid request", func(t *testing.T) {
		payload := `{"id": 2, "bar": "baz"}`
		req, _ := http.NewRequest("POST", "/foo", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var foo Foo
		err = json.NewDecoder(resp.Body).Decode(&foo)
		assert.Nil(t, err)
		assert.Equal(t, Foo{ID: 2, Bar: "baz"}, foo)
	})

	t.Run("invalid request", func(t *testing.T) {
		payload := `{"id": "invalid", "bar": "baz"}`
		req, _ := http.NewRequest("POST", "/foo", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		var response map[string]string
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.Nil(t, err)
		assert.Contains(t, response["error"], "cannot unmarshal")
	})
}
