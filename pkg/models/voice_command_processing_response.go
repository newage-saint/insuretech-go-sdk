package models

// VoiceCommandProcessingResponse represents a voice_command_processing_response
type VoiceCommandProcessingResponse struct {
	ResponseAudioUrl string  `json:"response_audio_url,omitempty"`
	ConfidenceScore  float64 `json:"confidence_score,omitempty"`
	Error            *Error  `json:"error,omitempty"`
	VoiceCommandId   string  `json:"voice_command_id,omitempty"`
	CommandType      string  `json:"command_type,omitempty"`
	ResponseText     string  `json:"response_text,omitempty"`
}
