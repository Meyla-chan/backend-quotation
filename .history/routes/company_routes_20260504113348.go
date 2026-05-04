package routes

import (
	"backend-quotation/handlers"

	"github.com/gin-gonic/gin"
)

func CompanyRoutes(r *gin.Engine) {
	r.GET("/companies", handlers.GetCompanies)
	r.GET("/companies/:id", handlers.GetCompanyByID)
	r.POST("/companies", handlers.CreateCompany)
	r.PUT("/companies/:id", handlers.UpdateCompany)
	r.DELETE("/companies/:id", handlers.DeleteCompany)
}