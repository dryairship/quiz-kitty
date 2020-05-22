package config

import (
	"log"

	"github.com/spf13/viper"
)

var PORT string
var CORRECT_TOKEN string
var ACCESS_TOKEN string

func init() {
	viper.SetConfigName("quiz-bot-config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("[WARNING] Unable to locate configuration file")
	}

	viper.SetEnvPrefix("quizbot")
	viper.AutomaticEnv()

	PORT = viper.GetString("port")
	CORRECT_TOKEN = viper.GetString("correct_token")
	ACCESS_TOKEN = viper.GetString("access_token")
}
