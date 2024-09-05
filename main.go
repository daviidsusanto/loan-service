package main

import (
	"loan-service/databases"
	"loan-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	databases.ConnectDB()

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.POST("/loans", handlers.CreateLoan)
	v1.PUT("/loans/:id/approve", handlers.ApproveLoan)
	v1.PUT("/loans/:id/invest", handlers.RecordInvestment)
	v1.PUT("/loans/:id/disburse", handlers.DisburseLoan)
	v1.GET("/loans/:id", handlers.GetLoan)
	v1.GET("/loans", handlers.ListLoans)

	router.Run(":8080")
}
