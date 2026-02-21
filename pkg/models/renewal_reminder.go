package models

import (
	"time"
)

// RenewalReminder represents a renewal_reminder
type RenewalReminder struct {
	Channel           *ReminderChannel `json:"channel"`
	ScheduledAt       time.Time        `json:"scheduled_at"`
	NotificationId    string           `json:"notification_id,omitempty"`
	AuditInfo         interface{}      `json:"audit_info"`
	Id                string           `json:"id"`
	Status            interface{}      `json:"status"`
	SentAt            time.Time        `json:"sent_at,omitempty"`
	RenewalScheduleId string           `json:"renewal_schedule_id"`
	DaysBeforeRenewal int              `json:"days_before_renewal"`
}
