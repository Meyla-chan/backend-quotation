package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {AuthRoutes(r)
    CompanyRoutes(r)
    EmployeeRoutes(r)
    RegisterBankAccountRoutes(r) // Pastikan ini terpanggil
}