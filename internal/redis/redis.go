package redis

import (
	"log"
	"todo-list/config"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func NewClient() *redis.Client {
	redisPass := config.MainConfig.Redis.Password
	redisPort := config.MainConfig.Redis.Port

	Client = redis.NewClient(&redis.Options{
		Addr:     "redis:" + redisPort,
		Password: redisPass,
		DB:       0,
	})
	err := Client.Ping().Err()
	if err == nil {
		log.Println("Клиент Redis успешно создан")
	} else {
		log.Fatal(err)
	}
	return Client
}
