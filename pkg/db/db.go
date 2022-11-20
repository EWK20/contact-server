package db

import (
	"log"
	"os"

	"contact-server/pkg/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Init() {

	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Error loading .env file: %v \n", envErr.Error())
	}

	dsn := os.Getenv("PRODDB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Could not connect to db")
	}

	log.Println("Database connection is successful")

	db.AutoMigrate(&model.Contact{})
	DB = db

}
