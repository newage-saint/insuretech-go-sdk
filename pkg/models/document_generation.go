package models

import (
	"time"
)

// DocumentGeneration represents a document_generation
type DocumentGeneration struct {
	FileUrl            string      `json:"file_url,omitempty"`
	QrCodeData         string      `json:"qr_code_data,omitempty"`
	GeneratedAt        time.Time   `json:"generated_at,omitempty"`
	AuditInfo          *AuditInfo  `json:"audit_info,omitempty"`
	EntityType         string      `json:"entity_type"`
	EntityId           string      `json:"entity_id"`
	Status             interface{} `json:"status"`
	FileSizeBytes      string      `json:"file_size_bytes,omitempty"`
	GeneratedBy        string      `json:"generated_by,omitempty"`
	Id                 string      `json:"id"`
	DocumentTemplateId string      `json:"document_template_id"`
	Data               string      `json:"data"`
}
