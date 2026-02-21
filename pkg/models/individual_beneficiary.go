package models

import (
	"time"
)

// IndividualBeneficiary represents a individual_beneficiary
type IndividualBeneficiary struct {
	PermanentAddress       *Address           `json:"permanent_address,omitempty"`
	NomineeName            string             `json:"nominee_name,omitempty"`
	FullName               string             `json:"full_name,omitempty"`
	DateOfBirth            time.Time          `json:"date_of_birth,omitempty"`
	BirthCertificateNumber string             `json:"birth_certificate_number,omitempty"`
	MaritalStatus          *MaritalStatus     `json:"marital_status,omitempty"`
	NomineeRelationship    string             `json:"nominee_relationship,omitempty"`
	AuditInfo              *AuditInfo         `json:"audit_info,omitempty"`
	BeneficiaryId          string             `json:"beneficiary_id,omitempty"`
	NidNumber              string             `json:"nid_number,omitempty"`
	PassportNumber         string             `json:"passport_number,omitempty"`
	TinNumber              string             `json:"tin_number,omitempty"`
	Occupation             string             `json:"occupation,omitempty"`
	FullNameBn             string             `json:"full_name_bn,omitempty"`
	Gender                 *BeneficiaryGender `json:"gender,omitempty"`
	PresentAddress         *Address           `json:"present_address,omitempty"`
	ContactInfo            *ContactInfo       `json:"contact_info,omitempty"`
}
