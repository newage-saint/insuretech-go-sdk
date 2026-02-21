package models

import (
	"time"
)

// Task represents a task
type Task struct {
	CreatedBy         string        `json:"created_by,omitempty"`
	RelatedEntityType string        `json:"related_entity_type,omitempty"`
	DueDate           time.Time     `json:"due_date,omitempty"`
	CompletedAt       time.Time     `json:"completed_at,omitempty"`
	AuditInfo         interface{}   `json:"audit_info"`
	Id                string        `json:"id"`
	Title             string        `json:"title"`
	Description       string        `json:"description,omitempty"`
	Type              *TaskTaskType `json:"type"`
	Priority          interface{}   `json:"priority"`
	AssignedTo        string        `json:"assigned_to,omitempty"`
	RelatedEntityId   string        `json:"related_entity_id,omitempty"`
	Status            interface{}   `json:"status"`
}
