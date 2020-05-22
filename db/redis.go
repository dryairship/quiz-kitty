package db

import (
	"log"
)

func SetUserData(user, data string) {
	err := redisClient.Set(user, data, 0).Err()
	if err != nil {
		log.Println("[ERROR] Redis can't set user data. Error: ", err)
	}
}

func GetUserData(user string) string {
	data, err := redisClient.Get(user).Result()
	if err != nil {
		log.Println("[ERROR] Redis can't get user data. User: ", user, ", Error: ", err)
	}
	return data
}
