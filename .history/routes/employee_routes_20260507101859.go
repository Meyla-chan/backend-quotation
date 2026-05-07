package routes

import (
	"backend-quotation/handlers"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(r *gin.Engine) {
	r.POST("/employees", handlers.CreateEmployee)
	
	r.GET("/employees", handlers.GetEmployees)       
	r.GET("/employees/:id", handlers.GetEmployeeByID)   
	r.PUT("/employees/:id", handlers.UpdateEmployee)   
	r.DELETE("/employees/:id", handlers.DeleteEmployee)  

}
}


