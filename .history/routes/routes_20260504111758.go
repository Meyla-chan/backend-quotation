package routes

import (
	"backend-quotation/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/companies", handlers.GetCompanies)
	r.POST("/companies", handlers.CreateCompany)

}