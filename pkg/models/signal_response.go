package models

import (
	"time"
)

// SignalResponse represents a signal_response
type SignalResponse struct {
	PeerId                string                      `json:"peer_id,omitempty"`
	PeerJoined            *PeerJoinedEvent            `json:"peer_joined,omitempty"`
	PeerStateChanged      *PeerStateChangedEvent      `json:"peer_state_changed,omitempty"`
	TrackMuted            *TrackMutedEvent            `json:"track_muted,omitempty"`
	Error                 *ErrorResponse              `json:"error,omitempty"`
	Pong                  *PongResponse               `json:"pong,omitempty"`
	Timestamp             time.Time                   `json:"timestamp,omitempty"`
	AnswerReceived        *AnswerReceivedEvent        `json:"answer_received,omitempty"`
	IceCandidateReceived  *ICECandidateReceivedEvent  `json:"ice_candidate_received,omitempty"`
	TrackUnpublished      *TrackUnpublishedEvent      `json:"track_unpublished,omitempty"`
	RenegotiationRequired *RenegotiationRequiredEvent `json:"renegotiation_required,omitempty"`
	PeerLeft              *PeerLeftEvent              `json:"peer_left,omitempty"`
	TrackPublished        *TrackPublishedEvent        `json:"track_published,omitempty"`
	RoomId                string                      `json:"room_id,omitempty"`
	OfferReceived         *OfferReceivedEvent         `json:"offer_received,omitempty"`
}
