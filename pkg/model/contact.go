package model

type Contact struct {
	Id            int    `json:"id" gorm:"primaryKey" `
	FirstName     string `json:"firstName" `
	LastName      string `json:"lastName" `
	Email         string `json:"email" `
	ContactNumber string `json:"contactNumber" `
	Reason        string `json:"reason" `
	ProjectCode   string `json:"projectCode" `
	Message       string `json:"message" `
}

func (u *Contact) TableName() string {
	// custom table name, this is default
	return "contact"
}
