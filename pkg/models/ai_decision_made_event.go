package models

import (
	"time"
)

// AIDecisionMadeEvent represents a ai_decision_made_event
type AIDecisionMadeEvent struct {
	EventId         string    `json:"event_id,omitempty"`
	AgentId         string    `json:"agent_id,omitempty"`
	AgentName       string    `json:"agent_name,omitempty"`
	DecisionType    string    `json:"decision_type,omitempty"`
	ConfidenceScore float64   `json:"confidence_score,omitempty"`
	Timestamp       time.Time `json:"timestamp,omitempty"`
	EntityType      string    `json:"entity_type,omitempty"`
	EntityId        string    `json:"entity_id,omitempty"`
	Decision        string    `json:"decision,omitempty"`
	Reasoning       []string  `json:"reasoning,omitempty"`
	CorrelationId   string    `json:"correlation_id,omitempty"`
}
