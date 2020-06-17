package db

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/dryairship/quiz-kitty/config"
)

var redisClient *redis.Client
var mongoClient *mongo.Client

func connectToRedis() bool {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.REDIS_ADDRESS,
		Password: config.REDIS_PASSWORD,
		DB:       config.REDIS_DB_INDEX,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Printf("[ERROR] Cannot Ping Redis. Error: %v\n", err)
		return false
	} else {
		log.Println("[INFO] Successfully pinged Redis")
		return true
	}
}

func connectToMongo() bool {
	protocol := "mongodb"
	if config.MONGO_DNS_SRV {
		protocol += "+srv"
	}

	authString := ""
	if config.MONGO_USING_AUTH {
		authString = fmt.Sprintf(
			"%s:%s@",
			url.QueryEscape(config.MONGO_USER),
			url.QueryEscape(config.MONGO_PASSWORD),
		)
	}

	connectURL := fmt.Sprintf(
		"%s://%s%s/%s%s",
		protocol,
		authString,
		config.MONGO_ADDRESS,
		config.MONGO_DB_NAME,
		config.MONGO_EXTRA_OPTIONS,
	)

	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(connectURL))
	if err != nil {
		log.Printf("[ERROR] Cannot connect to Mongo. Error: %v\n", err)
		return false
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Printf("[ERROR] Cannot Ping Mongo. Error: %v\n", err)
		return false
	} else {
		log.Println("[INFO] Successfully pinged Mongo")
		return true
	}
}

func init() {
	for !connectToRedis() {
		log.Println("[INFO] Retrying to connect to Redis")
	}
	for !connectToMongo() {
		log.Println("[INFO] Retrying to connect to Mongo")
	}
}
