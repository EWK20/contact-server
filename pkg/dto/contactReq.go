package dto

type ContactReq struct {
	FirstName     string `json:"firstName" `
	LastName      string `json:"lastName" `
	Email         string `json:"email" `
	ContactNumber string `json:"contactNumber" `
	Reason        string `json:"reason" `
	ProjectCode   string `json:"projectCode" `
	Message       string `json:"message" `
}
