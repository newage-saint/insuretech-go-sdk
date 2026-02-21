package models

import (
	"time"
)

// DocumentGeneration represents a document_generation
type DocumentGeneration struct {
	EntityType         string      `json:"entity_type"`
	EntityId           string      `json:"entity_id"`
	Status             interface{} `json:"status"`
	FileUrl            string      `json:"file_url,omitempty"`
	FileSizeBytes      string      `json:"file_size_bytes,omitempty"`
	QrCodeData         string      `json:"qr_code_data,omitempty"`
	Id                 string      `json:"id"`
	DocumentTemplateId string      `json:"document_template_id"`
	Data               string      `json:"data"`
	GeneratedBy        string      `json:"generated_by,omitempty"`
	GeneratedAt        time.Time   `json:"generated_at,omitempty"`
	AuditInfo          interface{} `json:"audit_info"`
}
