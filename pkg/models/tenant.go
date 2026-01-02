package models

// Tenant represents a tenant
type Tenant struct {
	Type           *TenantType `json:"type"`
	Status         interface{} `json:"status"`
	Branding       string      `json:"branding,omitempty"`
	AuditInfo      *AuditInfo  `json:"audit_info,omitempty"`
	Name           string      `json:"name"`
	ParentTenantId string      `json:"parent_tenant_id,omitempty"`
	Config         string      `json:"config,omitempty"`
	Id             string      `json:"id"`
	Code           string      `json:"code"`
}
