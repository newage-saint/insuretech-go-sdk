package models

import (
	"time"
)

// VoiceTranscript represents a voice_transcript
type VoiceTranscript struct {
	Id             string       `json:"id"`
	Text           string       `json:"text"`
	Language       string       `json:"language"`
	Confidence     float64      `json:"confidence,omitempty"`
	Timestamp      time.Time    `json:"timestamp"`
	AuditInfo      interface{}  `json:"audit_info"`
	VoiceSessionId string       `json:"voice_session_id"`
	Speaker        *SpeakerType `json:"speaker"`
	SequenceNumber int          `json:"sequence_number"`
}
