package models

// VoiceSessionStartResponse represents a voice_session_start_response
type VoiceSessionStartResponse struct {
	Error          *Error `json:"error,omitempty"`
	VoiceSessionId string `json:"voice_session_id,omitempty"`
	SessionId      string `json:"session_id,omitempty"`
	Message        string `json:"message,omitempty"`
}
