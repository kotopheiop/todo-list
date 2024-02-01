package routes

import (
	"todo-list/cmd/app/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Создание подроута для '/api'
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/task", handlers.CreateTaskEndpoint).Methods("POST")
	apiRouter.HandleFunc("/task/{id}", handlers.GetTaskEndpoint).Methods("GET")
	apiRouter.HandleFunc("/task/{id}", handlers.UpdateTaskEndpoint).Methods("PUT")
	apiRouter.HandleFunc("/task/{id}", handlers.DeleteTaskEndpoint).Methods("DELETE")
	apiRouter.HandleFunc("/task/{id}/complete", handlers.CompleteTaskEndpoint).Methods("PUT")

	router.HandleFunc("/api/tasks", handlers.GetAllTasksEndpoint).Methods("GET")

	return router
}
