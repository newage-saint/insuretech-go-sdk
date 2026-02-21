package models

// FraudRule represents a fraud_rule
type FraudRule struct {
	FraudRuleId string          `json:"fraud_rule_id"`
	Description string          `json:"description,omitempty"`
	Conditions  string          `json:"conditions"`
	AuditInfo   interface{}     `json:"audit_info"`
	Name        string          `json:"name"`
	Category    *RuleCategory   `json:"category"`
	RiskLevel   *FraudRiskLevel `json:"risk_level"`
	ScoreWeight int             `json:"score_weight"`
	IsActive    bool            `json:"is_active,omitempty"`
}
