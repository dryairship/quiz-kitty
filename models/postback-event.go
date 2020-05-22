package models

type PostbackPayload struct {
	Payload string `json:"payload"`
}

type PostbackEvent struct {
	Sender    User            `json:"sender"`
	Timestamp int64           `json:"timestamp"`
	Postback  PostbackPayload `json:"postback"`
}
