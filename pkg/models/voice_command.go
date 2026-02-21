package models

import (
	"time"
)

// VoiceCommand represents a voice_command
type VoiceCommand struct {
	ResponseAudioUrl string       `json:"response_audio_url,omitempty"`
	Type             *CommandType `json:"type"`
	ConfidenceScore  float64      `json:"confidence_score,omitempty"`
	ExecutedAt       time.Time    `json:"executed_at"`
	AuditInfo        interface{}  `json:"audit_info"`
	Id               string       `json:"id"`
	VoiceSessionId   string       `json:"voice_session_id"`
	CommandText      string       `json:"command_text,omitempty"`
	AudioUrl         string       `json:"audio_url,omitempty"`
	Parameters       string       `json:"parameters,omitempty"`
	Status           interface{}  `json:"status"`
	ResponseText     string       `json:"response_text,omitempty"`
}
