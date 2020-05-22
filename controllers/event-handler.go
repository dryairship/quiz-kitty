package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/dryairship/messenger-quiz-bot/models"
)

func EventHandler(c *gin.Context) {
	var body models.MessageRequestBody
	err := c.ShouldBindBodyWith(&body, binding.JSON)
	if err == nil {
		fmt.Println(body)
	} else {
		var body models.PostbackRequestBody
		err := c.BindJSON(&body)
		if err != nil {
			c.AbortWithStatus(403)
			return
		}
		fmt.Println(body)
	}

	c.String(200, "EVENT_RECEIVED")
}
