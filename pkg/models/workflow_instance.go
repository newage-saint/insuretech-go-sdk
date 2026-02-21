package models

import (
	"time"
)

// WorkflowInstance represents a workflow_instance
type WorkflowInstance struct {
	EntityType           string      `json:"entity_type"`
	CurrentStep          string      `json:"current_step,omitempty"`
	EntityId             string      `json:"entity_id"`
	Status               interface{} `json:"status"`
	Context              string      `json:"context,omitempty"`
	InitiatedBy          string      `json:"initiated_by"`
	StartedAt            time.Time   `json:"started_at"`
	CompletedAt          time.Time   `json:"completed_at,omitempty"`
	AuditInfo            interface{} `json:"audit_info"`
	Id                   string      `json:"id"`
	WorkflowDefinitionId string      `json:"workflow_definition_id"`
}
