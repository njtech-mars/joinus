package main

import (
	"embed"
	"server/configs"
	"server/lib"
	"server/routes"

	"github.com/gofiber/fiber/v2"
)

//go:embed all:build
var staticFiles embed.FS

func main() {
	configs.SetupConfig()
	configs.SetupEmail()
	configs.ConnectDb()

	app := fiber.New()

	lib.SetupCors(app)

	routes.EmailRoute(app)
	routes.AdminRoute(app)
	routes.ApplicantsRoute(app)
	routes.ServeStatic(app, staticFiles)

	app.Listen(":" + configs.Config.Port)
}
