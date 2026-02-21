package models

import (
	"time"
)

// ClaimDocument represents a claim_document
type ClaimDocument struct {
	DocumentId   string    `json:"document_id,omitempty"`
	ClaimId      string    `json:"claim_id,omitempty"`
	DocumentType string    `json:"document_type,omitempty"`
	FileUrl      string    `json:"file_url,omitempty"`
	FileHash     string    `json:"file_hash,omitempty"`
	Verified     bool      `json:"verified,omitempty"`
	VerifiedBy   string    `json:"verified_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UploadedAt   time.Time `json:"uploaded_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
