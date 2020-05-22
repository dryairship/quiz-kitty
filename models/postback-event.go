package models

type PostbackPayload struct {
	Payload string `json:"payload"`
}

type PostbackEvent struct {
	Postback PostbackPayload `json:"postback"`
}
