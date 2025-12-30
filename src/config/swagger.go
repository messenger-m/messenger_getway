package config

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func SetupSwagger(app *fiber.App) {
	cfg := swagger.Config{
		FilePath: "./docs/swagger.json",
		Path:     "docs",
		Title:    "Swagger API Docs",
	}

	app.Use(swagger.New(cfg))
}
