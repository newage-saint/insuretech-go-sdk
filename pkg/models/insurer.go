package models

import (
	"time"
)

// Insurer represents a insurer
type Insurer struct {
	NameBn             string       `json:"name_bn,omitempty"`
	HeadOfficeAddress  interface{}  `json:"head_office_address"`
	LogoUrl            string       `json:"logo_url,omitempty"`
	WebsiteUrl         string       `json:"website_url,omitempty"`
	Type               *InsurerType `json:"type"`
	IdraLicenseExpiry  time.Time    `json:"idra_license_expiry,omitempty"`
	AuditInfo          interface{}  `json:"audit_info"`
	Code               string       `json:"code"`
	Name               string       `json:"name"`
	Status             interface{}  `json:"status"`
	IdraLicenseNumber  string       `json:"idra_license_number,omitempty"`
	ContactInfo        interface{}  `json:"contact_info"`
	Id                 string       `json:"id"`
	TradeLicenseNumber string       `json:"trade_license_number,omitempty"`
	TinNumber          string       `json:"tin_number,omitempty"`
	RegisteredAddress  interface{}  `json:"registered_address"`
	FinancialRating    string       `json:"financial_rating,omitempty"`
	PaidUpCapital      *Money       `json:"paid_up_capital,omitempty"`
}
