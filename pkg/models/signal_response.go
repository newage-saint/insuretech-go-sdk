package models

import (
	"time"
)

// SignalResponse represents a signal_response
type SignalResponse struct {
	PeerId                string                      `json:"peer_id,omitempty"`
	RoomId                string                      `json:"room_id,omitempty"`
	PeerLeft              *PeerLeftEvent              `json:"peer_left,omitempty"`
	Pong                  *PongResponse               `json:"pong,omitempty"`
	Timestamp             time.Time                   `json:"timestamp,omitempty"`
	IceCandidateReceived  *ICECandidateReceivedEvent  `json:"ice_candidate_received,omitempty"`
	RenegotiationRequired *RenegotiationRequiredEvent `json:"renegotiation_required,omitempty"`
	PeerStateChanged      *PeerStateChangedEvent      `json:"peer_state_changed,omitempty"`
	TrackPublished        *TrackPublishedEvent        `json:"track_published,omitempty"`
	TrackUnpublished      *TrackUnpublishedEvent      `json:"track_unpublished,omitempty"`
	TrackMuted            *TrackMutedEvent            `json:"track_muted,omitempty"`
	Error                 *ErrorResponse              `json:"error,omitempty"`
	OfferReceived         *OfferReceivedEvent         `json:"offer_received,omitempty"`
	AnswerReceived        *AnswerReceivedEvent        `json:"answer_received,omitempty"`
	PeerJoined            *PeerJoinedEvent            `json:"peer_joined,omitempty"`
}
