package models

import (
	"time"
)

// WorkflowTaskAssignedEvent represents a workflow_task_assigned_event
type WorkflowTaskAssignedEvent struct {
	CorrelationId      string    `json:"correlation_id,omitempty"`
	Timestamp          time.Time `json:"timestamp,omitempty"`
	EventId            string    `json:"event_id,omitempty"`
	TaskId             string    `json:"task_id,omitempty"`
	WorkflowInstanceId string    `json:"workflow_instance_id,omitempty"`
	StepName           string    `json:"step_name,omitempty"`
	AssignedTo         string    `json:"assigned_to,omitempty"`
}
