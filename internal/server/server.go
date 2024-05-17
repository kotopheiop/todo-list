package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"todo-list/config"
	"todo-list/internal/controller/interfaces"
	"todo-list/internal/router/routes"
)

func NewServer(cfg *config.Config, handler interfaces.DataHandler) *fiber.App {
	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
		AppName:      cfg.App.Name + " " + cfg.App.Version,
	})

	app.Use(cors.New(cors.Config{
		AllowMethods:     "POST,GET,PUT,DELETE",
		AllowHeaders:     "*",
		AllowCredentials: true,
	}))

	routes.SetupRoutes(app, handler)

	return app
}
