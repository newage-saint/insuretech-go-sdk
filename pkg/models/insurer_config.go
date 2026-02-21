package models

// InsurerConfig represents a insurer_config
type InsurerConfig struct {
	ApiVersion                string              `json:"api_version,omitempty"`
	AuthType                  *AuthenticationType `json:"auth_type,omitempty"`
	AuthCredentials           string              `json:"auth_credentials,omitempty"`
	WebhookUrl                string              `json:"webhook_url,omitempty"`
	AutoUnderwritingEnabled   bool                `json:"auto_underwriting_enabled,omitempty"`
	RealTimeClaimNotification bool                `json:"real_time_claim_notification,omitempty"`
	WebhookSecret             string              `json:"webhook_secret,omitempty"`
	ClaimSettlementDays       int                 `json:"claim_settlement_days,omitempty"`
	PaymentTerms              string              `json:"payment_terms,omitempty"`
	AuditInfo                 interface{}         `json:"audit_info"`
	Id                        string              `json:"id"`
	BusinessModel             string              `json:"business_model,omitempty"`
	UnderwritingThreshold     int                 `json:"underwriting_threshold,omitempty"`
	InsurerId                 string              `json:"insurer_id"`
	ApiBaseUrl                string              `json:"api_base_url,omitempty"`
}
