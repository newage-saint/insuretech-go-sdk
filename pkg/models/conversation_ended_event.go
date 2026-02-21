package models

import (
	"time"
)

// ConversationEndedEvent represents a conversation_ended_event
type ConversationEndedEvent struct {
	UserId           string    `json:"user_id,omitempty"`
	DurationSeconds  int       `json:"duration_seconds,omitempty"`
	ConversationId   string    `json:"conversation_id,omitempty"`
	AgentId          string    `json:"agent_id,omitempty"`
	MessageCount     int       `json:"message_count,omitempty"`
	ResolutionStatus string    `json:"resolution_status,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
	CorrelationId    string    `json:"correlation_id,omitempty"`
	EventId          string    `json:"event_id,omitempty"`
}
