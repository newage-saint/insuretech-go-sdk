package models

// FieldSorting represents a field_sorting
type FieldSorting struct {
	Field     string         `json:"field,omitempty"`
	Direction *SortDirection `json:"direction,omitempty"`
}
