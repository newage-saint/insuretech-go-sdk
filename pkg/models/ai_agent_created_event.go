package models

import (
	"time"
)

// AIAgentCreatedEvent represents a ai_agent_created_event
type AIAgentCreatedEvent struct {
	CorrelationId string    `json:"correlation_id,omitempty"`
	EventId       string    `json:"event_id,omitempty"`
	AgentId       string    `json:"agent_id,omitempty"`
	AgentName     string    `json:"agent_name,omitempty"`
	AgentType     string    `json:"agent_type,omitempty"`
	ModelName     string    `json:"model_name,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	CreatedBy     string    `json:"created_by,omitempty"`
}
