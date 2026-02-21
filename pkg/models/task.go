package models

import (
	"time"
)

// Task represents a task
type Task struct {
	CompletedAt       time.Time     `json:"completed_at,omitempty"`
	Type              *TaskTaskType `json:"type"`
	Priority          interface{}   `json:"priority"`
	AssignedTo        string        `json:"assigned_to,omitempty"`
	RelatedEntityType string        `json:"related_entity_type,omitempty"`
	RelatedEntityId   string        `json:"related_entity_id,omitempty"`
	AuditInfo         interface{}   `json:"audit_info"`
	Id                string        `json:"id"`
	Title             string        `json:"title"`
	Description       string        `json:"description,omitempty"`
	Status            interface{}   `json:"status"`
	CreatedBy         string        `json:"created_by,omitempty"`
	DueDate           time.Time     `json:"due_date,omitempty"`
}
