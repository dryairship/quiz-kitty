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

var MONGO_ADDRESS string
var MONGO_DNS_SRV bool
var MONGO_USING_AUTH bool
var MONGO_DB_NAME string
var MONGO_USER string
var MONGO_PASSWORD string
var MONGO_EXTRA_OPTIONS string

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

	MONGO_ADDRESS = viper.GetString("mongo.address")
	MONGO_DNS_SRV = viper.GetBool("mongo.dns_srv")
	MONGO_USING_AUTH = viper.GetBool("mongo.using_auth")
	MONGO_DB_NAME = viper.GetString("mongo.db_name")
	MONGO_USER = viper.GetString("mongo.user")
	MONGO_PASSWORD = viper.GetString("mongo.password")
	MONGO_EXTRA_OPTIONS = viper.GetString("mongo.extra_options")
}
