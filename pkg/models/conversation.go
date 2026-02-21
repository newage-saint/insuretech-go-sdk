package models

import (
	"time"
)

// Conversation represents a conversation
type Conversation struct {
	UserId         string                 `json:"user_id"`
	AgentId        string                 `json:"agent_id"`
	Messages       []*Message             `json:"messages,omitempty"`
	Status         interface{}            `json:"status"`
	StartedAt      time.Time              `json:"started_at"`
	EndedAt        time.Time              `json:"ended_at,omitempty"`
	Context        map[string]interface{} `json:"context,omitempty"`
	ConversationId string                 `json:"conversation_id"`
}
