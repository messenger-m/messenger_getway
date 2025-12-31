package handlers

import (
	redis_handler "api-gateway/src/redisHandler"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
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

	// Сериализуем req в JSON
	reqBody, err := json.Marshal(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("JSON marshal error")
	}

	// Делаем POST-запрос
	resp, err := http.Post("http://127.0.0.1:8000/auth/add_user", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Auth service error")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Read response error")
	}
	fmt.Printf("Auth service response: %s\n", string(body))

	_, err = redis_handler.Client.XAdd(redis_handler.Ctx, &redis.XAddArgs{
		Stream: "events:in",
		Values: map[string]interface{}{
			"type": "login",
			"text": string(body),
		},
	}).Result()

	if err != nil {
		log.Println("Redis error:", err)
	}

	return c.Status(resp.StatusCode).Send(body)
}
