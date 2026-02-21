package models

import (
	"time"
)

// ApiKey represents a api_key
type ApiKey struct {
	KeyHash            string           `json:"key_hash"`
	Name               string           `json:"name"`
	OwnerType          *ApiKeyOwnerType `json:"owner_type"`
	Scopes             []string         `json:"scopes,omitempty"`
	LastUsedAt         time.Time        `json:"last_used_at,omitempty"`
	IpWhitelist        []string         `json:"ip_whitelist,omitempty"`
	Id                 string           `json:"id"`
	OwnerId            string           `json:"owner_id"`
	Status             interface{}      `json:"status"`
	RateLimitPerMinute int              `json:"rate_limit_per_minute"`
	ExpiresAt          time.Time        `json:"expires_at,omitempty"`
	AuditInfo          interface{}      `json:"audit_info"`
}
