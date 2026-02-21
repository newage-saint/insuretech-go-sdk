package models

// FraudCaseCreationRequest represents a fraud_case_creation_request
type FraudCaseCreationRequest struct {
	InvestigationNotes string `json:"investigation_notes,omitempty"`
	FraudAlertId       string `json:"fraud_alert_id"`
	Priority           string `json:"priority,omitempty"`
	InvestigatorId     string `json:"investigator_id"`
}
