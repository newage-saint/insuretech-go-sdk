package models

import (
	"time"
)

// VoiceSession represents a voice_session
type VoiceSession struct {
	Status          *SessionStatus `json:"status,omitempty"`
	Context         string         `json:"context,omitempty"`
	DurationSeconds int            `json:"duration_seconds,omitempty"`
	AuditInfo       interface{}    `json:"audit_info"`
	PhoneNumber     string         `json:"phone_number,omitempty"`
	Language        string         `json:"language,omitempty"`
	Intent          string         `json:"intent,omitempty"`
	StartedAt       time.Time      `json:"started_at,omitempty"`
	EndedAt         time.Time      `json:"ended_at,omitempty"`
	Id              string         `json:"id,omitempty"`
	SessionId       string         `json:"session_id,omitempty"`
	UserId          string         `json:"user_id,omitempty"`
}
