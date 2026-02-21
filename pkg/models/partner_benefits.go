package models

import (
	"time"
)

// PartnerBenefits represents a partner_benefits
type PartnerBenefits struct {
	MaxDiscount               float64   `json:"max_discount,omitempty"`
	AutoApprovalThreshold     string    `json:"auto_approval_threshold,omitempty"`
	PreAuthorizationRequired  bool      `json:"pre_authorization_required,omitempty"`
	ServiceLocations          []string  `json:"service_locations,omitempty"`
	EffectiveTo               time.Time `json:"effective_to,omitempty"`
	DiscountType              string    `json:"discount_type,omitempty"`
	CashlessLimit             string    `json:"cashless_limit,omitempty"`
	NationwideCoverage        bool      `json:"nationwide_coverage,omitempty"`
	EffectiveFrom             time.Time `json:"effective_from,omitempty"`
	AuthorizationValidityDays int       `json:"authorization_validity_days,omitempty"`
	Notes                     string    `json:"notes,omitempty"`
	DiscountEnabled           bool      `json:"discount_enabled,omitempty"`
	CashlessEnabled           bool      `json:"cashless_enabled,omitempty"`
	RequiredDocuments         []string  `json:"required_documents,omitempty"`
	DiscountPercentage        float64   `json:"discount_percentage,omitempty"`
	MinDiscount               float64   `json:"min_discount,omitempty"`
}
