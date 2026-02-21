package models

// ChatRequest represents a chat_request
type ChatRequest struct {
	ConversationId string                 `json:"conversation_id"`
	UserId         string                 `json:"user_id"`
	AgentId        string                 `json:"agent_id"`
	Message        string                 `json:"message,omitempty"`
	Context        map[string]interface{} `json:"context,omitempty"`
}
