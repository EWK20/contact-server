package repo

import (
	"contact-server/internal/dto"
	"contact-server/internal/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// type ContactRepo interface {
// 	FindAll() (*[]models.Contact, error)
// 	Find(contactId uint) (*models.Contact, error)
// 	Edit(contact models.Contact) (*models.Contact, error)
// 	Create(contact models.Contact) (*models.Contact, error)
// 	Delete(contact models.Contact) error
// }

type ContactDB struct {
	conn *gorm.DB
}

func NewContactRepo() (*ContactDB, error) {

	dsn := os.Getenv("DEVDB")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println("Contact Repo - Could not connect to db")
		return nil, err
	}
	db.AutoMigrate(&models.Contact{})

	log.Println("Database connection is successful")

	return &ContactDB{
		conn: db,
	}, nil
}

func (db *ContactDB) FindAll() (*[]models.Contact, error) {
	var contacts []models.Contact

	if err := db.conn.Find(&contacts).Error; err != nil {
		log.Println("Contact Repo - Could not retrieve all contacts")
		return nil, err
	}

	return &contacts, nil
}

func (db *ContactDB) Find(contactId uint) (*models.Contact, error) {
	var contact models.Contact

	if err := db.conn.First(&contact, contactId).Error; err != nil {
		log.Println("Contact Repo - Could not retrieve contacts")
		return nil, err
	}

	return &contact, nil
}

func (db *ContactDB) Edit(contact models.Contact, req dto.ContactReq) (*models.Contact, error) {

	if err := db.conn.Model(&contact).Updates(models.Contact{
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Email:         req.Email,
		ContactNumber: req.ContactNumber,
		Reason:        req.Reason,
		Message:       req.Message,
	}).Error; err != nil {
		log.Println("Contact Repo - Could not update contact")
		return nil, err
	}

	return &contact, nil

}

func (db *ContactDB) Create(contact models.Contact) (*models.Contact, error) {

	if err := db.conn.Create(&contact).Error; err != nil {
		log.Println("Contact Repo - Could not create contact")
		return nil, err
	}

	return &contact, nil
}

func (db *ContactDB) Delete(contact models.Contact) error {

	if err := db.conn.Delete(&contact).Error; err != nil {
		log.Println("Contact Repo - Could not delete contact")
		return nil
	}

	return nil

}
