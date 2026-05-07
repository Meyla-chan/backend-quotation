package routes

import (
	"backend-quotation/handlers"
	"backend-quotation/middleware"

	"github.com/gin-gonic/gin"
)

func CompanyRoutes(r *gin.Engine) {

	company := r.Group("/")
	company.Use(middleware.AuthMiddleware())

	r.GET("/companies", handlers.GetCompanies)
	r.GET("/companies/:id", handlers.GetCompanyByID)
	r.POST("/companies", handlers.CreateCompany)
	r.PUT("/companies/:id", handlers.UpdateCompany)
	r.DELETE("/companies/:id", handlers.DeleteCompany)
}
