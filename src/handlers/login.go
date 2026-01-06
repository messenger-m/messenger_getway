package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"api-gateway/src/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary      Auth
// @Router       /auth/login [post]
// @Description  proxy to auth service

func LoginHandler(c *fiber.Ctx) error {
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
	resp, err := http.Post("http://127.0.0.1:8000/login", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Auth service error")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Read response error")
	}
	fmt.Printf("Auth service response: %s\n", string(body))

	// _, err = redis_handler.Client.XAdd(redis_handler.Ctx, &redis.XAddArgs{
	// 	Stream: "events:in",
	// 	Values: map[string]interface{}{
	// 		"type": "login",
	// 		"text": string(body),
	// 	},
	// }).Result()

	// if err != nil {
	// 	log.Println("Redis error:", err)
	// }

	return c.Status(resp.StatusCode).Send(body)
}
