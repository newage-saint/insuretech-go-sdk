package models

import (
	"time"
)

// WorkflowTask represents a workflow_task
type WorkflowTask struct {
	CompletedAt        time.Time         `json:"completed_at,omitempty"`
	AuditInfo          interface{}       `json:"audit_info"`
	WorkflowInstanceId string            `json:"workflow_instance_id"`
	StepName           string            `json:"step_name"`
	Type               *WorkflowTaskType `json:"type"`
	Status             interface{}       `json:"status"`
	Decision           string            `json:"decision,omitempty"`
	DueDate            time.Time         `json:"due_date,omitempty"`
	Id                 string            `json:"id"`
	AssignedTo         string            `json:"assigned_to,omitempty"`
	Comments           string            `json:"comments,omitempty"`
}
