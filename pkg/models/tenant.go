package models

// Tenant represents a tenant
type Tenant struct {
	Branding       string      `json:"branding,omitempty"`
	AuditInfo      interface{} `json:"audit_info"`
	Name           string      `json:"name"`
	Status         interface{} `json:"status"`
	ParentTenantId string      `json:"parent_tenant_id,omitempty"`
	TenantId       string      `json:"tenant_id"`
	Code           string      `json:"code"`
	Type           *TenantType `json:"type"`
	Config         string      `json:"config,omitempty"`
}
