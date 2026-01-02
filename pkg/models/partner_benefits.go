package models

import (
	"time"
)

// PartnerBenefits represents a partner_benefits
type PartnerBenefits struct {
	EffectiveFrom             time.Time `json:"effective_from,omitempty"`
	DiscountType              string    `json:"discount_type,omitempty"`
	CashlessEnabled           bool      `json:"cashless_enabled,omitempty"`
	AutoApprovalThreshold     string    `json:"auto_approval_threshold,omitempty"`
	PreAuthorizationRequired  bool      `json:"pre_authorization_required,omitempty"`
	DiscountEnabled           bool      `json:"discount_enabled,omitempty"`
	MinDiscount               float64   `json:"min_discount,omitempty"`
	MaxDiscount               float64   `json:"max_discount,omitempty"`
	NationwideCoverage        bool      `json:"nationwide_coverage,omitempty"`
	EffectiveTo               time.Time `json:"effective_to,omitempty"`
	RequiredDocuments         []string  `json:"required_documents,omitempty"`
	ServiceLocations          []string  `json:"service_locations,omitempty"`
	DiscountPercentage        float64   `json:"discount_percentage,omitempty"`
	CashlessLimit             string    `json:"cashless_limit,omitempty"`
	AuthorizationValidityDays int       `json:"authorization_validity_days,omitempty"`
	Notes                     string    `json:"notes,omitempty"`
}
