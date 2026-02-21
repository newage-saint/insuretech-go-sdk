package models

import (
	"time"
)

// FraudAlert represents a fraud_alert
type FraudAlert struct {
	FraudRuleId string      `json:"fraud_rule_id"`
	Details     string      `json:"details,omitempty"`
	AssignedTo  string      `json:"assigned_to,omitempty"`
	ResolvedAt  time.Time   `json:"resolved_at,omitempty"`
	AuditInfo   interface{} `json:"audit_info"`
	Id          string      `json:"id"`
	RiskLevel   string      `json:"risk_level"`
	FraudScore  int         `json:"fraud_score"`
	Status      interface{} `json:"status"`
	AlertNumber string      `json:"alert_number"`
	EntityType  string      `json:"entity_type"`
	EntityId    string      `json:"entity_id"`
}
