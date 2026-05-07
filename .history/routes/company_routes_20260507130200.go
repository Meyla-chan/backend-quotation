package routes

import (
	"backend-quotation/handlers"
	"backend-quotation/middleware"

	"github.com/gin-gonic/gin"
)

func CompanyRoutes(r *gin.Engine) {

	company := r.Group("/")

	// pakai middleware JWT
	company.Use(middleware.AuthMiddleware())

	company.GET("/companies", handlers.GetCompanies)
	company.GET("/companies/:id", handlers.GetCompanyByID)

	company.POST("/companies", handlers.CreateCompany)

	company.PUT("/companies/:id", handlers.UpdateCompany)

	company.DELETE("/companies/:id", handlers.DeleteCompany)
}
