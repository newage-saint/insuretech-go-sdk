package models

// FraudRule represents a fraud_rule
type FraudRule struct {
	Id          string          `json:"id"`
	Conditions  string          `json:"conditions"`
	RiskLevel   *FraudRiskLevel `json:"risk_level"`
	ScoreWeight int             `json:"score_weight"`
	Name        string          `json:"name"`
	Category    *RuleCategory   `json:"category"`
	Description string          `json:"description,omitempty"`
	IsActive    bool            `json:"is_active,omitempty"`
	AuditInfo   *AuditInfo      `json:"audit_info,omitempty"`
}
