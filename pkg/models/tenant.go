package models

// Tenant represents a tenant
type Tenant struct {
	TenantId       string      `json:"tenant_id"`
	Code           string      `json:"code"`
	ParentTenantId string      `json:"parent_tenant_id,omitempty"`
	Config         string      `json:"config,omitempty"`
	Name           string      `json:"name"`
	Type           *TenantType `json:"type"`
	Status         interface{} `json:"status"`
	Branding       string      `json:"branding,omitempty"`
	AuditInfo      interface{} `json:"audit_info"`
}
