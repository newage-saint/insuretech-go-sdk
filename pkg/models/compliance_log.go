package models

import (
	"time"
)

// ComplianceLog represents a compliance_log
type ComplianceLog struct {
	Id          string            `json:"id"`
	Type        *ComplianceType   `json:"type"`
	EntityType  string            `json:"entity_type"`
	EntityId    string            `json:"entity_id"`
	Status      *ComplianceStatus `json:"status"`
	Regulation  string            `json:"regulation"`
	Description string            `json:"description"`
	Evidence    string            `json:"evidence,omitempty"`
	PerformedBy string            `json:"performed_by,omitempty"`
	Timestamp   time.Time         `json:"timestamp"`
}
