package types

type MessageEnvelope struct {
	ToUserID string
	Data     interface{}
}

type MessagePayload struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}
