package services

import (
	"contact-server/internal/dto"
	"contact-server/internal/models"
	"log"
)

type ContactRepo interface {
	FindAll() (*[]models.Contact, error)
	Find(uint) (*models.Contact, error)
	Edit(models.Contact, dto.ContactReq) (*models.Contact, error)
	Create(models.Contact) (*models.Contact, error)
	Delete(models.Contact) error
}

type ContactService struct {
	contactRepo ContactRepo
}

func NewContactService(r ContactRepo) *ContactService {
	return &ContactService{
		contactRepo: r,
	}
}

func (s *ContactService) GetAllContacts() (*[]dto.ContactRes, error) {

	contacts, err := s.contactRepo.FindAll()
	if err != nil {
		log.Println("Contact Service - Could not retrieve all contacts")
		return nil, err
	}

	var resArr []dto.ContactRes

	for _, contact := range *contacts {
		res := dto.ContactRes{
			FirstName:     contact.FirstName,
			LastName:      contact.LastName,
			Email:         contact.Email,
			ContactNumber: contact.ContactNumber,
			Reason:        contact.Reason,
			Message:       contact.Message,
		}

		resArr = append(resArr, res)
	}

	return &resArr, nil
}

func (s *ContactService) GetContact(contactId uint) (*dto.ContactRes, error) {

	contact, err := s.contactRepo.Find(contactId)
	if err != nil {
		log.Println("Contact Service - Could not retrieve contact")
		return nil, err
	}

	res := dto.ContactRes{
		FirstName:     contact.FirstName,
		LastName:      contact.LastName,
		Email:         contact.Email,
		ContactNumber: contact.ContactNumber,
		Reason:        contact.Reason,
		Message:       contact.Message,
	}
	return &res, nil
}

func (s *ContactService) EditContact(contactId uint, req dto.ContactReq) (*dto.ContactRes, error) {
	return nil, nil
}

func (s *ContactService) CreateContact(req dto.ContactReq) (*dto.ContactRes, error) {
	return nil, nil
}

func (s *ContactService) DeleteContact(contactId uint) error {
	return nil
}
