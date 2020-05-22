package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/dryairship/messenger-quiz-bot/config"
)

func VerificationHandler(c *gin.Context) {
	query := c.Request.URL.Query()

	mode := query["hub.mode"]
	token := query["hub.verify_token"]
	challenge := query["hub.challenge"]

	if len(mode) != 1 || len(token) != 1 || len(challenge) != 1 || mode[0] != "subscribe" || token[0] != config.CORRECT_TOKEN {
		c.AbortWithStatus(403)
	} else {
		c.String(200, challenge[0])
	}
}
