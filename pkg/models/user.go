package models

import (
	"time"
)

// User represents a user
type User struct {
	LoginAttempts          int         `json:"login_attempts,omitempty"`
	LockedUntil            time.Time   `json:"locked_until,omitempty"`
	Username               string      `json:"username,omitempty"`
	PreferredLanguage      string      `json:"preferred_language,omitempty"`
	UserId                 string      `json:"user_id,omitempty"`
	Email                  string      `json:"email,omitempty"`
	Status                 *UserStatus `json:"status,omitempty"`
	UpdatedAt              time.Time   `json:"updated_at,omitempty"`
	BiometricToken         string      `json:"biometric_token,omitempty"`
	PendingClaimsCount     int         `json:"pending_claims_count,omitempty"`
	LastLoginSessionType   string      `json:"last_login_session_type,omitempty"`
	NotificationPreference string      `json:"notification_preference,omitempty"`
	WalletPaymentMethod    string      `json:"wallet_payment_method,omitempty"`
	LastLoginAt            time.Time   `json:"last_login_at,omitempty"`
	PreferredAuthMethod    string      `json:"preferred_auth_method,omitempty"`
	UpdatedBy              string      `json:"updated_by,omitempty"`
	DeletedAt              time.Time   `json:"deleted_at,omitempty"`
	ActivePoliciesCount    int         `json:"active_policies_count,omitempty"`
	WalletBalance          *Money      `json:"wallet_balance,omitempty"`
	MobileNumber           string      `json:"mobile_number,omitempty"`
	PasswordHash           string      `json:"password_hash,omitempty"`
	CreatedAt              time.Time   `json:"created_at,omitempty"`
	CreatedBy              string      `json:"created_by,omitempty"`
}
