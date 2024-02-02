package handlers

import (
	"encoding/json"
	"log"
	"sort"
	"strconv"
	"time"
	"todo-list/tools/redis"

	"github.com/gofiber/fiber/v2"
)

type Task struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Complete  bool      `json:"complete"`
	CreatedAt time.Time `json:"created_at"`
}

func GetAllTasksEndpoint(c *fiber.Ctx) error {
	log.Println("GetAllTasksEndpoint")

	keys, err := redis.Client.Keys("*").Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	tasks := make([]Task, 0)
	for _, key := range keys {
		dataType, err := redis.Client.Type(key).Result()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		if dataType != "hash" {
			continue
		}

		taskMap, err := redis.Client.HGetAll(key).Result()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		id, _ := strconv.Atoi(key)
		complete, _ := strconv.ParseBool(taskMap["complete"])

		task := Task{
			ID:       id,
			Name:     taskMap["name"],
			Complete: complete,
		}
		tasks = append(tasks, task)
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})

	return c.JSON(tasks)
}

func CreateTaskEndpoint(c *fiber.Ctx) error {
	log.Println("CreateTaskEndpoint")

	var task Task
	err := json.Unmarshal(c.Body(), &task)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	taskID, err := redis.Client.Incr("taskID").Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	task.ID = int(taskID)
	task.Complete = false
	task.CreatedAt = time.Now()

	err = redis.Client.HMSet(strconv.Itoa(task.ID), map[string]interface{}{
		"name":       task.Name,
		"complete":   strconv.FormatBool(task.Complete),
		"created_at": task.CreatedAt.String(),
	}).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}
func GetTaskEndpoint(c *fiber.Ctx) error {
	log.Println("GetTaskEndpoint")

	id := c.Params("id")
	result, err := redis.Client.HGetAll(id).Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if len(result) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}

	complete, _ := strconv.ParseBool(result["complete"])
	taskID, _ := strconv.Atoi(id)

	return c.JSON(&Task{ID: taskID, Name: result["name"], Complete: complete})
}

func CompleteTaskEndpoint(c *fiber.Ctx) error {
	log.Println("CompleteTaskEndpoint")

	id := c.Params("id")
	currentValue, err := redis.Client.HGet(id, "complete").Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	newValue := "true"
	if currentValue == "true" {
		newValue = "false"
	}

	_, err = redis.Client.HSet(id, "complete", newValue).Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendString("Статус задачи изменен")
}

func DeleteTaskEndpoint(c *fiber.Ctx) error {
	log.Println("DeleteTaskEndpoint")

	id := c.Params("id")
	_, err := redis.Client.Del(id).Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendString("Задача удалена")
}

func UpdateTaskEndpoint(c *fiber.Ctx) error {
	log.Println("UpdateTaskEndpoint")

	id := c.Params("id")

	var updatedTask Task
	err := json.Unmarshal(c.Body(), &updatedTask)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = redis.Client.HMSet(id, map[string]interface{}{
		"name": updatedTask.Name, // Обновляем толя имя таски
	}).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendString("Задача обновлена")
}
