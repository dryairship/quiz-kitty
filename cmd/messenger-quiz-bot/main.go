package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/dryairship/messenger-quiz-bot/config"
	"github.com/dryairship/messenger-quiz-bot/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/webhook", controllers.VerificationHandler)
	r.POST("/webhook", controllers.EventHandler)

	log.Println("[INFO] Starting server on port ", config.PORT)
	log.Fatal(r.Run(":" + config.PORT))
}
