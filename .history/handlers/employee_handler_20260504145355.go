package handlers

import (
	"backend-quotation/config"
	"backend-quotation/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEmployee(c *gin.Context) {
	var emp models.Employee

	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	query := `
	INSERT INTO employees (name, email, phone, company_id)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`

	err := config.DB.QueryRow(
		query,
		emp.Name,
		emp.Email,
		emp.Phone,
		emp.CompanyID,
	).Scan(&emp.ID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, emp)
}