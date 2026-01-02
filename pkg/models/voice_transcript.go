package models

import (
	"time"
)

// VoiceTranscript represents a voice_transcript
type VoiceTranscript struct {
	Speaker        *SpeakerType `json:"speaker"`
	Text           string       `json:"text"`
	Confidence     float64      `json:"confidence,omitempty"`
	Timestamp      time.Time    `json:"timestamp"`
	AuditInfo      *AuditInfo   `json:"audit_info,omitempty"`
	Id             string       `json:"id"`
	VoiceSessionId string       `json:"voice_session_id"`
	Language       string       `json:"language"`
	SequenceNumber int          `json:"sequence_number"`
}
