package models

import (
	"time"
)

// IndividualBeneficiary represents a individual_beneficiary
type IndividualBeneficiary struct {
	PresentAddress         *Address           `json:"present_address,omitempty"`
	AuditInfo              *AuditInfo         `json:"audit_info,omitempty"`
	BeneficiaryId          string             `json:"beneficiary_id,omitempty"`
	FullName               string             `json:"full_name,omitempty"`
	TinNumber              string             `json:"tin_number,omitempty"`
	MaritalStatus          *MaritalStatus     `json:"marital_status,omitempty"`
	Occupation             string             `json:"occupation,omitempty"`
	NomineeName            string             `json:"nominee_name,omitempty"`
	NomineeRelationship    string             `json:"nominee_relationship,omitempty"`
	PermanentAddress       *Address           `json:"permanent_address,omitempty"`
	FullNameBn             string             `json:"full_name_bn,omitempty"`
	Gender                 *BeneficiaryGender `json:"gender,omitempty"`
	ContactInfo            *ContactInfo       `json:"contact_info,omitempty"`
	DateOfBirth            time.Time          `json:"date_of_birth,omitempty"`
	NidNumber              string             `json:"nid_number,omitempty"`
	PassportNumber         string             `json:"passport_number,omitempty"`
	BirthCertificateNumber string             `json:"birth_certificate_number,omitempty"`
}
