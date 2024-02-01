package routes

import (
	"todo-list/cmd/app/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// TODO: Переделать на запросы по методов GET, POST, PUT, DELETE...
	router.HandleFunc("/api/tasks", handlers.GetAllTasksEndpoint).Methods("GET")
	router.HandleFunc("/api/task", handlers.CreateTaskEndpoint).Methods("POST")
	router.HandleFunc("/api/task/{id}", handlers.GetTaskEndpoint).Methods("GET")
	router.HandleFunc("/api/task/{id}/update", handlers.UpdateTaskEndpoint).Methods("POST")
	router.HandleFunc("/api/task/{id}/delete", handlers.DeleteTaskEndpoint).Methods("POST")
	router.HandleFunc("/api/task/{id}/complete", handlers.CompleteTaskEndpoint).Methods("POST")

	return router
}
