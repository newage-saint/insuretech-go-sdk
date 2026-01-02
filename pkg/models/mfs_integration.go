package models

// MFSIntegration represents a mfs_integration
type MFSIntegration struct {
	MerchantId     string      `json:"merchant_id,omitempty"`
	ApiCredentials string      `json:"api_credentials,omitempty"`
	IsActive       bool        `json:"is_active,omitempty"`
	Config         string      `json:"config,omitempty"`
	AuditInfo      *AuditInfo  `json:"audit_info,omitempty"`
	Id             string      `json:"id"`
	Provider       interface{} `json:"provider"`
	ApiBaseUrl     string      `json:"api_base_url"`
}
