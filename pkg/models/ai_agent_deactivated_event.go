package models

import (
	"time"
)

// AIAgentDeactivatedEvent represents a ai_agent_deactivated_event
type AIAgentDeactivatedEvent struct {
	EventId       string    `json:"event_id,omitempty"`
	AgentId       string    `json:"agent_id,omitempty"`
	AgentName     string    `json:"agent_name,omitempty"`
	Reason        string    `json:"reason,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
}
