package models

import (
	"time"
)

// Session represents a session
type Session struct {
	UserId                string           `json:"user_id,omitempty"`
	SessionType           *SessionType     `json:"session_type,omitempty"`
	SessionTokenHash      string           `json:"session_token_hash,omitempty"`
	AccessTokenJti        string           `json:"access_token_jti,omitempty"`
	RefreshTokenJti       string           `json:"refresh_token_jti,omitempty"`
	AccessTokenExpiresAt  time.Time        `json:"access_token_expires_at,omitempty"`
	DeviceId              string           `json:"device_id,omitempty"`
	RefreshTokenExpiresAt time.Time        `json:"refresh_token_expires_at,omitempty"`
	CreatedAt             time.Time        `json:"created_at,omitempty"`
	IsActive              bool             `json:"is_active,omitempty"`
	IpAddress             string           `json:"ip_address,omitempty"`
	ExpiresAt             time.Time        `json:"expires_at,omitempty"`
	LastActivityAt        time.Time        `json:"last_activity_at,omitempty"`
	SessionId             string           `json:"session_id,omitempty"`
	UserAgent             string           `json:"user_agent,omitempty"`
	CsrfToken             string           `json:"csrf_token,omitempty"`
	DeviceName            string           `json:"device_name,omitempty"`
	DeviceType            *AuthnDeviceType `json:"device_type,omitempty"`
}
