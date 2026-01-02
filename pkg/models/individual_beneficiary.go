package models

import (
	"time"
)

// IndividualBeneficiary represents a individual_beneficiary
type IndividualBeneficiary struct {
	PassportNumber         string             `json:"passport_number,omitempty"`
	Occupation             string             `json:"occupation,omitempty"`
	NomineeRelationship    string             `json:"nominee_relationship,omitempty"`
	Id                     string             `json:"id"`
	FullNameBn             string             `json:"full_name_bn,omitempty"`
	DateOfBirth            time.Time          `json:"date_of_birth"`
	BirthCertificateNumber string             `json:"birth_certificate_number,omitempty"`
	NomineeName            string             `json:"nominee_name,omitempty"`
	FullName               string             `json:"full_name"`
	MaritalStatus          *MaritalStatus     `json:"marital_status,omitempty"`
	PermanentAddress       *Address           `json:"permanent_address,omitempty"`
	BeneficiaryId          string             `json:"beneficiary_id"`
	Gender                 *BeneficiaryGender `json:"gender"`
	TinNumber              string             `json:"tin_number,omitempty"`
	ContactInfo            *ContactInfo       `json:"contact_info,omitempty"`
	PresentAddress         *Address           `json:"present_address,omitempty"`
	AuditInfo              *AuditInfo         `json:"audit_info,omitempty"`
	NidNumber              string             `json:"nid_number,omitempty"`
}
