package app

import (
	"github.com/gofiber/fiber/v2/log"
	"todo-list/config"
	"todo-list/internal/controller/handlers"
	"todo-list/internal/server"
)

func Run(cfg *config.Config) {
	handler := handlers.NewDataHandler(cfg.Handler.DBClient)
	app := server.NewServer(cfg, handler)

	log.Fatal(app.Listen(":8080"))
}
