package models

import (
	"time"
)

// IndividualBeneficiary represents a individual_beneficiary
type IndividualBeneficiary struct {
	DateOfBirth            time.Time          `json:"date_of_birth,omitempty"`
	TinNumber              string             `json:"tin_number,omitempty"`
	MaritalStatus          *MaritalStatus     `json:"marital_status,omitempty"`
	FullName               string             `json:"full_name,omitempty"`
	Gender                 *BeneficiaryGender `json:"gender,omitempty"`
	PassportNumber         string             `json:"passport_number,omitempty"`
	NomineeName            string             `json:"nominee_name,omitempty"`
	BeneficiaryId          string             `json:"beneficiary_id,omitempty"`
	FullNameBn             string             `json:"full_name_bn,omitempty"`
	NidNumber              string             `json:"nid_number,omitempty"`
	Occupation             string             `json:"occupation,omitempty"`
	ContactInfo            *ContactInfo       `json:"contact_info,omitempty"`
	PresentAddress         *Address           `json:"present_address,omitempty"`
	NomineeRelationship    string             `json:"nominee_relationship,omitempty"`
	BirthCertificateNumber string             `json:"birth_certificate_number,omitempty"`
	PermanentAddress       *Address           `json:"permanent_address,omitempty"`
	AuditInfo              *AuditInfo         `json:"audit_info,omitempty"`
}
