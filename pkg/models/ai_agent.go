package models

import (
	"time"
)

// AIAgent represents a ai_agent
type AIAgent struct {
	AgentId      string                 `json:"agent_id"`
	Type         *AgentType             `json:"type"`
	LastActiveAt time.Time              `json:"last_active_at,omitempty"`
	AgentName    string                 `json:"agent_name"`
	ModelName    string                 `json:"model_name"`
	Status       interface{}            `json:"status"`
	Capabilities map[string]interface{} `json:"capabilities,omitempty"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
}
