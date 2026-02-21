package models

import (
	"time"
)

// User represents a user
type User struct {
	Email                  string      `json:"email,omitempty"`
	PasswordHash           string      `json:"password_hash,omitempty"`
	LoginAttempts          int         `json:"login_attempts,omitempty"`
	Username               string      `json:"username,omitempty"`
	ActivePoliciesCount    int         `json:"active_policies_count,omitempty"`
	PendingClaimsCount     int         `json:"pending_claims_count,omitempty"`
	WalletBalance          *Money      `json:"wallet_balance,omitempty"`
	PreferredAuthMethod    string      `json:"preferred_auth_method,omitempty"`
	UpdatedBy              string      `json:"updated_by,omitempty"`
	NotificationPreference string      `json:"notification_preference,omitempty"`
	UserId                 string      `json:"user_id,omitempty"`
	UpdatedAt              time.Time   `json:"updated_at,omitempty"`
	LastLoginAt            time.Time   `json:"last_login_at,omitempty"`
	CreatedBy              string      `json:"created_by,omitempty"`
	PreferredLanguage      string      `json:"preferred_language,omitempty"`
	MobileNumber           string      `json:"mobile_number,omitempty"`
	Status                 *UserStatus `json:"status,omitempty"`
	CreatedAt              time.Time   `json:"created_at,omitempty"`
	LastLoginSessionType   string      `json:"last_login_session_type,omitempty"`
	LockedUntil            time.Time   `json:"locked_until,omitempty"`
	DeletedAt              time.Time   `json:"deleted_at,omitempty"`
	WalletPaymentMethod    string      `json:"wallet_payment_method,omitempty"`
	BiometricToken         string      `json:"biometric_token,omitempty"`
}
