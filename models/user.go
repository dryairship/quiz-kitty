package models

type UserState int

const (
	USER_STATE_ACTIVE_QUESTION UserState = iota
	USER_STATE_WANT_QUESTION
	USER_STATE_IDLE
)

type User struct {
	Id string `json:"id"`
}

type RedisUserData struct {
	State             UserState
	CorrectAnswerText string
	CorrectAnswerChar byte
	MaxAcceptableChar byte
}
