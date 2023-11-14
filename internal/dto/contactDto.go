package dto

type ContactReq struct {
	FirstName     string
	LastName      string
	Email         string
	ContactNumber string
	Reason        string
	ProjectCode   string
	Message       string
}
type ContactRes struct {
	FirstName     string `json:"firstName,omitempty"`
	LastName      string `json:"lastName,omitempty"`
	Email         string `json:"email,omitempty"`
	ContactNumber string `json:"contactNumber,omitempty"`
	Reason        string `json:"reason,omitempty"`
	ProjectCode   string `json:"projectCode,omitempty"`
	Message       string `json:"message,omitempty"`
}
