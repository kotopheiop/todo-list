package routes

import (
	"github.com/gofiber/fiber/v2"
	"todo-list/cmd/app/handlers"
)

func SetupRoutes(app *fiber.App) {
	// Создание подроута для '/api'
	apiRouter := app.Group("/api")

	apiRouter.Post("/task", handlers.CreateTaskEndpoint)
	apiRouter.Get("/task/:id", handlers.GetTaskEndpoint)
	apiRouter.Put("/task/:id", handlers.UpdateTaskEndpoint)
	apiRouter.Delete("/task/:id", handlers.DeleteTaskEndpoint)
	apiRouter.Put("/task/:id/complete", handlers.CompleteTaskEndpoint)

	apiRouter.Get("/tasks", handlers.GetAllTasksEndpoint)
}
