package models

import (
	"time"
)

// DocumentGeneration represents a document_generation
type DocumentGeneration struct {
	FileUrl            string      `json:"file_url,omitempty"`
	FileSizeBytes      string      `json:"file_size_bytes,omitempty"`
	QrCodeData         string      `json:"qr_code_data,omitempty"`
	GeneratedAt        time.Time   `json:"generated_at,omitempty"`
	Id                 string      `json:"id"`
	DocumentTemplateId string      `json:"document_template_id"`
	GeneratedBy        string      `json:"generated_by,omitempty"`
	AuditInfo          interface{} `json:"audit_info"`
	EntityType         string      `json:"entity_type"`
	EntityId           string      `json:"entity_id"`
	Data               string      `json:"data"`
	Status             interface{} `json:"status"`
}
