package databases

import (
	"log"

	"loan-service/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DBTest *gorm.DB

// ConnectDBTest initializes the test database connection
func ConnectDBTest() {
	var err error

	// Connect to an in-memory SQLite database
	DBTest, err = gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("failed to connect to the test database: %v", err)
	}

	// Automatically migrate the schema for your models
	DBTest.AutoMigrate(&models.Loan{}, &models.ApprovalData{}, &models.Investment{}, &models.DisbursementData{})
}

// CleanUpTestData clears the test database after each test
func CleanUpTestData() {
	DBTest.DropTableIfExists(&models.Loan{}, &models.ApprovalData{}, &models.Investment{}, &models.DisbursementData{})
}
