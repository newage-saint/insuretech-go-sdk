package models

import (
	"time"
)

// RenewalReminder represents a renewal_reminder
type RenewalReminder struct {
	Channel           *ReminderChannel `json:"channel"`
	ScheduledAt       time.Time        `json:"scheduled_at"`
	SentAt            time.Time        `json:"sent_at,omitempty"`
	AuditInfo         interface{}      `json:"audit_info"`
	Id                string           `json:"id"`
	DaysBeforeRenewal int              `json:"days_before_renewal"`
	Status            interface{}      `json:"status"`
	NotificationId    string           `json:"notification_id,omitempty"`
	RenewalScheduleId string           `json:"renewal_schedule_id"`
}
