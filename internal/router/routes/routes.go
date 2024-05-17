package routes

import (
	"github.com/gofiber/fiber/v2"
	"todo-list/internal/controller/interfaces"
)

func SetupRoutes(app *fiber.App, handler interfaces.DataHandler) {
	// Создание подроута для '/api'
	apiRouter := app.Group("/api")

	apiRouter.Post("/task", handler.CreateTaskEndpoint)
	apiRouter.Get("/task/:id", handler.GetTaskEndpoint)
	apiRouter.Put("/task/:id", handler.UpdateTaskEndpoint)
	apiRouter.Delete("/task/:id", handler.DeleteTaskEndpoint)
	apiRouter.Put("/task/:id/complete", handler.CompleteTaskEndpoint)

	apiRouter.Get("/tasks", handler.GetAllTasksEndpoint)
}
