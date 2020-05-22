package controllers

import (
	"strings"

	"github.com/dryairship/messenger-quiz-bot/db"
	"github.com/dryairship/messenger-quiz-bot/models"
)

func handleQuestionAPIProblem(user *models.User) {
	SendTextMessageToUser(user, &models.TextMessage{
		Text: "I'm sorry I cannot handle your request right now. There is some problem with our questions API.",
	})
	db.SetRedisUserData(user.Id, &models.RedisUserData{
		State: models.USER_STATE_IDLE,
	})
}

func askIfUserWantsQuestion(user *models.User) {
	SendTextMessageToUser(user, &models.TextMessage{
		Text: "Hey! Do you want me to ask you a question?\n\nA ) Yes\nB ) No",
	})
	db.SetRedisUserData(user.Id, &models.RedisUserData{
		State: models.USER_STATE_WANT_QUESTION,
	})
}

func handleUserWantsNewQuestion(user *models.User) {
	question, err := GetQuestion()
	if err != nil {
		handleQuestionAPIProblem(user)
		return
	}

	textMessage, correctChar, maxChar, correctText := question.ToTextMessage()
	SendTextMessageToUser(user, &textMessage)
	db.SetRedisUserData(user.Id, &models.RedisUserData{
		State:             models.USER_STATE_ACTIVE_QUESTION,
		CorrectAnswerChar: correctChar,
		CorrectAnswerText: correctText,
		MaxAcceptableChar: maxChar,
	})
}

func handleUserDoesNotWantNewQuestion(user *models.User) {
	SendTextMessageToUser(user, &models.TextMessage{
		Text: "Bye! Come back again soon!",
	})
	db.SetRedisUserData(user.Id, &models.RedisUserData{
		State: models.USER_STATE_IDLE,
	})
}

func handleInvalidUserAnswer(user *models.User, maxChar byte) {
	SendTextMessageToUser(user, &models.TextMessage{
		Text: "That is not a valid response. Please choose a letter from A-" + string(maxChar) + ".",
	})
}

func handleValidUserAnswer(user *models.User, correctAnswerText *string, isAnswerCorrect bool) {
	if isAnswerCorrect {
		SendTextMessageToUser(user, &models.TextMessage{
			Text: "That's right! :D\n" + *correctAnswerText + " is the correct answer!\n\nDo you want me to ask you another question?\n\nA ) Yes\nB ) No",
		})
	} else {
		SendTextMessageToUser(user, &models.TextMessage{
			Text: "That's incorrect! :(\n" + *correctAnswerText + " is the correct answer!\n\nDo you want me to ask you another question?\n\nA ) Yes\nB ) No",
		})
	}
	db.SetRedisUserData(user.Id, &models.RedisUserData{
		State: models.USER_STATE_WANT_QUESTION,
	})
}

func HandleTextMessage(user *models.User, message *string) {
	userData, err := db.GetRedisUserData(user.Id)
	if err != nil {
		askIfUserWantsQuestion(user)
		return
	}

	switch userData.State {
	case models.USER_STATE_ACTIVE_QUESTION:
		if len(*message) != 1 {
			handleInvalidUserAnswer(user, userData.MaxAcceptableChar)
			return
		}

		userChar := strings.ToUpper(*message)[0]
		if userChar > userData.MaxAcceptableChar {
			handleInvalidUserAnswer(user, userData.MaxAcceptableChar)
			return
		}

		handleValidUserAnswer(user, &userData.CorrectAnswerText, userData.CorrectAnswerChar == userChar)

	case models.USER_STATE_WANT_QUESTION:
		userAnswer := strings.ToUpper(*message)

		if userAnswer == "YES" {
			handleUserWantsNewQuestion(user)
			return
		} else if userAnswer == "NO" {
			handleUserDoesNotWantNewQuestion(user)
			return
		}

		if len(userAnswer) != 1 {
			handleInvalidUserAnswer(user, userData.MaxAcceptableChar)
			return
		}

		if userAnswer[0] == 'A' || userAnswer[0] == 'Y' {
			handleUserWantsNewQuestion(user)
			return
		} else if userAnswer[0] == 'B' || userAnswer[0] == 'N' {
			handleUserDoesNotWantNewQuestion(user)
			return
		}

		handleInvalidUserAnswer(user, userData.MaxAcceptableChar)

	case models.USER_STATE_IDLE:
		askIfUserWantsQuestion(user)
	}
}
