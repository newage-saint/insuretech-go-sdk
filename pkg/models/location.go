package models

// Location represents a location
type Location struct {
	Longitude float64 `json:"longitude,omitempty"`
	Altitude  float64 `json:"altitude,omitempty"`
	Accuracy  float64 `json:"accuracy,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
}
