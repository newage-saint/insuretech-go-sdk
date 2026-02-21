package models

import (
	"time"
)

// Task represents a task
type Task struct {
	RelatedEntityId   string        `json:"related_entity_id,omitempty"`
	DueDate           time.Time     `json:"due_date,omitempty"`
	CompletedAt       time.Time     `json:"completed_at,omitempty"`
	AuditInfo         interface{}   `json:"audit_info"`
	Description       string        `json:"description,omitempty"`
	Type              *TaskTaskType `json:"type"`
	Priority          interface{}   `json:"priority"`
	Status            interface{}   `json:"status"`
	AssignedTo        string        `json:"assigned_to,omitempty"`
	RelatedEntityType string        `json:"related_entity_type,omitempty"`
	Id                string        `json:"id"`
	Title             string        `json:"title"`
	CreatedBy         string        `json:"created_by,omitempty"`
}
