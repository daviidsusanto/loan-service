package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"loan-service/databases"
	"loan-service/models"

	"github.com/gin-gonic/gin"
)

func mockDatabase() {
	databases.ConnectDBTest()
	databases.DB = databases.DBTest
}

func TestCreateLoan(t *testing.T) {
	mockDatabase()
	defer databases.CleanUpTestData()

	router := gin.Default()
	router.POST("/v1/loans", CreateLoan)

	loan := models.Loan{
		BorrowerID:      "123",
		PrincipalAmount: 1000.0,
		Rate:            5.0,
		ROI:             10.0,
		AgreementLetter: "www.example.com/Sample_Agreement.pdf",
	}

	loanJSON, err := json.Marshal(loan)
	if err != nil {
		t.Fatalf("Error marshalling loan data: %v", err)
	}

	req, err := http.NewRequest("POST", "/v1/loans", bytes.NewBuffer(loanJSON))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", status)
	}

	var createdLoan models.Loan
	if err := json.Unmarshal(rr.Body.Bytes(), &createdLoan); err != nil {
		t.Fatalf("Error unmarshalling response: %v", err)
	}

	if createdLoan.State != models.StateProposed {
		t.Errorf("Expected loan state 'proposed', got %v", createdLoan.State)
	}

	if createdLoan.BorrowerID != "123" {
		t.Errorf("Expected borrower ID '123', got %v", createdLoan.BorrowerID)
	}
}
