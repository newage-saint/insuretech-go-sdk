package models

import (
	"time"
)

// RenewalReminder represents a renewal_reminder
type RenewalReminder struct {
	Id                string           `json:"id"`
	RenewalScheduleId string           `json:"renewal_schedule_id"`
	Status            interface{}      `json:"status"`
	ScheduledAt       time.Time        `json:"scheduled_at"`
	NotificationId    string           `json:"notification_id,omitempty"`
	AuditInfo         *AuditInfo       `json:"audit_info,omitempty"`
	DaysBeforeRenewal int              `json:"days_before_renewal"`
	Channel           *ReminderChannel `json:"channel"`
	SentAt            time.Time        `json:"sent_at,omitempty"`
}
