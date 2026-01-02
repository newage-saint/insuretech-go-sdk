package models

import (
	"time"
)

// ApiKeyUsage represents a api_key_usage
type ApiKeyUsage struct {
	HttpMethod      string    `json:"http_method"`
	StatusCode      int       `json:"status_code"`
	ResponseTimeMs  int       `json:"response_time_ms,omitempty"`
	RequestPayload  string    `json:"request_payload,omitempty"`
	ResponsePayload string    `json:"response_payload,omitempty"`
	Id              string    `json:"id"`
	Endpoint        string    `json:"endpoint"`
	RequestIp       string    `json:"request_ip,omitempty"`
	UserAgent       string    `json:"user_agent,omitempty"`
	TraceId         string    `json:"trace_id,omitempty"`
	Timestamp       time.Time `json:"timestamp"`
	ApiKeyId        string    `json:"api_key_id"`
}
