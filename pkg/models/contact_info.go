package models

// ContactInfo represents a contact_info
type ContactInfo struct {
	Email           string `json:"email,omitempty"`
	AlternateMobile string `json:"alternate_mobile,omitempty"`
	Landline        string `json:"landline,omitempty"`
	MobileNumber    string `json:"mobile_number,omitempty"`
}
