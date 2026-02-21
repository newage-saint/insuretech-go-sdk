package models

import (
	"time"
)

// InsurerProduct represents a insurer_product
type InsurerProduct struct {
	ProductId           string      `json:"product_id"`
	Status              interface{} `json:"status"`
	MinEntryAge         int         `json:"min_entry_age,omitempty"`
	MaxTermYears        int         `json:"max_term_years,omitempty"`
	Features            string      `json:"features,omitempty"`
	EffectiveFrom       time.Time   `json:"effective_from"`
	MinPremium          *Money      `json:"min_premium,omitempty"`
	MaxPremium          *Money      `json:"max_premium,omitempty"`
	Exclusions          string      `json:"exclusions,omitempty"`
	AuditInfo           interface{} `json:"audit_info"`
	Id                  string      `json:"id"`
	InsurerId           string      `json:"insurer_id"`
	MaxSumAssured       *Money      `json:"max_sum_assured,omitempty"`
	MaxEntryAge         int         `json:"max_entry_age,omitempty"`
	FreeLookPeriodDays  int         `json:"free_look_period_days,omitempty"`
	CommissionConfigId  string      `json:"commission_config_id,omitempty"`
	EffectiveTo         time.Time   `json:"effective_to,omitempty"`
	Code                string      `json:"code"`
	Name                string      `json:"name"`
	MinSumAssured       *Money      `json:"min_sum_assured,omitempty"`
	MaxMaturityAge      int         `json:"max_maturity_age,omitempty"`
	MinTermYears        int         `json:"min_term_years,omitempty"`
	PremiumPaymentModes []string    `json:"premium_payment_modes,omitempty"`
	MedicalRequired     bool        `json:"medical_required,omitempty"`
	MedicalThreshold    *Money      `json:"medical_threshold,omitempty"`
}
