package models

import (
	"time"
)

// User represents a user
type User struct {
	NotificationPreference string      `json:"notification_preference,omitempty"`
	WalletPaymentMethod    string      `json:"wallet_payment_method,omitempty"`
	ActivePoliciesCount    int         `json:"active_policies_count,omitempty"`
	Email                  string      `json:"email,omitempty"`
	PasswordHash           string      `json:"password_hash,omitempty"`
	CreatedAt              time.Time   `json:"created_at,omitempty"`
	LoginAttempts          int         `json:"login_attempts,omitempty"`
	LockedUntil            time.Time   `json:"locked_until,omitempty"`
	BiometricToken         string      `json:"biometric_token,omitempty"`
	UpdatedBy              string      `json:"updated_by,omitempty"`
	UserId                 string      `json:"user_id,omitempty"`
	Status                 *UserStatus `json:"status,omitempty"`
	UpdatedAt              time.Time   `json:"updated_at,omitempty"`
	LastLoginSessionType   string      `json:"last_login_session_type,omitempty"`
	PreferredAuthMethod    string      `json:"preferred_auth_method,omitempty"`
	DeletedAt              time.Time   `json:"deleted_at,omitempty"`
	WalletBalance          *Money      `json:"wallet_balance,omitempty"`
	Username               string      `json:"username,omitempty"`
	PendingClaimsCount     int         `json:"pending_claims_count,omitempty"`
	MobileNumber           string      `json:"mobile_number,omitempty"`
	LastLoginAt            time.Time   `json:"last_login_at,omitempty"`
	CreatedBy              string      `json:"created_by,omitempty"`
	PreferredLanguage      string      `json:"preferred_language,omitempty"`
}
