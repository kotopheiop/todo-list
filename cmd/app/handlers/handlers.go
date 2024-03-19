package handlers

import (
	"log"
	"strconv"
	"todo-list/cmd/app/database"
	"todo-list/cmd/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllTasksEndpoint(c *fiber.Ctx) error {
	log.Println("GetAllTasksEndpoint")

	var tasks []models.Task
	database.DB.Find(&tasks)

	return c.JSON(tasks)
}

func CreateTaskEndpoint(c *fiber.Ctx) error {
	log.Println("CreateTaskEndpoint")

	var task models.Task
	err := c.BodyParser(&task)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	task.Complete = false

	database.DB.Create(&task)

	return c.Status(fiber.StatusCreated).JSON(task)
}

func GetTaskEndpoint(c *fiber.Ctx) error {
	log.Println("GetTaskEndpoint")

	id, _ := strconv.Atoi(c.Params("id"))

	var task models.Task
	database.DB.First(&task, id)

	if task.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}

	return c.JSON(task)
}

func CompleteTaskEndpoint(c *fiber.Ctx) error {
	log.Println("CompleteTaskEndpoint")

	id, _ := strconv.Atoi(c.Params("id"))

	var task models.Task
	database.DB.First(&task, id)

	if task.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}

	task.Complete = !task.Complete

	database.DB.Save(&task)

	return c.SendString("Статус задачи изменен")
}

func DeleteTaskEndpoint(c *fiber.Ctx) error {
	log.Println("DeleteTaskEndpoint")
	log.Println(c.Request())
	id, _ := strconv.Atoi(c.Params("id"))

	var task models.Task
	database.DB.First(&task, id)

	if task.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}

	database.DB.Delete(&task)

	return c.SendString("Задача удалена")
}

func UpdateTaskEndpoint(c *fiber.Ctx) error {
	log.Println("UpdateTaskEndpoint")

	id, _ := strconv.Atoi(c.Params("id"))

	var task models.Task
	database.DB.First(&task, id)

	if task.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}

	var updatedTask models.Task
	err := c.BodyParser(&updatedTask)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	task.Name = updatedTask.Name

	database.DB.Save(&task)

	return c.SendString("Задача обновлена")
}
