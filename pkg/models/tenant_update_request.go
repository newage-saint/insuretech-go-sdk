package models

// TenantUpdateRequest represents a tenant_update_request
type TenantUpdateRequest struct {
	Status   string `json:"status,omitempty"`
	Config   string `json:"config,omitempty"`
	Branding string `json:"branding,omitempty"`
	TenantId string `json:"tenant_id"`
	Name     string `json:"name"`
}
