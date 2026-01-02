package models

// ContactInfo represents a contact_info
type ContactInfo struct {
	MobileNumber    string `json:"mobile_number,omitempty"`
	Email           string `json:"email,omitempty"`
	AlternateMobile string `json:"alternate_mobile,omitempty"`
	Landline        string `json:"landline,omitempty"`
}
