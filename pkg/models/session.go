package models

import (
	"time"
)

// Session represents a session
type Session struct {
	IpAddress             string           `json:"ip_address,omitempty"`
	CreatedAt             time.Time        `json:"created_at"`
	SessionId             string           `json:"session_id"`
	UserId                string           `json:"user_id"`
	RefreshToken          string           `json:"refresh_token"`
	UserAgent             string           `json:"user_agent,omitempty"`
	DeviceId              string           `json:"device_id,omitempty"`
	DeviceType            *AuthnDeviceType `json:"device_type"`
	LastActivityAt        time.Time        `json:"last_activity_at"`
	IsActive              bool             `json:"is_active"`
	AccessToken           string           `json:"access_token"`
	AccessTokenExpiresAt  time.Time        `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time        `json:"refresh_token_expires_at"`
}
