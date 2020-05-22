package models

type MessageEntry struct {
	Messaging []MessageEvent `json:"messaging"`
}

type PostbackEntry struct {
	Messaging []PostbackEvent `json:"messaging"`
}
