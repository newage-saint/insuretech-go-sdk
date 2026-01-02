package models

import (
	"time"
)

// InsurerProduct represents a insurer_product
type InsurerProduct struct {
	MaxEntryAge         int         `json:"max_entry_age,omitempty"`
	MaxTermYears        int         `json:"max_term_years,omitempty"`
	MedicalThreshold    *Money      `json:"medical_threshold,omitempty"`
	EffectiveTo         time.Time   `json:"effective_to,omitempty"`
	Name                string      `json:"name"`
	MinSumAssured       *Money      `json:"min_sum_assured,omitempty"`
	MinPremium          *Money      `json:"min_premium,omitempty"`
	MinEntryAge         int         `json:"min_entry_age,omitempty"`
	MinTermYears        int         `json:"min_term_years,omitempty"`
	MedicalRequired     bool        `json:"medical_required,omitempty"`
	Features            string      `json:"features,omitempty"`
	Exclusions          string      `json:"exclusions,omitempty"`
	ProductId           string      `json:"product_id"`
	Code                string      `json:"code"`
	MaxSumAssured       *Money      `json:"max_sum_assured,omitempty"`
	PremiumPaymentModes []string    `json:"premium_payment_modes,omitempty"`
	EffectiveFrom       time.Time   `json:"effective_from"`
	AuditInfo           *AuditInfo  `json:"audit_info,omitempty"`
	Status              interface{} `json:"status"`
	MaxMaturityAge      int         `json:"max_maturity_age,omitempty"`
	FreeLookPeriodDays  int         `json:"free_look_period_days,omitempty"`
	CommissionConfigId  string      `json:"commission_config_id,omitempty"`
	Id                  string      `json:"id"`
	InsurerId           string      `json:"insurer_id"`
	MaxPremium          *Money      `json:"max_premium,omitempty"`
}
