package routes

import (
	"backend-quotation/handlers"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(r *gin.Engine) {
	r.POST("/employees", handlers.CreateEmployee)
	
}
