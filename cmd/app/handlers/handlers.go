package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"todo-list/tools/redis"

	"github.com/gorilla/mux"
)

type Task struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Complete  string    `json:"complete"`
	CreatedAt time.Time `json:"created_at"`
}

func GetAllTasksEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Println("GetAllTasksEndpoint")

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

		task := Task{
			ID:       key,
			Name:     taskMap["name"],
			Complete: taskMap["complete"],
		}
		tasks = append(tasks, task)
	}

	json.NewEncoder(response).Encode(tasks)
}

func CreateTaskEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Println("CreateTaskEndpoint")

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
	task.ID = strconv.FormatInt(taskID, 10)
	task.Complete = "false"
	task.CreatedAt = time.Now()

	err = redis.Client.HMSet(task.ID, map[string]interface{}{
		"name":       task.Name,
		"complete":   task.Complete,
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
	fmt.Println("GetTaskEndpoint")
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

	json.NewEncoder(response).Encode(&Task{ID: params["id"], Name: result["name"], Complete: result["complete"]})
}

func CompleteTaskEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Println("CompleteTaskEndpoint")

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
	fmt.Println("DeleteTaskEndpoint")

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
	fmt.Println("UpdateTaskEndpoint")

	params := mux.Vars(request)

	var updatedTask Task
	err := json.NewDecoder(request.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	err = redis.Client.HMSet(params["id"], map[string]interface{}{
		"name": updatedTask.Name,
	}).Err()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Задача обновлена"))
}
