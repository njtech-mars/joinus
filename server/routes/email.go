package routes

import (
	"server/controllers"
	"server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func EmailRoute(app *fiber.App) {
	app.Get("/api/email", middlewares.Auth, controllers.EmailQuery)
	app.Post("/api/email", middlewares.Limitter(2, 30), controllers.EmailSend)
}
