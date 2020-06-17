package config

import (
	"log"

	"github.com/spf13/viper"
)

var PORT string
var VERIFICATION_TOKEN string
var ACCESS_TOKEN string
var REDIS_ADDRESS string
var REDIS_PASSWORD string
var REDIS_DB_INDEX int

func init() {
	viper.SetConfigName("quiz-kitty-config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("[WARNING] Unable to locate configuration file")
	}

	viper.AutomaticEnv()

	PORT = viper.GetString("port")
	VERIFICATION_TOKEN = viper.GetString("verification_token")
	ACCESS_TOKEN = viper.GetString("access_token")

	REDIS_ADDRESS = viper.GetString("redis.address")
	REDIS_PASSWORD = viper.GetString("redis.password")
	REDIS_DB_INDEX = viper.GetInt("redis.db_index")
}
