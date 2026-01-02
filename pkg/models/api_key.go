package models

import (
	"time"
)

// ApiKey represents a api_key
type ApiKey struct {
	RateLimitPerMinute int              `json:"rate_limit_per_minute"`
	IpWhitelist        []string         `json:"ip_whitelist,omitempty"`
	AuditInfo          *AuditInfo       `json:"audit_info,omitempty"`
	Id                 string           `json:"id"`
	KeyHash            string           `json:"key_hash"`
	Scopes             []string         `json:"scopes,omitempty"`
	ExpiresAt          time.Time        `json:"expires_at,omitempty"`
	LastUsedAt         time.Time        `json:"last_used_at,omitempty"`
	Name               string           `json:"name"`
	OwnerType          *ApiKeyOwnerType `json:"owner_type"`
	OwnerId            string           `json:"owner_id"`
	Status             interface{}      `json:"status"`
}
