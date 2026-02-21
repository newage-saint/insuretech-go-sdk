package models

import (
	"time"
)

// SignalResponse represents a signal_response
type SignalResponse struct {
	TrackMuted            *TrackMutedEvent            `json:"track_muted,omitempty"`
	RoomId                string                      `json:"room_id,omitempty"`
	Timestamp             time.Time                   `json:"timestamp,omitempty"`
	RenegotiationRequired *RenegotiationRequiredEvent `json:"renegotiation_required,omitempty"`
	Pong                  *PongResponse               `json:"pong,omitempty"`
	OfferReceived         *OfferReceivedEvent         `json:"offer_received,omitempty"`
	PeerStateChanged      *PeerStateChangedEvent      `json:"peer_state_changed,omitempty"`
	Error                 *ErrorResponse              `json:"error,omitempty"`
	PeerId                string                      `json:"peer_id,omitempty"`
	AnswerReceived        *AnswerReceivedEvent        `json:"answer_received,omitempty"`
	IceCandidateReceived  *ICECandidateReceivedEvent  `json:"ice_candidate_received,omitempty"`
	TrackPublished        *TrackPublishedEvent        `json:"track_published,omitempty"`
	TrackUnpublished      *TrackUnpublishedEvent      `json:"track_unpublished,omitempty"`
	PeerJoined            *PeerJoinedEvent            `json:"peer_joined,omitempty"`
	PeerLeft              *PeerLeftEvent              `json:"peer_left,omitempty"`
}
