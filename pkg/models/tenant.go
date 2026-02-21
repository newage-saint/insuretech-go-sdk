package models

// Tenant represents a tenant
type Tenant struct {
	Type           *TenantType `json:"type"`
	ParentTenantId string      `json:"parent_tenant_id,omitempty"`
	Branding       string      `json:"branding,omitempty"`
	AuditInfo      interface{} `json:"audit_info"`
	TenantId       string      `json:"tenant_id"`
	Status         interface{} `json:"status"`
	Config         string      `json:"config,omitempty"`
	Code           string      `json:"code"`
	Name           string      `json:"name"`
}
