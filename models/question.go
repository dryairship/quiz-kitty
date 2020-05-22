package models

import (
	"math/rand"
)

type Question struct {
	Category         string   `json:"category"`
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

func (question Question) ToTextMessage() (TextMessage, string) {
	var text string
	text += question.Question
	text += "\n\n"

	perm := rand.Perm(len(question.IncorrectAnswers) + 1)
	var currentChar byte = 'A'
	var correctAnswer string

	for i := range perm {
		if perm[i] == 0 {
			correctAnswer = string(currentChar)
			text += correctAnswer
			text += " ) "
			text += question.CorrectAnswer
			text += "\n"
		} else {
			text += string(currentChar)
			text += " ) "
			text += question.IncorrectAnswers[perm[i]-1]
			text += "\n"
		}
		currentChar++
	}

	return TextMessage{Text: text}, correctAnswer
}
