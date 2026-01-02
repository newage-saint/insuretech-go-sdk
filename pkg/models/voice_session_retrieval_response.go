package models

// VoiceSessionRetrievalResponse represents a voice_session_retrieval_response
type VoiceSessionRetrievalResponse struct {
	VoiceSession *VoiceSession   `json:"voice_session,omitempty"`
	Commands     []*VoiceCommand `json:"commands,omitempty"`
	Error        *Error          `json:"error,omitempty"`
}
