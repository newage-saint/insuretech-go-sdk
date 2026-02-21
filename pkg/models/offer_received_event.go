package models

import (
	"time"
)

// OfferReceivedEvent represents a offer_received_event
type OfferReceivedEvent struct {
	RoomId     string    `json:"room_id,omitempty"`
	Offer      *SDPOffer `json:"offer,omitempty"`
	ReceivedAt time.Time `json:"received_at,omitempty"`
}
