package models

import (
	"time"
)

// VoiceSessionStartedEvent represents a voice_session_started_event
type VoiceSessionStartedEvent struct {
	Language       string    `json:"language,omitempty"`
	CorrelationId  string    `json:"correlation_id,omitempty"`
	Timestamp      time.Time `json:"timestamp,omitempty"`
	EventId        string    `json:"event_id,omitempty"`
	VoiceSessionId string    `json:"voice_session_id,omitempty"`
	SessionId      string    `json:"session_id,omitempty"`
	UserId         string    `json:"user_id,omitempty"`
}
