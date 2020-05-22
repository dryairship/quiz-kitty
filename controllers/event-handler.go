package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/dryairship/messenger-quiz-bot/models"
)

func EventHandler(c *gin.Context) {
	var body models.MessageRequestBody
	err := c.ShouldBindBodyWith(&body, binding.JSON)
	if err == nil {
		go HandleTextMessage(&body.Entry[0].Messaging[0].Sender, &body.Entry[0].Messaging[0].Message.Text)
	}
	// Commented out because we are not using PostbackRequestBody as of now
	// else {
	// 	var body models.PostbackRequestBody
	// 	err := c.BindJSON(&body)
	// 	if err != nil {
	// 		c.AbortWithStatus(403)
	// 		return
	// 	}
	// }

	c.String(200, "EVENT_RECEIVED")
}
