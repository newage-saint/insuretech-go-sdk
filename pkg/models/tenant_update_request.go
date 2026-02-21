package models

// TenantUpdateRequest represents a tenant_update_request
type TenantUpdateRequest struct {
	Name     string `json:"name"`
	Status   string `json:"status,omitempty"`
	Config   string `json:"config,omitempty"`
	Branding string `json:"branding,omitempty"`
	TenantId string `json:"tenant_id"`
}
