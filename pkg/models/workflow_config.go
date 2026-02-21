package models

import (
	"time"
)

// WorkflowConfig represents a workflow_config
type WorkflowConfig struct {
	BusinessId  string              `json:"business_id,omitempty"`
	Rules       string              `json:"rules,omitempty"`
	ConfigType  *WorkflowConfigType `json:"config_type,omitempty"`
	Title       string              `json:"title,omitempty"`
	Description string              `json:"description,omitempty"`
	IsEnabled   bool                `json:"is_enabled,omitempty"`
	CreatedAt   time.Time           `json:"created_at,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at,omitempty"`
	ConfigId    string              `json:"config_id,omitempty"`
}
