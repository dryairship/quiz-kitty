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
		outgoingMsg := models.OutgoingMessage(models.OutgoingTextMessage{
			Recipient: body.Entry[0].Messaging[0].Sender,
			Message:   body.Entry[0].Messaging[0].Message,
		})
		SendMessage(&outgoingMsg)
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
