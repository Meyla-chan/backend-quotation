package routes

import (
	"backend-quotation/handlers"
	"github.com/gin-gonic/gin"
)

func QuotationRoutes(r *gin.Engine) {
	quotationGroup := r.Group("/quotations")
	{
		quotationGroup.POST("/", handlers.CreateQuotation)
		quotationGroup.GET("/", handlers.GetAllQuotations)
		quotationGroup.GET("/:id", handlers.GetQuotationByID) 
		quotationGroup.PUT("/:id", handlers.UpdateQuotation)
		quotationGroup.DELETE("/:id", handlers.DeleteQuotation)
	}
}