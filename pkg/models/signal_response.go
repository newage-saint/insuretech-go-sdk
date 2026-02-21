package models

import (
	"time"
)

// SignalResponse represents a signal_response
type SignalResponse struct {
	TrackMuted            *TrackMutedEvent            `json:"track_muted,omitempty"`
	PeerId                string                      `json:"peer_id,omitempty"`
	Timestamp             time.Time                   `json:"timestamp,omitempty"`
	IceCandidateReceived  *ICECandidateReceivedEvent  `json:"ice_candidate_received,omitempty"`
	PeerLeft              *PeerLeftEvent              `json:"peer_left,omitempty"`
	TrackPublished        *TrackPublishedEvent        `json:"track_published,omitempty"`
	OfferReceived         *OfferReceivedEvent         `json:"offer_received,omitempty"`
	PeerStateChanged      *PeerStateChangedEvent      `json:"peer_state_changed,omitempty"`
	TrackUnpublished      *TrackUnpublishedEvent      `json:"track_unpublished,omitempty"`
	RoomId                string                      `json:"room_id,omitempty"`
	AnswerReceived        *AnswerReceivedEvent        `json:"answer_received,omitempty"`
	PeerJoined            *PeerJoinedEvent            `json:"peer_joined,omitempty"`
	Pong                  *PongResponse               `json:"pong,omitempty"`
	RenegotiationRequired *RenegotiationRequiredEvent `json:"renegotiation_required,omitempty"`
	Error                 *ErrorResponse              `json:"error,omitempty"`
}
