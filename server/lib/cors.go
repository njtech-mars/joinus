package lib

import (
	"server/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupCors(app *fiber.App) {
	if configs.Config.Cors != "" {
		app.Use(cors.New(cors.Config{
			AllowCredentials: true,
			AllowOrigins:     configs.Config.Cors,
		}))
	}
}
