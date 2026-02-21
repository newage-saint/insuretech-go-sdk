package models

import (
	"time"
)

// FraudAlert represents a fraud_alert
type FraudAlert struct {
	EntityId    string      `json:"entity_id"`
	RiskLevel   string      `json:"risk_level"`
	FraudScore  int         `json:"fraud_score"`
	ResolvedAt  time.Time   `json:"resolved_at,omitempty"`
	EntityType  string      `json:"entity_type"`
	FraudRuleId string      `json:"fraud_rule_id"`
	Details     string      `json:"details,omitempty"`
	Status      interface{} `json:"status"`
	AssignedTo  string      `json:"assigned_to,omitempty"`
	AuditInfo   interface{} `json:"audit_info"`
	Id          string      `json:"id"`
	AlertNumber string      `json:"alert_number"`
}
