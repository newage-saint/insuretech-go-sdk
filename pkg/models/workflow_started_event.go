package models

import (
	"time"
)

// WorkflowStartedEvent represents a workflow_started_event
type WorkflowStartedEvent struct {
	CorrelationId        string    `json:"correlation_id,omitempty"`
	Timestamp            time.Time `json:"timestamp,omitempty"`
	EventId              string    `json:"event_id,omitempty"`
	WorkflowInstanceId   string    `json:"workflow_instance_id,omitempty"`
	WorkflowDefinitionId string    `json:"workflow_definition_id,omitempty"`
	EntityType           string    `json:"entity_type,omitempty"`
	EntityId             string    `json:"entity_id,omitempty"`
	InitiatedBy          string    `json:"initiated_by,omitempty"`
}
