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
=======
}
>>>>>>> 1d842f5a8d4645f8c6dbd7bb093aa3fa6d8be12a
