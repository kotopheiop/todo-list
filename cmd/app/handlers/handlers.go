package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
	"todo-list/tools/redis"

	"github.com/gorilla/mux"
)

// TODO: пересмотреть типы, оч.надо)
type Task struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Complete  bool      `json:"complete"`
	CreatedAt time.Time `json:"created_at"`
}

func GetAllTasksEndpoint(response http.ResponseWriter, request *http.Request) {
	log.Println("GetAllTasksEndpoint")

	keys, err := redis.Client.Keys("*").Result()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	tasks := make([]Task, 0)
	for _, key := range keys {
		dataType, err := redis.Client.Type(key).Result()
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		if dataType != "hash" {
			continue
		}

		taskMap, err := redis.Client.HGetAll(key).Result()
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
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

	json.NewEncoder(response).Encode(tasks)
}

func CreateTaskEndpoint(response http.ResponseWriter, request *http.Request) {
	log.Println("CreateTaskEndpoint")

	var task Task
	err := json.NewDecoder(request.Body).Decode(&task)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	taskID, err := redis.Client.Incr("taskID").Result()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
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
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(task)
}

func GetTaskEndpoint(response http.ResponseWriter, request *http.Request) {
	log.Println("GetTaskEndpoint")
	params := mux.Vars(request)

	result, err := redis.Client.HGetAll(params["id"]).Result()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(result) == 0 {
		http.Error(response, "Task not found", http.StatusNotFound)
		return
	}

	complete, _ := strconv.ParseBool(result["complete"])
	id, _ := strconv.Atoi(params["id"])

	json.NewEncoder(response).Encode(&Task{ID: id, Name: result["name"], Complete: complete})
}

func CompleteTaskEndpoint(response http.ResponseWriter, request *http.Request) {
	log.Println("CompleteTaskEndpoint")

	params := mux.Vars(request)

	currentValue, err := redis.Client.HGet(params["id"], "complete").Result()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	newValue := "true"
	if currentValue == "true" {
		newValue = "false"
	}

	_, err = redis.Client.HSet(params["id"], "complete", newValue).Result()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Статус задачи изменен"))
}

func DeleteTaskEndpoint(response http.ResponseWriter, request *http.Request) {
	log.Println("DeleteTaskEndpoint")

	params := mux.Vars(request)
	_, err := redis.Client.Del(params["id"]).Result()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Задача удалена"))
}

func UpdateTaskEndpoint(response http.ResponseWriter, request *http.Request) {
	log.Println("UpdateTaskEndpoint")

	params := mux.Vars(request)

	var updatedTask Task
	err := json.NewDecoder(request.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	err = redis.Client.HMSet(params["id"], map[string]interface{}{
		"name": updatedTask.Name, // Обновляем толя имя таски
	}).Err()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Задача обновлена"))
}
