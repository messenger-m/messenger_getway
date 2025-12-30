package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/login", LoginHandler)
	app.Get("/ws", websocket.New(WebSocketHandler))

}
