## **Overview**
This is a Loan Service built using the Go Gin framework, PostgreSQL, and GORM. It provides APIs to manage loans through various states: proposed, approved, invested, and disbursed.

## **Features**
1. **Create a Loan:** Endpoint to create a new loan.
2. **Approve a Loan:** Endpoint to approve a loan.
3. **Record Investment:** Endpoint to record investments in a loan.
4. **Disburse a Loan:** Endpoint to disburse a loan.
5. **Get Loan:** Endpoint to retrieve loan details by ID.
6. **List Loans:** Endpoint to list loans with pagination.

## **Requirements**
- Go (1.18 or higher)
- PostgreSQL 15
- Docker (optional, for database container)

## **Setup**
1. Clone the Repository
```git clone https://github.com/daviidsusanto/loan-service.git
cd loan-service`