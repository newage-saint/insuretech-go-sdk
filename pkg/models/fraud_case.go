package models

import (
	"time"
)

// FraudCase represents a fraud_case
type FraudCase struct {
	CaseNumber         string       `json:"case_number"`
	FraudAlertId       string       `json:"fraud_alert_id"`
	Priority           interface{}  `json:"priority"`
	InvestigationNotes string       `json:"investigation_notes,omitempty"`
	Evidence           string       `json:"evidence,omitempty"`
	ClosedAt           time.Time    `json:"closed_at,omitempty"`
	Id                 string       `json:"id"`
	Status             interface{}  `json:"status"`
	Outcome            *CaseOutcome `json:"outcome,omitempty"`
	InvestigatorId     string       `json:"investigator_id,omitempty"`
	AuditInfo          interface{}  `json:"audit_info"`
}
