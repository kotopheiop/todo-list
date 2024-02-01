package routes

import (
	"todo-list/cmd/app/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Создание подроута для '/api/task'
	taskRouter := router.PathPrefix("/api/task").Subrouter()

	taskRouter.HandleFunc("", handlers.CreateTaskEndpoint).Methods("POST")
	taskRouter.HandleFunc("/{id}", handlers.GetTaskEndpoint).Methods("GET")
	taskRouter.HandleFunc("/{id}", handlers.UpdateTaskEndpoint).Methods("PUT")
	taskRouter.HandleFunc("/{id}", handlers.DeleteTaskEndpoint).Methods("DELETE")
	taskRouter.HandleFunc("/{id}/complete", handlers.CompleteTaskEndpoint).Methods("PUT")

	router.HandleFunc("/api/tasks", handlers.GetAllTasksEndpoint).Methods("GET")

	return router
}
