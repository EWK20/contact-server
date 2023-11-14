package server

import (
	"contact-server/internal/dto"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContactService interface {
	GetAllContacts() (*[]dto.ContactRes, error)
	GetContact(contactId uint) (*dto.ContactRes, error)
	EditContact(contactId uint, req dto.ContactReq) (*dto.ContactRes, error)
	CreateContact(req dto.ContactReq) (*dto.ContactRes, error)
	DeleteContact(contactId uint) error
}

type Server struct {
	contactService ContactService
}

func New(service ContactService) *Server {
	return &Server{
		contactService: service,
	}
}

func (s *Server) Serve() {

	router := gin.Default()
	api := router.Group("/v1/api")

	//! Put all logic in controller layer

	//* Welcome
	api.GET("/", func(ctx *gin.Context) {
		log.Println("Web Server Running...")
		ctx.JSON(http.StatusOK, "Web Server Running")
	})

	//* Contact Service
	api.GET("/contact/get-all", func(ctx *gin.Context) {
		contacts, err := s.contactService.GetAllContacts()
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		ctx.JSON(http.StatusOK, contacts)
	})

	api.GET("/contact", func(ctx *gin.Context) {
		contactId, err := strconv.Atoi(ctx.Query("id"))
		if err != nil {
			log.Println("Could not convert query param to integer")
			ctx.AbortWithStatus(http.StatusBadRequest)
		}

		res, err := s.contactService.GetContact(uint(contactId))
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		ctx.JSON(http.StatusOK, res)
	})

	api.PATCH("/contact/edit-contact", func(ctx *gin.Context) {
		contactId, err := strconv.Atoi(ctx.Query("id"))
		if err != nil {
			log.Println("Could not convert query param to integer")
			ctx.AbortWithStatus(http.StatusBadRequest)
		}

		var req dto.ContactReq

		err = ctx.Bind(&req)
		if err != nil {
			log.Println("Checkin Controller - Could not retrieve guest data")
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		res, err := s.contactService.EditContact(uint(contactId), req)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		ctx.JSON(http.StatusOK, res)
	})

	api.POST("/contact/create-contact", func(ctx *gin.Context) {

		var req dto.ContactReq

		err := ctx.Bind(&req)
		if err != nil {
			log.Println("Checkin Controller - Could not retrieve guest data")
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		res, err := s.contactService.CreateContact(req)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		ctx.JSON(http.StatusOK, res)
	})

	api.DELETE("/contact", func(ctx *gin.Context) {
		contactId, err := strconv.Atoi(ctx.Query("id"))
		if err != nil {
			log.Println("Could not convert query param to integer")
			ctx.AbortWithStatus(http.StatusBadRequest)
		}

		err = s.contactService.DeleteContact(uint(contactId))
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		ctx.Status(http.StatusOK)
	})

	err := router.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
