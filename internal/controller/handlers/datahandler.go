package handlers

import (
	"log"
	"todo-list/internal/controller/handlers/mysql"
	"todo-list/internal/controller/handlers/redis"
	"todo-list/internal/controller/interfaces"
	MysqlClient "todo-list/internal/mysql"
	RedisClient "todo-list/internal/redis"
)

func NewDataHandler(dbClient string) interfaces.DataHandler {
	switch dbClient {
	case "mysql":
		MysqlClient.NewClient()
		return &mysql.Handler{}
	case "redis":
		RedisClient.NewClient()
		return &redis.Handler{}
	default:
		log.Fatalf("Не поддерживаемый клент базы данных: %s", dbClient)
		return nil
	}
}
