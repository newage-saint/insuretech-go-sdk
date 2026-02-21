package models

import (
	"time"
)

// WorkflowInstance represents a workflow_instance
type WorkflowInstance struct {
	Context              string      `json:"context,omitempty"`
	CompletedAt          time.Time   `json:"completed_at,omitempty"`
	EntityType           string      `json:"entity_type"`
	EntityId             string      `json:"entity_id"`
	InitiatedBy          string      `json:"initiated_by"`
	StartedAt            time.Time   `json:"started_at"`
	AuditInfo            interface{} `json:"audit_info"`
	Id                   string      `json:"id"`
	WorkflowDefinitionId string      `json:"workflow_definition_id"`
	Status               interface{} `json:"status"`
	CurrentStep          string      `json:"current_step,omitempty"`
}
