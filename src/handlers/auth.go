package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// @Summary      Auth
// @Router       /auth/login [post]
// @Description  proxy to auth service
type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func LoginHandler(c *fiber.Ctx) error {
	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad request")
	}

	fmt.Printf("Received login request: %+v\n", req)

	return c.JSON(fiber.Map{
		"status": "ok",
		"login":  req.Login,
	})
}
