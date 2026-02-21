package models

// TenantCreationResponse represents a tenant_creation_response
type TenantCreationResponse struct {
	Message  string `json:"message,omitempty"`
	Error    *Error `json:"error,omitempty"`
	TenantId string `json:"tenant_id,omitempty"`
}
