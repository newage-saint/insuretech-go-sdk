package models

import (
	"time"
)

// PartnerBenefits represents a partner_benefits
type PartnerBenefits struct {
	DiscountPercentage        float64   `json:"discount_percentage,omitempty"`
	DiscountType              string    `json:"discount_type,omitempty"`
	CashlessEnabled           bool      `json:"cashless_enabled,omitempty"`
	AuthorizationValidityDays int       `json:"authorization_validity_days,omitempty"`
	RequiredDocuments         []string  `json:"required_documents,omitempty"`
	MinDiscount               float64   `json:"min_discount,omitempty"`
	PreAuthorizationRequired  bool      `json:"pre_authorization_required,omitempty"`
	NationwideCoverage        bool      `json:"nationwide_coverage,omitempty"`
	MaxDiscount               float64   `json:"max_discount,omitempty"`
	CashlessLimit             string    `json:"cashless_limit,omitempty"`
	ServiceLocations          []string  `json:"service_locations,omitempty"`
	EffectiveFrom             time.Time `json:"effective_from,omitempty"`
	DiscountEnabled           bool      `json:"discount_enabled,omitempty"`
	AutoApprovalThreshold     string    `json:"auto_approval_threshold,omitempty"`
	Notes                     string    `json:"notes,omitempty"`
	EffectiveTo               time.Time `json:"effective_to,omitempty"`
}
