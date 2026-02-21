package models

import (
	"time"
)

// PartnerBenefits represents a partner_benefits
type PartnerBenefits struct {
	AutoApprovalThreshold     string    `json:"auto_approval_threshold,omitempty"`
	PreAuthorizationRequired  bool      `json:"pre_authorization_required,omitempty"`
	DiscountEnabled           bool      `json:"discount_enabled,omitempty"`
	MaxDiscount               float64   `json:"max_discount,omitempty"`
	DiscountType              string    `json:"discount_type,omitempty"`
	NationwideCoverage        bool      `json:"nationwide_coverage,omitempty"`
	Notes                     string    `json:"notes,omitempty"`
	MinDiscount               float64   `json:"min_discount,omitempty"`
	CashlessLimit             string    `json:"cashless_limit,omitempty"`
	AuthorizationValidityDays int       `json:"authorization_validity_days,omitempty"`
	EffectiveTo               time.Time `json:"effective_to,omitempty"`
	DiscountPercentage        float64   `json:"discount_percentage,omitempty"`
	CashlessEnabled           bool      `json:"cashless_enabled,omitempty"`
	RequiredDocuments         []string  `json:"required_documents,omitempty"`
	ServiceLocations          []string  `json:"service_locations,omitempty"`
	EffectiveFrom             time.Time `json:"effective_from,omitempty"`
}
