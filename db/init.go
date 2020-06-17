package db

import (
	"log"

	"github.com/go-redis/redis"

	"github.com/dryairship/quiz-kitty/config"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.REDIS_ADDRESS,
		Password: config.REDIS_PASSWORD,
		DB:       config.REDIS_DB_INDEX,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("[ERROR] Cannot Ping Redis. Error: %v\n", err)
	}
}
