package models

// RequestContext represents a request_context
type RequestContext struct {
	RequestId string                 `json:"request_id,omitempty"`
	Headers   map[string]interface{} `json:"headers,omitempty"`
	Tenant    *TenantContext         `json:"tenant,omitempty"`
}
