package models

import (
	"time"
)

// PolicyServiceRequest represents a policy_service_request
type PolicyServiceRequest struct {
	RequestId   string                `json:"request_id"`
	PolicyId    string                `json:"policy_id"`
	CustomerId  string                `json:"customer_id"`
	RequestType *ServiceRequestType   `json:"request_type,omitempty"`
	RequestData string                `json:"request_data,omitempty"`
	ProcessedAt time.Time             `json:"processed_at,omitempty"`
	Status      *ServiceRequestStatus `json:"status,omitempty"`
	ProcessedBy string                `json:"processed_by,omitempty"`
	CreatedAt   time.Time             `json:"created_at,omitempty"`
}
