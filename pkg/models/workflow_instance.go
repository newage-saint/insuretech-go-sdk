package models

import (
	"time"
)

// WorkflowInstance represents a workflow_instance
type WorkflowInstance struct {
	EntityId             string      `json:"entity_id"`
	Status               interface{} `json:"status"`
	CurrentStep          string      `json:"current_step,omitempty"`
	InitiatedBy          string      `json:"initiated_by"`
	CompletedAt          time.Time   `json:"completed_at,omitempty"`
	AuditInfo            *AuditInfo  `json:"audit_info,omitempty"`
	EntityType           string      `json:"entity_type"`
	Context              string      `json:"context,omitempty"`
	StartedAt            time.Time   `json:"started_at"`
	Id                   string      `json:"id"`
	WorkflowDefinitionId string      `json:"workflow_definition_id"`
}
