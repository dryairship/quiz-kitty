package controllers

import (
	"strings"

	"github.com/dryairship/messenger-quiz-bot/db"
	"github.com/dryairship/messenger-quiz-bot/models"
)

func askIfUserWantsQuestion(user *models.User) {

}

func handleUserWantsNewQuestion(user *models.User) {

}

func handleUserDoesNotWantNewQuestion(user *models.User) {

}

func handleInvalidUserAnswer(user *models.User, maxChar byte) {

}

func handleValidUserAnswer(user *models.User, correctAnswerText *string, isAnswerCorrect bool) {

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
