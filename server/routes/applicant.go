package routes

import (
	"server/controllers"
	"server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ApplicantsRoute(app *fiber.App) {
	app.Get("/api/applicant", middlewares.Limitter(5, 30), controllers.ApplicantQuery)
	app.Post("/api/applicant", middlewares.Limitter(2, 30), controllers.ApplicantCreate)

	app.Get("/api/applicants", middlewares.Auth, controllers.ApplicantsQuery)
	app.Get("/api/applicants/download", middlewares.Auth, controllers.ApplicantsDownload)
}
