package interfaces

import "github.com/gofiber/fiber/v2"

type DataHandler interface {
	GetAllTasksEndpoint(c *fiber.Ctx) error
	CreateTaskEndpoint(c *fiber.Ctx) error
	GetTaskEndpoint(c *fiber.Ctx) error
	CompleteTaskEndpoint(c *fiber.Ctx) error
	DeleteTaskEndpoint(c *fiber.Ctx) error
	UpdateTaskEndpoint(c *fiber.Ctx) error
}
