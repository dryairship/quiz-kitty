package models

type TextMessage struct {
	Text string `json:"text"`
}

type MessageEvent struct {
	Sender    User        `json:"sender"`
	Timestamp int64       `json:"timestamp"`
	Message   TextMessage `json:"message"`
}
