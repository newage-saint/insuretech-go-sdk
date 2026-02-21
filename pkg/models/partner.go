package models

import (
	"time"
)

// Partner represents a partner
type Partner struct {
	OrganizationName          string               `json:"organization_name"`
	ContactPhone              string               `json:"contact_phone"`
	Commission                *CommissionStructure `json:"commission,omitempty"`
	PartnerId                 string               `json:"partner_id"`
	TinNumber                 string               `json:"tin_number"`
	ClaimsAssistanceRate      float64              `json:"claims_assistance_rate,omitempty"`
	OnboardedAt               time.Time            `json:"onboarded_at,omitempty"`
	CreatedAt                 time.Time            `json:"created_at"`
	Benefits                  *PartnerBenefits     `json:"benefits,omitempty"`
	TradeLicense              string               `json:"trade_license"`
	BankName                  string               `json:"bank_name,omitempty"`
	BankBranch                string               `json:"bank_branch,omitempty"`
	ContactEmail              string               `json:"contact_email"`
	AcquisitionCommissionRate float64              `json:"acquisition_commission_rate"`
	UpdatedAt                 time.Time            `json:"updated_at"`
	FocalPersonId             string               `json:"focal_person_id,omitempty"`
	Type                      *PartnerType         `json:"type"`
	Status                    interface{}          `json:"status"`
	BankAccount               string               `json:"bank_account"`
	RenewalCommissionRate     float64              `json:"renewal_commission_rate"`
	DeletedAt                 time.Time            `json:"deleted_at,omitempty"`
}
