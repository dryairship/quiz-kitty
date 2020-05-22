package db

import (
	"encoding/json"
	"log"

	"github.com/dryairship/messenger-quiz-bot/models"
)

func SetRedisUserData(user string, data *models.RedisUserData) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("[ERROR] Can't JSON marshal Redis user data. Error: ", err)
		return
	}

	err = redisClient.Set(user, jsonData, 0).Err()
	if err != nil {
		log.Println("[ERROR] Redis can't set user data. Error: ", err)
	}
}

func GetRedisUserData(user string) (*models.RedisUserData, error) {
	jsonData, err := redisClient.Get(user).Result()
	if err != nil {
		log.Println("[ERROR] Redis can't get user data. User: ", user, ", Error: ", err)
		return nil, err
	}

	var data models.RedisUserData
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Println("[ERROR] Cannot unmarshal Redis bytes to Redis user data. User: ", user, ", Error: ", err)
		return nil, err
	}

	return &data, nil
}
