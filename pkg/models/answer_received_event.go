package models

import (
	"time"
)

// AnswerReceivedEvent represents a answer_received_event
type AnswerReceivedEvent struct {
	ReceivedAt time.Time  `json:"received_at,omitempty"`
	RoomId     string     `json:"room_id,omitempty"`
	Answer     *SDPAnswer `json:"answer,omitempty"`
}
