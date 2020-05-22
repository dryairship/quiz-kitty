package models

type MessageRequestBody struct {
	Object string         `json:"object"`
	Entry  []MessageEntry `json:"entry"`
}

type PostbackRequestBody struct {
	Object string          `json:"object"`
	Entry  []PostbackEntry `json:"entry"`
}

type QuestionAPIResponseBody struct {
	ResponseCode int        `json:"response_code"`
	Results      []Question `json:"results"`
}
