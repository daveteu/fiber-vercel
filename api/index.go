package handler

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	adaptor.FiberHandler(greet).ServeHTTP(w, r)
}

func greet(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}
