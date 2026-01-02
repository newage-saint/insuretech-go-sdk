package models

// FraudRuleDeactivationResponse represents a fraud_rule_deactivation_response
type FraudRuleDeactivationResponse struct {
	Message string `json:"message,omitempty"`
	Error   *Error `json:"error,omitempty"`
}
