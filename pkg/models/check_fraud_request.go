package models

// CheckFraudRequest represents a check_fraud_request
type CheckFraudRequest struct {
	EntityId   string                 `json:"entity_id"`
	Data       map[string]interface{} `json:"data,omitempty"`
	EntityType string                 `json:"entity_type"`
}
