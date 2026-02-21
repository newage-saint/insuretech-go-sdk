package models

// InsurerConfig represents a insurer_config
type InsurerConfig struct {
	ApiBaseUrl                string              `json:"api_base_url,omitempty"`
	AutoUnderwritingEnabled   bool                `json:"auto_underwriting_enabled,omitempty"`
	RealTimeClaimNotification bool                `json:"real_time_claim_notification,omitempty"`
	PaymentTerms              string              `json:"payment_terms,omitempty"`
	AuthType                  *AuthenticationType `json:"auth_type,omitempty"`
	WebhookUrl                string              `json:"webhook_url,omitempty"`
	BusinessModel             string              `json:"business_model,omitempty"`
	AuditInfo                 interface{}         `json:"audit_info"`
	WebhookSecret             string              `json:"webhook_secret,omitempty"`
	ClaimSettlementDays       int                 `json:"claim_settlement_days,omitempty"`
	Id                        string              `json:"id"`
	InsurerId                 string              `json:"insurer_id"`
	ApiVersion                string              `json:"api_version,omitempty"`
	AuthCredentials           string              `json:"auth_credentials,omitempty"`
	UnderwritingThreshold     int                 `json:"underwriting_threshold,omitempty"`
}
