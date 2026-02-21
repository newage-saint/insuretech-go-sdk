package models

// RequestContext represents a request_context
type RequestContext struct {
	Tenant    *TenantContext         `json:"tenant,omitempty"`
	RequestId string                 `json:"request_id,omitempty"`
	Headers   map[string]interface{} `json:"headers,omitempty"`
}
