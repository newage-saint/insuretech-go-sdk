package models

import (
	"time"
)

// VoiceSession represents a voice_session
type VoiceSession struct {
	Id              string         `json:"id,omitempty"`
	UserId          string         `json:"user_id,omitempty"`
	Language        string         `json:"language,omitempty"`
	Status          *SessionStatus `json:"status,omitempty"`
	Intent          string         `json:"intent,omitempty"`
	Context         string         `json:"context,omitempty"`
	DurationSeconds int            `json:"duration_seconds,omitempty"`
	SessionId       string         `json:"session_id,omitempty"`
	PhoneNumber     string         `json:"phone_number,omitempty"`
	StartedAt       time.Time      `json:"started_at,omitempty"`
	EndedAt         time.Time      `json:"ended_at,omitempty"`
	AuditInfo       *AuditInfo     `json:"audit_info,omitempty"`
}
