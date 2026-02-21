package models

import (
	"time"
)

// ApiKey represents a api_key
type ApiKey struct {
	OwnerId            string           `json:"owner_id"`
	Scopes             []string         `json:"scopes,omitempty"`
	ExpiresAt          time.Time        `json:"expires_at,omitempty"`
	LastUsedAt         time.Time        `json:"last_used_at,omitempty"`
	Status             interface{}      `json:"status"`
	RateLimitPerMinute int              `json:"rate_limit_per_minute"`
	IpWhitelist        []string         `json:"ip_whitelist,omitempty"`
	AuditInfo          interface{}      `json:"audit_info"`
	Id                 string           `json:"id"`
	KeyHash            string           `json:"key_hash"`
	Name               string           `json:"name"`
	OwnerType          *ApiKeyOwnerType `json:"owner_type"`
}
