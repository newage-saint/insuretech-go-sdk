package models

import (
	"time"
)

// ConversationEndedEvent represents a conversation_ended_event
type ConversationEndedEvent struct {
	ResolutionStatus string    `json:"resolution_status,omitempty"`
	CorrelationId    string    `json:"correlation_id,omitempty"`
	ConversationId   string    `json:"conversation_id,omitempty"`
	UserId           string    `json:"user_id,omitempty"`
	MessageCount     int       `json:"message_count,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
	EventId          string    `json:"event_id,omitempty"`
	AgentId          string    `json:"agent_id,omitempty"`
	DurationSeconds  int       `json:"duration_seconds,omitempty"`
}
