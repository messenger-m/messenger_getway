// @title           Api Gateway

package main

import (
	"api-gateway/src/config"
	"api-gateway/src/handlers"
	redis_handler "api-gateway/src/redisHandler"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	redis_handler.InitRedis()
	config.SetupSwagger(app)
	handlers.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
