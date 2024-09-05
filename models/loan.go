package models

import "time"

type LoanState string

const (
	StateProposed  LoanState = "proposed"
	StateApproved  LoanState = "approved"
	StateInvested  LoanState = "invested"
	StateDisbursed LoanState = "disbursed"
)

type Loan struct {
	ID               uint             `gorm:"primary_key" json:"id"`
	BorrowerID       string           `json:"borrower_id"`
	PrincipalAmount  float64          `json:"principal_amount"`
	Rate             float64          `json:"rate"`
	ROI              float64          `json:"roi"`
	AgreementLetter  string           `json:"agreement_letter"`
	State            LoanState        `json:"state"`
	ApprovalData     ApprovalData     `json:"approval_data" gorm:"foreignKey:LoanID"`
	Investments      []Investment     `json:"investments" gorm:"foreignKey:LoanID"`
	DisbursementData DisbursementData `json:"disbursement_data" gorm:"foreignKey:LoanID"`
	CreatedAt        time.Time        `json:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at"`
}

type ApprovalData struct {
	ID               uint   `gorm:"primaryKey"`
	LoanID           uint   `json:"loan_id"`
	FieldValidatorID string `json:"field_validator_id"`
	ApprovalDate     string `json:"approval_date"`
	ProofOfVisit     string `json:"proof_of_visit"`
}

type Investment struct {
	ID       uint    `gorm:"primaryKey"`
	LoanID   uint    `json:"loan_id"`
	Investor string  `json:"investor"`
	Amount   float64 `json:"amount"`
}

type DisbursementData struct {
	ID               uint   `gorm:"primaryKey"`
	LoanID           uint   `json:"loan_id"`
	FieldOfficerID   string `json:"field_officer_id"`
	DisbursementDate string `json:"disbursement_date"`
	AgreementLetter  string `json:"agreement_letter"`
}
