package models

import (
	"time"
)

// Task represents a task
type Task struct {
	CompletedAt       time.Time     `json:"completed_at,omitempty"`
	Id                string        `json:"id"`
	Description       string        `json:"description,omitempty"`
	Priority          interface{}   `json:"priority"`
	AssignedTo        string        `json:"assigned_to,omitempty"`
	RelatedEntityId   string        `json:"related_entity_id,omitempty"`
	AuditInfo         interface{}   `json:"audit_info"`
	Title             string        `json:"title"`
	Type              *TaskTaskType `json:"type"`
	Status            interface{}   `json:"status"`
	CreatedBy         string        `json:"created_by,omitempty"`
	RelatedEntityType string        `json:"related_entity_type,omitempty"`
	DueDate           time.Time     `json:"due_date,omitempty"`
}
