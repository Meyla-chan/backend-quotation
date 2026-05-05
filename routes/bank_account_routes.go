package routes

import (
	"backend-quotation/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterBankAccountRoutes(r *gin.Engine) {
    r.GET("/bank_accounts", handlers.GetBankAccounts)          // Read All
    r.GET("/bank_accounts/:id", handlers.GetBankAccountByID)   // Read By ID
    r.POST("/bank_accounts", handlers.CreateBankAccount)       // Create
    r.PUT("/bank_accounts/:id", handlers.UpdateBankAccount)    // Update
}