package handlers

import (
	"api-gateway/src/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// @Summary      Auth
// @Router       /auth/login [post]
// @Description  proxy to auth service

func RegisterHandler(c *fiber.Ctx) error {
	var req models.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad request")
	}

	fmt.Printf("Received login request: %+v\n", req)

	// Сериализуем req в JSON
	reqBody, err := json.Marshal(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("JSON marshal error")
	}
	fmt.Printf("Serialized request body: %s\n", string(reqBody))

	// Делаем POST-запрос
	resp, err := http.Post("http://127.0.0.1:8000/register", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Auth service error")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Read response error")
	}
	fmt.Printf("Auth service response: %s\n", string(body))
	return c.Status(resp.StatusCode).Send(body)
}
