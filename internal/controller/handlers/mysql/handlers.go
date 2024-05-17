package mysql

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
	"todo-list/internal/controller/interfaces"
	"todo-list/internal/mysql"
)

type Handler struct {
}

var _ interfaces.DataHandler = &Handler{}

func (h *Handler) GetAllTasksEndpoint(c *fiber.Ctx) error {
	log.Println("mysql - GetAllTasksEndpoint")

	var tasks []mysql.Task
	if err := mysql.DB.Find(&tasks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(tasks)
}

func (h *Handler) CreateTaskEndpoint(c *fiber.Ctx) error {
	log.Println("mysql - CreateTaskEndpoint")

	var task mysql.Task
	err := json.Unmarshal(c.Body(), &task)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	log.Println(task)
	if err := mysql.DB.Create(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *Handler) GetTaskEndpoint(c *fiber.Ctx) error {
	log.Println("mysql - GetTaskEndpoint")

	id := c.Params("id")
	var task mysql.Task
	if err := mysql.DB.First(&task, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(task)
}

func (h *Handler) CompleteTaskEndpoint(c *fiber.Ctx) error {
	log.Println("mysql - CompleteTaskEndpoint")

	id := c.Params("id")
	var task mysql.Task
	if err := mysql.DB.First(&task, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	task.Complete = !task.Complete
	if err := mysql.DB.Save(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendString("Статус задачи изменен")
}

func (h *Handler) DeleteTaskEndpoint(c *fiber.Ctx) error {
	log.Println("mysql - DeleteTaskEndpoint")

	id := c.Params("id")
	if err := mysql.DB.Delete(&mysql.Task{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendString("Задача удалена")
}

func (h *Handler) UpdateTaskEndpoint(c *fiber.Ctx) error {
	log.Println("mysql - UpdateTaskEndpoint")

	id := c.Params("id")

	var updatedTask mysql.Task
	err := json.Unmarshal(c.Body(), &updatedTask)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var task mysql.Task
	if err := mysql.DB.First(&task, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	task.Name = updatedTask.Name // Обновляем только имя задачи

	if err := mysql.DB.Save(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendString("Задача обновлена")
}
