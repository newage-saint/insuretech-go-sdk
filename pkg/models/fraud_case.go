package models

import (
	"time"
)

// FraudCase represents a fraud_case
type FraudCase struct {
	InvestigationNotes string       `json:"investigation_notes,omitempty"`
	Evidence           string       `json:"evidence,omitempty"`
	AuditInfo          interface{}  `json:"audit_info"`
	Id                 string       `json:"id"`
	CaseNumber         string       `json:"case_number"`
	Status             interface{}  `json:"status"`
	Outcome            *CaseOutcome `json:"outcome,omitempty"`
	InvestigatorId     string       `json:"investigator_id,omitempty"`
	ClosedAt           time.Time    `json:"closed_at,omitempty"`
	FraudAlertId       string       `json:"fraud_alert_id"`
	Priority           interface{}  `json:"priority"`
}
