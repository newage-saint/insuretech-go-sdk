package models

import (
	"time"
)

// ComplianceLog represents a compliance_log
type ComplianceLog struct {
	Evidence    string            `json:"evidence,omitempty"`
	PerformedBy string            `json:"performed_by,omitempty"`
	Timestamp   time.Time         `json:"timestamp"`
	Id          string            `json:"id"`
	Type        *ComplianceType   `json:"type"`
	Regulation  string            `json:"regulation"`
	EntityType  string            `json:"entity_type"`
	Description string            `json:"description"`
	EntityId    string            `json:"entity_id"`
	Status      *ComplianceStatus `json:"status"`
}
