package models

// FraudRuleUpdateRequest represents a fraud_rule_update_request
type FraudRuleUpdateRequest struct {
	FraudRule *FraudRule `json:"fraud_rule,omitempty"`
	RuleId    string     `json:"rule_id"`
}
