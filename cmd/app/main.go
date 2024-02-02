package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"todo-list/cmd/app/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowMethods:     "POST,GET,PUT,DELETE",
		AllowHeaders:     "*",
		AllowCredentials: true,
	}))

	routes.SetupRoutes(app)

	log.Println("Server start")

	log.Fatal(app.Listen(":8080"))
}
