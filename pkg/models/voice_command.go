package models

import (
	"time"
)

// VoiceCommand represents a voice_command
type VoiceCommand struct {
	Type             *CommandType `json:"type"`
	CommandText      string       `json:"command_text,omitempty"`
	Status           interface{}  `json:"status"`
	ResponseText     string       `json:"response_text,omitempty"`
	ConfidenceScore  float64      `json:"confidence_score,omitempty"`
	Id               string       `json:"id"`
	VoiceSessionId   string       `json:"voice_session_id"`
	AudioUrl         string       `json:"audio_url,omitempty"`
	Parameters       string       `json:"parameters,omitempty"`
	ResponseAudioUrl string       `json:"response_audio_url,omitempty"`
	ExecutedAt       time.Time    `json:"executed_at"`
	AuditInfo        *AuditInfo   `json:"audit_info,omitempty"`
}
