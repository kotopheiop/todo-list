package main

import (
	"log"
	"net/http"
	"time"
	"todo-list/cmd/app/routes"

	"github.com/rs/cors"
)

func main() {
	router := routes.NewRouter()

	c := cors.New(cors.Options{
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	srv := &http.Server{
		Handler:      handler,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
