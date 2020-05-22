package models

import (
	"html"
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

func (question Question) ToTextMessage() (TextMessage, byte, byte, string) {
	var text string
	text += html.UnescapeString(question.Question)
	text += "\n\n"

	perm := rand.Perm(len(question.IncorrectAnswers) + 1)
	var currentChar byte = 'A'
	var correctAnswer byte

	for i := range perm {
		text += string(currentChar)
		text += " ) "
		if perm[i] == 0 {
			correctAnswer = currentChar
			text += html.UnescapeString(question.CorrectAnswer)
		} else {
			text += html.UnescapeString(question.IncorrectAnswers[perm[i]-1])
		}
		text += "\n"
		currentChar++
	}

	return TextMessage{Text: text}, correctAnswer, currentChar - 1, question.CorrectAnswer
}
