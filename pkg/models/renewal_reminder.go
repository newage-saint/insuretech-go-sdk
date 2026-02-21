package models

import (
	"time"
)

// RenewalReminder represents a renewal_reminder
type RenewalReminder struct {
	ScheduledAt       time.Time        `json:"scheduled_at"`
	NotificationId    string           `json:"notification_id,omitempty"`
	AuditInfo         interface{}      `json:"audit_info"`
	RenewalScheduleId string           `json:"renewal_schedule_id"`
	SentAt            time.Time        `json:"sent_at,omitempty"`
	Id                string           `json:"id"`
	DaysBeforeRenewal int              `json:"days_before_renewal"`
	Channel           *ReminderChannel `json:"channel"`
	Status            interface{}      `json:"status"`
}
