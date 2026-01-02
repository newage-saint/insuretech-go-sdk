package models

// FraudRuleActivationResponse represents a fraud_rule_activation_response
type FraudRuleActivationResponse struct {
	Message string `json:"message,omitempty"`
	Error   *Error `json:"error,omitempty"`
}
