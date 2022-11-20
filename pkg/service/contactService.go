package service

import (
	"contact-server/pkg/db"
	"contact-server/pkg/dto"
	"contact-server/pkg/model"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func sendContactEmail(contact dto.ContactReq) error {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Error loading .env file: %v \n", envErr.Error())
	}

	from := os.Getenv("FROM")
	pass := os.Getenv("PASS")

	to := os.Getenv("TO")
	toArr := []string{to}

	host := os.Getenv("HOST")
	port := os.Getenv("EMAILPORT")
	address := host + ":" + port

	identity := os.Getenv("IDENTITY")
	subject := "Subject: NEW CONTACT REQUEST!!\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := "<html><body> <h3>Contact Details:</h3>\n<strong>Full Name:</strong> " + contact.FirstName + " " + contact.LastName + "<br/>" +
		"<strong>Email Address:</strong> " + contact.Email + "<br/>" +
		"<strong>Contact Number:</strong> " + contact.ContactNumber + "<br/>" +
		"<strong>Reason For Contact:</strong> " + contact.Reason + "<br/>" +
		"<strong>Project Code:</strong> " + contact.ProjectCode + "<br/>" +
		"<strong>Message:</strong> " + contact.Message + "<br/></body></html>"
	message := []byte(identity + subject + mime + body)

	auth := smtp.PlainAuth("", from, pass, host)

	err := smtp.SendMail(address, auth, from, toArr, message)
	if err != nil {
		return err
	}
	return nil

}

func GetAllContacts() ([]dto.ContactRes, error) {
	var contacts []model.Contact
	var contactArr []dto.ContactRes

	var contactRes dto.ContactRes

	if err := db.DB.Find(&contacts).Error; err != nil {
		log.Println("Get All Service - Could not retrieve contacts")
		return contactArr, err
	}

	for i := 0; i < len(contacts); i++ {
		contactRes.FirstName = contacts[i].FirstName
		contactRes.LastName = contacts[i].LastName
		contactRes.Email = contacts[i].Email
		contactRes.ContactNumber = contacts[i].ContactNumber
		contactRes.Reason = contacts[i].Reason
		contactRes.ProjectCode = contacts[i].ProjectCode
		contactRes.Message = contacts[i].Message

		contactArr = append(contactArr, contactRes)
	}

	return contactArr, nil
}

func GetContact(id int) (dto.ContactRes, error) {
	var contact model.Contact
	var contactRes dto.ContactRes

	if err := db.DB.Find(&contact, id).Error; err != nil {
		log.Println("Get By Id Service - Could not retrieve contacts")
		return contactRes, err
	}

	contactRes.FirstName = contact.FirstName
	contactRes.LastName = contact.LastName
	contactRes.Email = contact.Email
	contactRes.ContactNumber = contact.ContactNumber
	contactRes.Reason = contact.Reason
	contactRes.ProjectCode = contact.ProjectCode
	contactRes.Message = contact.Message

	return contactRes, nil
}

func CreateContact(contactReq dto.ContactReq) (dto.ContactRes, error) {

	var contact model.Contact
	var contactRes dto.ContactRes

	contact.FirstName = contactReq.FirstName
	contact.LastName = contactReq.LastName
	contact.Email = contactReq.Email
	contact.ContactNumber = contactReq.ContactNumber
	contact.Reason = contactReq.Reason
	contact.ProjectCode = contactReq.ProjectCode
	contact.Message = contactReq.Message

	if err := db.DB.Create(&contact).Error; err != nil {
		log.Println("Create Service - Could not create contact")
		return contactRes, err
	}

	sendContactEmail(contactReq)
	contactRes.FirstName = contact.FirstName

	return contactRes, nil
}

func DeleteContact(id int) error {
	var contact model.Contact

	if err := db.DB.Delete(contact, id).Error; err != nil {
		log.Fatalln("Delete Service - Could not delete contact")
		return err
	}
	return nil
}
