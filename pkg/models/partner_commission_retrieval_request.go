package models

import (
	"time"
)

// PartnerCommissionRetrievalRequest represents a partner_commission_retrieval_request
type PartnerCommissionRetrievalRequest struct {
	PartnerId string    `json:"partner_id"`
	StartDate time.Time `json:"start_date,omitempty"`
	EndDate   time.Time `json:"end_date,omitempty"`
}
