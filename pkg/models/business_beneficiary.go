package models

import (
	"time"
)

// BusinessBeneficiary represents a business_beneficiary
type BusinessBeneficiary struct {
	BusinessNameBn         string        `json:"business_name_bn,omitempty"`
	EmployeeCount          int           `json:"employee_count,omitempty"`
	IncorporationDate      time.Time     `json:"incorporation_date,omitempty"`
	RegisteredAddress      *Address      `json:"registered_address,omitempty"`
	BusinessName           string        `json:"business_name"`
	BinNumber              string        `json:"bin_number,omitempty"`
	ContactInfo            *ContactInfo  `json:"contact_info,omitempty"`
	BusinessAddress        *Address      `json:"business_address,omitempty"`
	FocalPersonName        string        `json:"focal_person_name"`
	Id                     string        `json:"id"`
	BeneficiaryId          string        `json:"beneficiary_id"`
	TradeLicenseNumber     string        `json:"trade_license_number"`
	TinNumber              string        `json:"tin_number"`
	BusinessType           *BusinessType `json:"business_type"`
	IndustrySector         string        `json:"industry_sector,omitempty"`
	FocalPersonDesignation string        `json:"focal_person_designation,omitempty"`
	FocalPersonNid         string        `json:"focal_person_nid,omitempty"`
	TradeLicenseIssueDate  time.Time     `json:"trade_license_issue_date,omitempty"`
	TradeLicenseExpiryDate time.Time     `json:"trade_license_expiry_date,omitempty"`
	FocalPersonContact     *ContactInfo  `json:"focal_person_contact,omitempty"`
	AuditInfo              *AuditInfo    `json:"audit_info,omitempty"`
}
