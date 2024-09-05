package handlers

import (
	"net/http"

	"loan-service/databases"
	"loan-service/models"

	"github.com/gin-gonic/gin"
)

// CreateLoan creates a new loan
func CreateLoan(c *gin.Context) {
	var loan models.Loan
	if err := c.ShouldBindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if loan.PrincipalAmount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Principal amount must be greater than zero"})
		return
	}

	loan.State = "proposed"
	if err := databases.DB.Create(&loan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create loan"})
		return
	}
	c.JSON(http.StatusOK, loan)
}

// ApproveLoan approves a loan
func ApproveLoan(c *gin.Context) {
	var loan models.Loan
	loanID := c.Param("id")

	if err := databases.DB.Where("id = ?", loanID).First(&loan).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	if loan.State != models.StateProposed {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Loan is not in proposed state"})
		return
	}

	var approvalData models.ApprovalData

	if err := c.ShouldBindJSON(&approvalData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loan.ApprovalData = approvalData
	loan.State = models.StateApproved

	if err := databases.DB.Save(&loan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update loan"})
		return
	}

	c.JSON(http.StatusOK, loan)
}

func RecordInvestment(c *gin.Context) {
	var loan models.Loan
	loanID := c.Param("id")

	// Fetch the loan and its investments in one query
	if err := databases.DB.Preload("Investments").Preload("DisbursementData").Preload("ApprovalData").
		Where("id = ?", loanID).First(&loan).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	if loan.State != models.StateApproved && loan.State != models.StateInvested {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Loan is not in a state that allows investment"})
		return
	}

	var investment models.Investment
	if err := c.ShouldBindJSON(&investment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the LoanID for the new investment
	investment.LoanID = loan.ID

	// Calculate total investment including the new one
	totalInvested := 0.0
	for _, inv := range loan.Investments {
		totalInvested += inv.Amount
	}
	totalInvested += investment.Amount

	if totalInvested > loan.PrincipalAmount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Total investment exceeds loan principal"})
		return
	}

	// Save the new investment
	if err := databases.DB.Create(&investment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record investment"})
		return
	}

	// Update the loan state if fully invested
	if totalInvested == loan.PrincipalAmount {
		loan.State = models.StateInvested
	}

	// Save the updated loan state
	if err := databases.DB.Save(&loan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update loan state"})
		return
	}

	// Return the updated loan with all related data
	c.JSON(http.StatusOK, loan)
}

func DisburseLoan(c *gin.Context) {
	var loan models.Loan
	loanID := c.Param("id")

	if err := databases.DB.Preload("ApprovalData").Preload("Investments").Preload("DisbursementData").Where("id = ?", loanID).First(&loan).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	if loan.State != models.StateInvested {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Loan is not in invested state"})
		return
	}

	var disbursementData models.DisbursementData
	if err := c.ShouldBindJSON(&disbursementData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if loan.DisbursementData.ID == 0 {
		disbursementData.LoanID = loan.ID
		if err := databases.DB.Create(&disbursementData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save disbursement data"})
			return
		}
	} else {
		disbursementData.ID = loan.DisbursementData.ID
		if err := databases.DB.Save(&disbursementData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update disbursement data"})
			return
		}
	}

	loan.State = models.StateDisbursed
	databases.DB.Save(&loan)

	c.JSON(http.StatusOK, loan)
}

func GetLoan(c *gin.Context) {
	var loan models.Loan
	loanID := c.Param("id")

	if err := databases.DB.Preload("ApprovalData").Preload("Investments").Preload("DisbursementData").Where("id = ?", loanID).First(&loan).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	c.JSON(http.StatusOK, loan)
}

func ListLoans(c *gin.Context) {
	var loans []models.Loan
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")

	if err := databases.DB.Preload("ApprovalData").Preload("Investments").Preload("DisbursementData").
		Limit(limit).Offset(offset).Find(&loans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch loans"})
		return
	}

	c.JSON(http.StatusOK, loans)
}
