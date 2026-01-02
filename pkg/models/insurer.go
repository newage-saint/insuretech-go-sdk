package models

import (
	"time"
)

// Insurer represents a insurer
type Insurer struct {
	WebsiteUrl         string       `json:"website_url,omitempty"`
	Name               string       `json:"name"`
	NameBn             string       `json:"name_bn,omitempty"`
	Status             interface{}  `json:"status"`
	TinNumber          string       `json:"tin_number,omitempty"`
	RegisteredAddress  *Address     `json:"registered_address,omitempty"`
	Type               *InsurerType `json:"type"`
	TradeLicenseNumber string       `json:"trade_license_number,omitempty"`
	ContactInfo        *ContactInfo `json:"contact_info,omitempty"`
	HeadOfficeAddress  *Address     `json:"head_office_address,omitempty"`
	LogoUrl            string       `json:"logo_url,omitempty"`
	FinancialRating    string       `json:"financial_rating,omitempty"`
	AuditInfo          *AuditInfo   `json:"audit_info,omitempty"`
	IdraLicenseExpiry  time.Time    `json:"idra_license_expiry,omitempty"`
	PaidUpCapital      *Money       `json:"paid_up_capital,omitempty"`
	Id                 string       `json:"id"`
	Code               string       `json:"code"`
	IdraLicenseNumber  string       `json:"idra_license_number,omitempty"`
}
