package models

import (
	"time"
)

// ComplianceLog represents a compliance_log
type ComplianceLog struct {
	EntityType  string            `json:"entity_type"`
	EntityId    string            `json:"entity_id"`
	Status      *ComplianceStatus `json:"status"`
	Evidence    string            `json:"evidence,omitempty"`
	Timestamp   time.Time         `json:"timestamp"`
	Id          string            `json:"id"`
	Type        *ComplianceType   `json:"type"`
	Regulation  string            `json:"regulation"`
	Description string            `json:"description"`
	PerformedBy string            `json:"performed_by,omitempty"`
}
