package models

import (
	"time"
)

// Task represents a task
type Task struct {
	Priority          interface{}   `json:"priority"`
	Status            interface{}   `json:"status"`
	AssignedTo        string        `json:"assigned_to,omitempty"`
	CreatedBy         string        `json:"created_by,omitempty"`
	RelatedEntityId   string        `json:"related_entity_id,omitempty"`
	Id                string        `json:"id"`
	Description       string        `json:"description,omitempty"`
	RelatedEntityType string        `json:"related_entity_type,omitempty"`
	DueDate           time.Time     `json:"due_date,omitempty"`
	CompletedAt       time.Time     `json:"completed_at,omitempty"`
	AuditInfo         *AuditInfo    `json:"audit_info,omitempty"`
	Title             string        `json:"title"`
	Type              *TaskTaskType `json:"type"`
}
