package models

import (
	"time"
)

// AIAgent represents a ai_agent
type AIAgent struct {
	Capabilities map[string]interface{} `json:"capabilities,omitempty"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
	AgentId      string                 `json:"agent_id"`
	ModelName    string                 `json:"model_name"`
	LastActiveAt time.Time              `json:"last_active_at,omitempty"`
	AgentName    string                 `json:"agent_name"`
	Type         *AgentType             `json:"type"`
	Status       interface{}            `json:"status"`
}
