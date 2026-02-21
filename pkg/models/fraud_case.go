package models

import (
	"time"
)

// FraudCase represents a fraud_case
type FraudCase struct {
	InvestigationNotes string       `json:"investigation_notes,omitempty"`
	Status             interface{}  `json:"status"`
	Id                 string       `json:"id"`
	FraudAlertId       string       `json:"fraud_alert_id"`
	Evidence           string       `json:"evidence,omitempty"`
	Outcome            *CaseOutcome `json:"outcome,omitempty"`
	InvestigatorId     string       `json:"investigator_id,omitempty"`
	ClosedAt           time.Time    `json:"closed_at,omitempty"`
	AuditInfo          interface{}  `json:"audit_info"`
	CaseNumber         string       `json:"case_number"`
	Priority           interface{}  `json:"priority"`
}
