package models

// MFSIntegration represents a mfs_integration
type MFSIntegration struct {
	Config         string      `json:"config,omitempty"`
	AuditInfo      interface{} `json:"audit_info"`
	Id             string      `json:"id"`
	Provider       interface{} `json:"provider"`
	ApiBaseUrl     string      `json:"api_base_url"`
	MerchantId     string      `json:"merchant_id,omitempty"`
	ApiCredentials string      `json:"api_credentials,omitempty"`
	IsActive       bool        `json:"is_active,omitempty"`
}
