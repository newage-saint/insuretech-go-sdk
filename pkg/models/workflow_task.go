package models

import (
	"time"
)

// WorkflowTask represents a workflow_task
type WorkflowTask struct {
	WorkflowInstanceId string            `json:"workflow_instance_id"`
	AssignedTo         string            `json:"assigned_to,omitempty"`
	DueDate            time.Time         `json:"due_date,omitempty"`
	CompletedAt        time.Time         `json:"completed_at,omitempty"`
	AuditInfo          interface{}       `json:"audit_info"`
	Id                 string            `json:"id"`
	StepName           string            `json:"step_name"`
	Type               *WorkflowTaskType `json:"type"`
	Status             interface{}       `json:"status"`
	Decision           string            `json:"decision,omitempty"`
	Comments           string            `json:"comments,omitempty"`
}
