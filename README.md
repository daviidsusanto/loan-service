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
```
git clone https://github.com/daviidsusanto/loan-service.git
cd loan-service
```
2. Set Up Environment Variables
Create a .env file in the root directory of the project with the following content:
```
POSTGRES_USER=youruser
POSTGRES_PASSWORD=yourpassword
POSTGRES_DB=yourdb

DB_HOST=postgres
DB_PORT=5432
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
```
3. Running Docker
To build and run the services, use:
```
docker-compose up --build
```
To run in detached mode, use:
```
docker-compose up -d --build
```
To stop and remove containers, networks, and volumes, use:
```
docker-compose down
```
4. Accessing API
To access API Endpoint you can access URL : http://localhost:8080/

## **Testing**
To run tests, use:
```
go test ./...
```

## **API Endpoints**
**Create Loan**
- Endpoint: POST /v1/loans
- Request Body:
```
{
  "borrower_id": "user-123",
  "principal_amount": 10000.0,
  "rate": 5.0,
  "roi": 10.0,
  "aggrement_letter": "www.example.com/aggrement-letter.pdf"
}
```
**Approve Loan**
- Endpoint: PUT /v1/loans/:id/approve
- Request Body:
```
{
  "field_validator_id": "validator-001",
  "approval_date": "2024-09-03",
  "proof_of_visit": "https://example.com/proof-of-visit.jpg"
}
```
**Record Investment**
- Endpoint: PUT /v1/loans/:id/invest
- Request Body:
```
{
  "investor": "investor-001",
  "amount": 10000
}
```
**Disburse Loan**
- Endpoint: PUT /v1/loans/:id/disburse
- Request Body:
```
{
  "agreement_letter": "https://example.com/signed-agreement.pdf",
  "field_officer_id": "field-officer-001",
  "disbursement_date": "2024-09-03"
}
```
**Get Loan By ID**
- Endpoint: GET /v1/loans/:id

**List Loans**
- Endpoint: GET /v1/loans
- Query Parameter:
    * limit: Number of loans to return (default: 10)
    * offset: Pagination offset (default: 0)