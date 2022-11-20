package controller

import (
	"contact-server/pkg/dto"
	"contact-server/pkg/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func enableCors(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}
}

func Welcome(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{"Msg": "Hello World"})
}

// GetContacts... Get all users
// @Summary Get all contacts
// @Description get all contacts
// @Tags ContactService
// @Success 302 {array} dto.ContactRes
// @Failure 500 {object} object
// @Router / [get]
func GetAllContacts(ctx *gin.Context) {

	contacts, err := service.GetAllContacts()
	if err != nil {
		log.Fatalln("Get All Controller - Could not retrieve contacts")
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	enableCors(ctx)

	ctx.IndentedJSON(http.StatusFound, contacts)
}

// GetContact... Get contact
// @Summary Get contact
// @Description get contact
// @Tags ContactService
// @Success 302 {object} dto.ContactRes
// @Failure 404,500 {object} object
// @Router /{id} [get]
func GetContact(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		log.Fatalln("Could not convert query param to integer")
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	contact, err := service.GetContact(id)
	if err != nil {
		log.Fatalln("Get By Id Controller - Could not retrieve contact")
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	enableCors(ctx)

	ctx.IndentedJSON(http.StatusFound, contact)
}

// CreateContact... Create contact
// @Summary Create contact
// @Description create contact
// @Tags ContactService
// @Success 201 {object} dto.ContactRes
// @Failure 400,500 {object} object
// @Router / [post]
func CreateContact(ctx *gin.Context) {

	var contactReq dto.ContactReq

	err := ctx.BindJSON(&contactReq)
	if err != nil {
		log.Fatalln("Could not retrieve contact data")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	newContact, err := service.CreateContact(contactReq)
	if err != nil {
		log.Fatalln("Create Controller - Failed to notify about new contact")
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	enableCors(ctx)

	ctx.IndentedJSON(http.StatusCreated, newContact)
}

// DeleteContact... Delete contact
// @Summary Delete contact
// @Description delete contact
// @Tags ContactService
// @Success 200  {object} object
// @Failure 500 {object} object
// @Router /{id} [delete]
func DeleteContact(ctx *gin.Context) {
	id, idErr := strconv.Atoi(ctx.Query("id"))
	if idErr != nil {
		log.Fatalln("Could not convert query param to integer")
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": idErr.Error()})
	}

	err := service.DeleteContact(id)
	if err != nil {
		log.Fatalln("Delete Controller - Could not delete contact")
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	enableCors(ctx)

	ctx.IndentedJSON(http.StatusOK, gin.H{"Msg": "Deleted"})
}

//! REDEPLOY - 15/11/22 @ 21:22
