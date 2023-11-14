package main

import (
	"contact-server/internal/repo"
	"contact-server/internal/server"
	"contact-server/internal/services"
	"fmt"
	"log"
)

func Run() error {
	db, err := repo.NewContactRepo()
	if err != nil {
		log.Println("Could not connect to db")
		return err
	}

	contactService := services.NewContactService(db)
	contactController := server.New(contactService)

	contactController.Serve()

	return nil
}

func main() {

	fmt.Println("INFO: starting contact server")

	if err := Run(); err != nil {
		log.Println(err)
	}

}
