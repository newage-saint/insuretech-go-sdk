package models

import (
	"time"
)

// Nominee represents a nominee
type Nominee struct {
	FullName            string    `json:"full_name,omitempty"`
	SharePercentage     float64   `json:"share_percentage,omitempty"`
	NidNumber           string    `json:"nid_number,omitempty"`
	PhoneNumber         string    `json:"phone_number,omitempty"`
	CreatedAt           time.Time `json:"created_at,omitempty"`
	Relationship        string    `json:"relationship,omitempty"`
	DateOfBirth         time.Time `json:"date_of_birth,omitempty"`
	UpdatedAt           time.Time `json:"updated_at,omitempty"`
	NomineeDobText      string    `json:"nominee_dob_text,omitempty"`
	NomineeSharePercent float64   `json:"nominee_share_percent,omitempty"`
	NomineeId           string    `json:"nominee_id,omitempty"`
	PolicyId            string    `json:"policy_id,omitempty"`
}
