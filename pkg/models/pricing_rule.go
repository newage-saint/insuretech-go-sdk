package models

// PricingRule represents a pricing_rule
type PricingRule struct {
	Conditions []*RuleCondition `json:"conditions,omitempty"`
	Action     *RuleAction      `json:"action,omitempty"`
	RuleId     string           `json:"rule_id,omitempty"`
	RuleName   string           `json:"rule_name,omitempty"`
	Type       *RuleType        `json:"type,omitempty"`
}
