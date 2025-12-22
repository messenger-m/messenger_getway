// @title           Api Gateway

package main

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// @Summary      Auth
// @Router       /auth/login [post]
// @Description  proxy to auth service
func loginHandler(c *fiber.Ctx) error {
	jsonData := c.Body()
	bodyReader := bytes.NewReader(jsonData)

	resp, err := http.Post("http://127.0.0.1:8000/auth/add_user", "application/json", bodyReader)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("auth service unavailable")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to read response")
	}

	c.Status(resp.StatusCode)
	return c.Send(body)
}

type Message struct {
	Text string `json:"text"`
}

func main() {
	app := fiber.New()

	cfg := swagger.Config{
		FilePath: "./docs/swagger.json",
		Path:     "docs",
		Title:    "Swagger API Docs",
	}

	app.Use(swagger.New(cfg))

	app.Post("/auth/login", loginHandler)

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		log.Println("WebSocket клиент подключён")

		for {
			var msg Message
			if err := c.ReadJSON(&msg); err != nil {
				log.Println("Клиент отключился:", err)
				break
			}

			log.Println("Получено:", msg.Text)

			msg.Text = "Привет от сервера!"
			if err := c.WriteJSON(msg); err != nil {
				log.Println("Ошибка отправки:", err)
				break
			}
		}
	}))

	log.Fatal(app.Listen(":3000"))
}
