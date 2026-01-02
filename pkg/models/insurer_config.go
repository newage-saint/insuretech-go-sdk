package models

// InsurerConfig represents a insurer_config
type InsurerConfig struct {
	AuthCredentials           string              `json:"auth_credentials,omitempty"`
	UnderwritingThreshold     int                 `json:"underwriting_threshold,omitempty"`
	AuditInfo                 *AuditInfo          `json:"audit_info,omitempty"`
	Id                        string              `json:"id"`
	AuthType                  *AuthenticationType `json:"auth_type,omitempty"`
	BusinessModel             string              `json:"business_model,omitempty"`
	RealTimeClaimNotification bool                `json:"real_time_claim_notification,omitempty"`
	InsurerId                 string              `json:"insurer_id"`
	ApiVersion                string              `json:"api_version,omitempty"`
	PaymentTerms              string              `json:"payment_terms,omitempty"`
	WebhookUrl                string              `json:"webhook_url,omitempty"`
	WebhookSecret             string              `json:"webhook_secret,omitempty"`
	AutoUnderwritingEnabled   bool                `json:"auto_underwriting_enabled,omitempty"`
	ClaimSettlementDays       int                 `json:"claim_settlement_days,omitempty"`
	ApiBaseUrl                string              `json:"api_base_url,omitempty"`
}
