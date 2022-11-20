package main

import (
	"contact-server/pkg/controller"
	"contact-server/pkg/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title PS Server API documentation
// @version 1.0.0
// @host localhost:8000
// @BasePath /v1/api
func main() {
	db.Init()

	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Error loading .env file: %v \n", envErr.Error())
	}

	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	v1 := router.Group("/v1")
	api := v1.Group("/api")

	//*Welcome
	api.GET("/", controller.Welcome)

	//* Contact Service
	api.GET("/contact/get-all", controller.GetAllContacts)
	api.GET("/contact/:id", controller.GetAllContacts)
	api.POST("/contact/new-contact", controller.CreateContact)
	api.DELETE("contact/:id", controller.DeleteContact)

	log.Println("Server is running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
