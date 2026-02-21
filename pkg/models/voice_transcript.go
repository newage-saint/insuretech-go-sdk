package models

import (
	"time"
)

// VoiceTranscript represents a voice_transcript
type VoiceTranscript struct {
	VoiceSessionId string       `json:"voice_session_id"`
	SequenceNumber int          `json:"sequence_number"`
	Timestamp      time.Time    `json:"timestamp"`
	AuditInfo      interface{}  `json:"audit_info"`
	Id             string       `json:"id"`
	Speaker        *SpeakerType `json:"speaker"`
	Text           string       `json:"text"`
	Language       string       `json:"language"`
	Confidence     float64      `json:"confidence,omitempty"`
}
