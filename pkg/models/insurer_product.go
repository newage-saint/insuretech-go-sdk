package models

import (
	"time"
)

// InsurerProduct represents a insurer_product
type InsurerProduct struct {
	ProductId           string      `json:"product_id"`
	Code                string      `json:"code"`
	Status              interface{} `json:"status"`
	MaxSumAssured       *Money      `json:"max_sum_assured,omitempty"`
	MaxMaturityAge      int         `json:"max_maturity_age,omitempty"`
	MinTermYears        int         `json:"min_term_years,omitempty"`
	FreeLookPeriodDays  int         `json:"free_look_period_days,omitempty"`
	Exclusions          string      `json:"exclusions,omitempty"`
	InsurerId           string      `json:"insurer_id"`
	MaxEntryAge         int         `json:"max_entry_age,omitempty"`
	PremiumPaymentModes []string    `json:"premium_payment_modes,omitempty"`
	MedicalRequired     bool        `json:"medical_required,omitempty"`
	CommissionConfigId  string      `json:"commission_config_id,omitempty"`
	EffectiveTo         time.Time   `json:"effective_to,omitempty"`
	Id                  string      `json:"id"`
	MaxPremium          *Money      `json:"max_premium,omitempty"`
	MinEntryAge         int         `json:"min_entry_age,omitempty"`
	MaxTermYears        int         `json:"max_term_years,omitempty"`
	EffectiveFrom       time.Time   `json:"effective_from"`
	AuditInfo           interface{} `json:"audit_info"`
	Name                string      `json:"name"`
	MinSumAssured       *Money      `json:"min_sum_assured,omitempty"`
	MinPremium          *Money      `json:"min_premium,omitempty"`
	MedicalThreshold    *Money      `json:"medical_threshold,omitempty"`
	Features            string      `json:"features,omitempty"`
}
