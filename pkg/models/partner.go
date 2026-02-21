package models

import (
	"time"
)

// Partner represents a partner
type Partner struct {
	ContactEmail              string               `json:"contact_email"`
	DeletedAt                 time.Time            `json:"deleted_at,omitempty"`
	Commission                *CommissionStructure `json:"commission,omitempty"`
	OrganizationName          string               `json:"organization_name"`
	Status                    interface{}          `json:"status"`
	OnboardedAt               time.Time            `json:"onboarded_at,omitempty"`
	CreatedAt                 time.Time            `json:"created_at"`
	UpdatedAt                 time.Time            `json:"updated_at"`
	ContactPhone              string               `json:"contact_phone"`
	ClaimsAssistanceRate      float64              `json:"claims_assistance_rate,omitempty"`
	Benefits                  *PartnerBenefits     `json:"benefits,omitempty"`
	PartnerId                 string               `json:"partner_id"`
	BankName                  string               `json:"bank_name,omitempty"`
	AcquisitionCommissionRate float64              `json:"acquisition_commission_rate"`
	RenewalCommissionRate     float64              `json:"renewal_commission_rate"`
	FocalPersonId             string               `json:"focal_person_id,omitempty"`
	Type                      *PartnerType         `json:"type"`
	TradeLicense              string               `json:"trade_license"`
	TinNumber                 string               `json:"tin_number"`
	BankAccount               string               `json:"bank_account"`
	BankBranch                string               `json:"bank_branch,omitempty"`
}
