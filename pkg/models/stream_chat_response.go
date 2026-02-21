package models

// StreamChatResponse represents a stream_chat_response
type StreamChatResponse struct {
	Error   *Error   `json:"error,omitempty"`
	Message *Message `json:"message,omitempty"`
}
