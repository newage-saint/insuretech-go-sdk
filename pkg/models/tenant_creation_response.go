package models

// TenantCreationResponse represents a tenant_creation_response
type TenantCreationResponse struct {
	Error    *Error `json:"error,omitempty"`
	TenantId string `json:"tenant_id,omitempty"`
	Message  string `json:"message,omitempty"`
}
