package models

import (
	"time"
)

// ApiKeyUsage represents a api_key_usage
type ApiKeyUsage struct {
	UserAgent       string    `json:"user_agent,omitempty"`
	Id              string    `json:"id"`
	HttpMethod      string    `json:"http_method"`
	StatusCode      int       `json:"status_code"`
	RequestPayload  string    `json:"request_payload,omitempty"`
	ResponsePayload string    `json:"response_payload,omitempty"`
	TraceId         string    `json:"trace_id,omitempty"`
	Timestamp       time.Time `json:"timestamp"`
	ApiKeyId        string    `json:"api_key_id"`
	Endpoint        string    `json:"endpoint"`
	ResponseTimeMs  int       `json:"response_time_ms,omitempty"`
	RequestIp       string    `json:"request_ip,omitempty"`
}
