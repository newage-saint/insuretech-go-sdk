package models

import (
	"time"
)

// Partner represents a partner
type Partner struct {
	OrganizationName          string               `json:"organization_name"`
	BankName                  string               `json:"bank_name,omitempty"`
	RenewalCommissionRate     float64              `json:"renewal_commission_rate"`
	FocalPersonId             string               `json:"focal_person_id,omitempty"`
	Commission                *CommissionStructure `json:"commission,omitempty"`
	ClaimsAssistanceRate      float64              `json:"claims_assistance_rate,omitempty"`
	Status                    interface{}          `json:"status"`
	BankBranch                string               `json:"bank_branch,omitempty"`
	ContactEmail              string               `json:"contact_email"`
	OnboardedAt               time.Time            `json:"onboarded_at,omitempty"`
	CreatedAt                 time.Time            `json:"created_at"`
	UpdatedAt                 time.Time            `json:"updated_at"`
	AcquisitionCommissionRate float64              `json:"acquisition_commission_rate"`
	DeletedAt                 time.Time            `json:"deleted_at,omitempty"`
	Benefits                  *PartnerBenefits     `json:"benefits,omitempty"`
	ContactPhone              string               `json:"contact_phone"`
	PartnerId                 string               `json:"partner_id"`
	Type                      *PartnerType         `json:"type"`
	TradeLicense              string               `json:"trade_license"`
	TinNumber                 string               `json:"tin_number"`
	BankAccount               string               `json:"bank_account"`
}
