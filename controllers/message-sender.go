package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dryairship/quiz-kitty/config"
	"github.com/dryairship/quiz-kitty/models"
)

var QUERY_URL string = "https://graph.facebook.com/v2.6/me/messages?access_token=" + config.ACCESS_TOKEN

func SendMessage(msg *models.OutgoingMessage) {
	requestBody, err := json.Marshal(msg)
	if err != nil {
		log.Printf("[ERROR] Can't JSON Marshal OutgoingMessage: %v\n\tError: %v", *msg, err)
		return
	}

	response, err := http.Post(QUERY_URL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Printf("[ERROR] Can't send POST request: %v\n", err)
		return
	}

	if response.StatusCode != http.StatusOK {
		defer response.Body.Close()
		responseBody, _ := ioutil.ReadAll(response.Body)
		log.Printf("[ERROR] Graph API Error: %s\n", string(responseBody))
	}
}

func SendTextMessageToUser(user *models.User, textMessage *models.TextMessage) {
	msg := models.OutgoingMessage(models.OutgoingTextMessage{
		Recipient: *user,
		Message:   *textMessage,
	})
	SendMessage(&msg)
}
