package models

// Tenant represents a tenant
type Tenant struct {
	Name           string      `json:"name"`
	Config         string      `json:"config,omitempty"`
	Branding       string      `json:"branding,omitempty"`
	TenantId       string      `json:"tenant_id"`
	Code           string      `json:"code"`
	Type           *TenantType `json:"type"`
	Status         interface{} `json:"status"`
	ParentTenantId string      `json:"parent_tenant_id,omitempty"`
	AuditInfo      interface{} `json:"audit_info"`
}
