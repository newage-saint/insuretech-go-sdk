package models

import (
	"time"
)

// AIAgent represents a ai_agent
type AIAgent struct {
	AgentId      string                 `json:"agent_id"`
	AgentName    string                 `json:"agent_name"`
	Status       interface{}            `json:"status"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
	Type         *AgentType             `json:"type"`
	ModelName    string                 `json:"model_name"`
	Capabilities map[string]interface{} `json:"capabilities,omitempty"`
	LastActiveAt time.Time              `json:"last_active_at,omitempty"`
}
