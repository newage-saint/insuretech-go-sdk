package models

import (
	"time"
)

// WorkflowTask represents a workflow_task
type WorkflowTask struct {
	DueDate            time.Time         `json:"due_date,omitempty"`
	CompletedAt        time.Time         `json:"completed_at,omitempty"`
	Id                 string            `json:"id"`
	Type               *WorkflowTaskType `json:"type"`
	AssignedTo         string            `json:"assigned_to,omitempty"`
	Status             interface{}       `json:"status"`
	AuditInfo          *AuditInfo        `json:"audit_info,omitempty"`
	WorkflowInstanceId string            `json:"workflow_instance_id"`
	StepName           string            `json:"step_name"`
	Decision           string            `json:"decision,omitempty"`
	Comments           string            `json:"comments,omitempty"`
}
