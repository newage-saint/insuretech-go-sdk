package models

// AnswerSendingRequest represents a answer_sending_request
type AnswerSendingRequest struct {
	ToPeerId   string `json:"to_peer_id"`
	Sdp        string `json:"sdp,omitempty"`
	FromPeerId string `json:"from_peer_id"`
}
