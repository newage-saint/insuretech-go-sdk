package models

import (
	"time"
)

// Partner represents a partner
type Partner struct {
	OrganizationName          string               `json:"organization_name"`
	Status                    interface{}          `json:"status"`
	TinNumber                 string               `json:"tin_number"`
	ContactPhone              string               `json:"contact_phone"`
	BankName                  string               `json:"bank_name,omitempty"`
	RenewalCommissionRate     float64              `json:"renewal_commission_rate"`
	OnboardedAt               time.Time            `json:"onboarded_at,omitempty"`
	DeletedAt                 time.Time            `json:"deleted_at,omitempty"`
	Type                      *PartnerType         `json:"type"`
	TradeLicense              string               `json:"trade_license"`
	BankAccount               string               `json:"bank_account"`
	UpdatedAt                 time.Time            `json:"updated_at"`
	BankBranch                string               `json:"bank_branch,omitempty"`
	ContactEmail              string               `json:"contact_email"`
	AcquisitionCommissionRate float64              `json:"acquisition_commission_rate"`
	ClaimsAssistanceRate      float64              `json:"claims_assistance_rate,omitempty"`
	CreatedAt                 time.Time            `json:"created_at"`
	FocalPersonId             string               `json:"focal_person_id,omitempty"`
	Commission                *CommissionStructure `json:"commission,omitempty"`
	Benefits                  *PartnerBenefits     `json:"benefits,omitempty"`
	PartnerId                 string               `json:"partner_id"`
}
