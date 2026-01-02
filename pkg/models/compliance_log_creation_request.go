package models

// ComplianceLogCreationRequest represents a compliance_log_creation_request
type ComplianceLogCreationRequest struct {
	EntityType  string `json:"entity_type"`
	EntityId    string `json:"entity_id"`
	Status      string `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
	Evidence    string `json:"evidence,omitempty"`
	Type        string `json:"type"`
	Regulation  string `json:"regulation,omitempty"`
}
