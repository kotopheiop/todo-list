package redis

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func init() {
	redisPass := os.Getenv("REDIS_PASSWORD")
	redisPort := os.Getenv("REDIS_PORT")
	Client = redis.NewClient(&redis.Options{
		Addr:     "redis:" + redisPort,
		Password: redisPass,
		DB:       0,
	})
	err := Client.Ping().Err()
	if err == nil {
		log.Println("Redis client has been successfully created")
	} else {
		log.Fatal(err)
	}

}
