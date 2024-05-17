package main

import (
	"todo-list/config"
	"todo-list/internal/app"
)

func main() {
	app.Run(config.MainConfig)
}
