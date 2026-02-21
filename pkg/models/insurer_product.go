package models

import (
	"time"
)

// InsurerProduct represents a insurer_product
type InsurerProduct struct {
	MaxPremium          *Money      `json:"max_premium,omitempty"`
	MinTermYears        int         `json:"min_term_years,omitempty"`
	MedicalRequired     bool        `json:"medical_required,omitempty"`
	Status              interface{} `json:"status"`
	MaxSumAssured       *Money      `json:"max_sum_assured,omitempty"`
	MinEntryAge         int         `json:"min_entry_age,omitempty"`
	MaxEntryAge         int         `json:"max_entry_age,omitempty"`
	MaxMaturityAge      int         `json:"max_maturity_age,omitempty"`
	MedicalThreshold    *Money      `json:"medical_threshold,omitempty"`
	FreeLookPeriodDays  int         `json:"free_look_period_days,omitempty"`
	CommissionConfigId  string      `json:"commission_config_id,omitempty"`
	InsurerId           string      `json:"insurer_id"`
	MinSumAssured       *Money      `json:"min_sum_assured,omitempty"`
	MinPremium          *Money      `json:"min_premium,omitempty"`
	PremiumPaymentModes []string    `json:"premium_payment_modes,omitempty"`
	Features            string      `json:"features,omitempty"`
	Exclusions          string      `json:"exclusions,omitempty"`
	EffectiveTo         time.Time   `json:"effective_to,omitempty"`
	AuditInfo           interface{} `json:"audit_info"`
	MaxTermYears        int         `json:"max_term_years,omitempty"`
	EffectiveFrom       time.Time   `json:"effective_from"`
	Id                  string      `json:"id"`
	ProductId           string      `json:"product_id"`
	Code                string      `json:"code"`
	Name                string      `json:"name"`
}
