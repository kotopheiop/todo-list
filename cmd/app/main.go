package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"todo-list/cmd/app/database"
	"todo-list/cmd/app/routes"
)

func main() {
	_, err := database.ConnectDB()
	if err != nil {
		// If unable to connect, panic
		panic("could not connect to db")
	}

	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
		AppName:      "Todo List App v1.0.1",
	})

	app.Use(cors.New(cors.Config{
		AllowMethods:     "POST,GET,PUT,DELETE",
		AllowHeaders:     "*",
		AllowCredentials: true,
	}))

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
