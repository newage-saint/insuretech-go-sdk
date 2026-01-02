package models

import (
	"time"
)

// FraudAlert represents a fraud_alert
type FraudAlert struct {
	FraudScore  int         `json:"fraud_score"`
	AssignedTo  string      `json:"assigned_to,omitempty"`
	Id          string      `json:"id"`
	EntityType  string      `json:"entity_type"`
	FraudRuleId string      `json:"fraud_rule_id"`
	Details     string      `json:"details,omitempty"`
	Status      interface{} `json:"status"`
	ResolvedAt  time.Time   `json:"resolved_at,omitempty"`
	AuditInfo   *AuditInfo  `json:"audit_info,omitempty"`
	AlertNumber string      `json:"alert_number"`
	EntityId    string      `json:"entity_id"`
	RiskLevel   string      `json:"risk_level"`
}
