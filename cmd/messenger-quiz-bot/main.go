package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dryairship/messenger-quiz-bot/config"
	"github.com/dryairship/messenger-quiz-bot/controllers"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	r := gin.Default()

	r.GET("/webhook", controllers.VerificationHandler)
	r.POST("/webhook", controllers.EventHandler)

	log.Println("[INFO] Starting server on port ", config.PORT)
	log.Fatal(r.Run(":" + config.PORT))
}
