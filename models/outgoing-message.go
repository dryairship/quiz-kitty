package models

type OutgoingMessage interface {
}

type OutgoingTextMessage struct {
	Recipient User        `json:"recipient"`
	Message   TextMessage `json:"message"`
}
