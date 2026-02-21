package models

import (
	"time"
)

// VoiceTranscript represents a voice_transcript
type VoiceTranscript struct {
	Id             string       `json:"id"`
	Speaker        *SpeakerType `json:"speaker"`
	Confidence     float64      `json:"confidence,omitempty"`
	SequenceNumber int          `json:"sequence_number"`
	Timestamp      time.Time    `json:"timestamp"`
	VoiceSessionId string       `json:"voice_session_id"`
	Text           string       `json:"text"`
	Language       string       `json:"language"`
	AuditInfo      interface{}  `json:"audit_info"`
}
