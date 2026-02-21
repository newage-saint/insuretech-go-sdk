package models

import (
	"time"
)

// ApiKeyUsage represents a api_key_usage
type ApiKeyUsage struct {
	ApiKeyId        string    `json:"api_key_id"`
	HttpMethod      string    `json:"http_method"`
	RequestIp       string    `json:"request_ip,omitempty"`
	UserAgent       string    `json:"user_agent,omitempty"`
	ResponsePayload string    `json:"response_payload,omitempty"`
	Endpoint        string    `json:"endpoint"`
	StatusCode      int       `json:"status_code"`
	ResponseTimeMs  int       `json:"response_time_ms,omitempty"`
	RequestPayload  string    `json:"request_payload,omitempty"`
	TraceId         string    `json:"trace_id,omitempty"`
	Timestamp       time.Time `json:"timestamp"`
	Id              string    `json:"id"`
}
