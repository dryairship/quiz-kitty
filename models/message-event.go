package models

type MessagePayloadElementButton struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Payload string `json:"payload"`
}

type MessagePayloadElement struct {
	Title    string                        `json:"title"`
	Subtitle string                        `json:"subtitle"`
	Buttons  []MessagePayloadElementButton `json:"buttons"`
}

type MessagePayload struct {
	TemplateType string                  `json:"template_type"`
	Elements     []MessagePayloadElement `json:"elements"`
}

type MessageAttachment struct {
	Type    string         `json:"type"`
	Payload MessagePayload `json:"payload"`
}

type TextMessage struct {
	Text string `json:"text"`
}

type GenericMessage struct {
	Attachment MessageAttachment `json:"attachment"`
}

type MessageEvent struct {
	Sender    User        `json:"sender"`
	Timestamp int64       `json:"timestamp"`
	Message   TextMessage `json:"message"`
}
