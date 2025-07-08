package models

import (
	"encoding/json"
)

// DTO For Input & Output
type WebhookRequestDTO struct {
	Id        int64
	Event     string
	Resource  string
	Timestamp int64
	Data      json.RawMessage
	Attempt   int
}

// DTO For Output
type WebhookResponseDTO struct {
	Body json.RawMessage
}
