package models

import (
	"time"
)

// Nominee represents a nominee
type Nominee struct {
	SharePercentage float64   `json:"share_percentage"`
	DateOfBirth     time.Time `json:"date_of_birth"`
	PhoneNumber     string    `json:"phone_number,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	Relationship    string    `json:"relationship"`
	NidNumber       string    `json:"nid_number,omitempty"`
	UpdatedAt       time.Time `json:"updated_at"`
	NomineeId       string    `json:"nominee_id"`
	PolicyId        string    `json:"policy_id"`
	FullName        string    `json:"full_name"`
}
