package models

// EndVoiceSessionResponse represents a end_voice_session_response
type EndVoiceSessionResponse struct {
	Message         string `json:"message,omitempty"`
	DurationSeconds int    `json:"duration_seconds,omitempty"`
	Error           *Error `json:"error,omitempty"`
}
