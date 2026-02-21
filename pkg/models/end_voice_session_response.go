package models

// EndVoiceSessionResponse represents a end_voice_session_response
type EndVoiceSessionResponse struct {
	Error           *Error `json:"error,omitempty"`
	Message         string `json:"message,omitempty"`
	DurationSeconds int    `json:"duration_seconds,omitempty"`
}
