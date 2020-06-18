package controllers

import (
	"fmt"
	"strings"

	"github.com/dryairship/quiz-kitty/db"
	"github.com/dryairship/quiz-kitty/models"
)

func handleQuestionAPIProblem(user *models.User) {
	SendTextMessageToUser(user, &models.TextMessage{
		Text: MESSAGE_API_PROBLEM,
	})
	db.SetRedisUserData(user.Id, &models.RedisUserData{
		State: models.USER_STATE_IDLE,
	})
}

func askIfUserWantsQuestion(user *models.User) {
	SendTextMessageToUser(user, &models.TextMessage{
		Text: MESSAGE_WANT_QUESTION,
	})
	db.SetRedisUserData(user.Id, &models.RedisUserData{
		State:             models.USER_STATE_WANT_QUESTION,
		MaxAcceptableChar: 'B',
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
		Text: MESSAGE_BYE,
	})
	db.SetRedisUserData(user.Id, &models.RedisUserData{
		State: models.USER_STATE_IDLE,
	})
}

func handleInvalidUserAnswer(user *models.User, maxChar byte) {
	SendTextMessageToUser(user, &models.TextMessage{
		Text: fmt.Sprintf(MESSAGE_INVALID_ANSWER, maxChar),
	})
}

func handleValidUserAnswer(user *models.User, correctAnswerText *string, isAnswerCorrect bool) {
	if isAnswerCorrect {
		SendTextMessageToUser(user, &models.TextMessage{
			Text: fmt.Sprintf(MESSAGE_CORRECT_ANSWER, *correctAnswerText),
		})
		db.UpdateUserScore(user.Id, 2)
	} else {
		SendTextMessageToUser(user, &models.TextMessage{
			Text: fmt.Sprintf(MESSAGE_INCORRECT_ANSWER, *correctAnswerText),
		})
		db.UpdateUserScore(user.Id, -1)
	}
	db.SetRedisUserData(user.Id, &models.RedisUserData{
		State:             models.USER_STATE_WANT_QUESTION,
		MaxAcceptableChar: 'B',
	})
}

func handleUserScore(user *models.User) {
	score := db.GetUserScore(user.Id)

	if score > 0 {
		SendTextMessageToUser(user, &models.TextMessage{
			Text: fmt.Sprintf(MESSAGE_POSITIVE_SCORE, score),
		})
	} else {
		SendTextMessageToUser(user, &models.TextMessage{
			Text: fmt.Sprintf(MESSAGE_NEGATIVE_SCORE, score),
		})
	}
}

func HandleTextMessage(user *models.User, message *string) {
	userData, err := db.GetRedisUserData(user.Id)
	if err != nil {
		askIfUserWantsQuestion(user)
		return
	}

	if strings.Contains(strings.ToLower(*message), "score") {
		handleUserScore(user)
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
