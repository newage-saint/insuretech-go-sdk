package models

import (
	"time"
)

// ComplianceLog represents a compliance_log
type ComplianceLog struct {
	EntityType  string            `json:"entity_type"`
	EntityId    string            `json:"entity_id"`
	Description string            `json:"description"`
	Evidence    string            `json:"evidence,omitempty"`
	Timestamp   time.Time         `json:"timestamp"`
	Id          string            `json:"id"`
	Status      *ComplianceStatus `json:"status"`
	PerformedBy string            `json:"performed_by,omitempty"`
	Type        *ComplianceType   `json:"type"`
	Regulation  string            `json:"regulation"`
}
