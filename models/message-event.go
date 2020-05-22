package models

type MessageText struct {
	Text string `json:"text"`
}

type MessageEvent struct {
	Message MessageText `json:"message"`
}
