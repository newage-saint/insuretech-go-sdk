package models

// TenantConfigRetrievalResponse represents a tenant_config_retrieval_response
type TenantConfigRetrievalResponse struct {
	Error   *Error          `json:"error,omitempty"`
	Configs []*TenantConfig `json:"configs,omitempty"`
}
