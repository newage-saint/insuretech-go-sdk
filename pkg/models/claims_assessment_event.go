package models

import (
	"time"
)

// ClaimsAssessmentEvent represents a claims_assessment_event
type ClaimsAssessmentEvent struct {
	AgentId           string    `json:"agent_id,omitempty"`
	RecommendedAmount *Money    `json:"recommended_amount,omitempty"`
	Reasons           []string  `json:"reasons,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
	CorrelationId     string    `json:"correlation_id,omitempty"`
	ClaimId           string    `json:"claim_id,omitempty"`
	AssessmentResult  string    `json:"assessment_result,omitempty"`
	ConfidenceScore   float64   `json:"confidence_score,omitempty"`
	EventId           string    `json:"event_id,omitempty"`
}
