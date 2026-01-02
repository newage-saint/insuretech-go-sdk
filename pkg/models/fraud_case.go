package models

import (
	"time"
)

// FraudCase represents a fraud_case
type FraudCase struct {
	Status             interface{}  `json:"status"`
	ClosedAt           time.Time    `json:"closed_at,omitempty"`
	AuditInfo          *AuditInfo   `json:"audit_info,omitempty"`
	Id                 string       `json:"id"`
	CaseNumber         string       `json:"case_number"`
	InvestigationNotes string       `json:"investigation_notes,omitempty"`
	Evidence           string       `json:"evidence,omitempty"`
	Outcome            *CaseOutcome `json:"outcome,omitempty"`
	InvestigatorId     string       `json:"investigator_id,omitempty"`
	FraudAlertId       string       `json:"fraud_alert_id"`
	Priority           interface{}  `json:"priority"`
}
