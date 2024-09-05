package databases

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"loan-service/models"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	// Load environment variables from .env file
	err = godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Parse the database port
	p := os.Getenv("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Println("Invalid port number")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)

	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect to the database")
	}

	// Run auto-migration to create/update tables based on the models
	DB.AutoMigrate(
		&models.Investment{},
		&models.ApprovalData{},
		&models.Loan{},
		&models.DisbursementData{},
	)

	fmt.Println("Connection opened to the database")
}
