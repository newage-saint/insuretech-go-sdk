package models

import (
	"time"
)

// WorkflowConfig represents a workflow_config
type WorkflowConfig struct {
	Title       string              `json:"title,omitempty"`
	IsEnabled   bool                `json:"is_enabled,omitempty"`
	Rules       string              `json:"rules,omitempty"`
	CreatedAt   time.Time           `json:"created_at,omitempty"`
	ConfigId    string              `json:"config_id,omitempty"`
	ConfigType  *WorkflowConfigType `json:"config_type,omitempty"`
	Description string              `json:"description,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at,omitempty"`
	BusinessId  string              `json:"business_id,omitempty"`
}
