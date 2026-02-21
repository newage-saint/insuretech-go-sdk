package models

import (
	"time"
)

// FraudCase represents a fraud_case
type FraudCase struct {
	Id                 string       `json:"id"`
	FraudAlertId       string       `json:"fraud_alert_id"`
	Priority           interface{}  `json:"priority"`
	InvestigationNotes string       `json:"investigation_notes,omitempty"`
	Evidence           string       `json:"evidence,omitempty"`
	InvestigatorId     string       `json:"investigator_id,omitempty"`
	ClosedAt           time.Time    `json:"closed_at,omitempty"`
	AuditInfo          interface{}  `json:"audit_info"`
	CaseNumber         string       `json:"case_number"`
	Status             interface{}  `json:"status"`
	Outcome            *CaseOutcome `json:"outcome,omitempty"`
}
