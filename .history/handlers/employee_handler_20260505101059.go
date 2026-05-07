package handlers

import (
	"backend-quotation/config"
	"backend-quotation/models"

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

func GetEmployees(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, name, email, phone, company_id FROM employees")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var employees []models.Employee

	for rows.Next() {
		var emp models.Employee
		err := rows.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Email,
			&emp.Phone,
			&emp.CompanyID,
		)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		employees = append(employees, emp)
	}

	c.JSON(200, employees)
}

func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")

	var emp models.Employee

	query := `
	SELECT id, name, email, phone, company_id
	FROM employees
	WHERE id=$1
	`

	err := config.DB.QueryRow(query, id).Scan(
		&emp.ID,
		&emp.Name,
		&emp.Email,
		&emp.Phone,
		&emp.CompanyID,
	)

	if err != nil {
		c.JSON(404, gin.H{"message": "Employee not found"})
		return
	}

	c.JSON(200, emp)
}

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")

	var emp models.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	query := `
	UPDATE employees
	SET name=$1, email=$2, phone=$3, company_id=$4
	WHERE id=$5
	`

	result, err := config.DB.Exec(
		query,
		emp.Name,
		emp.Email,
		emp.Phone,
		emp.CompanyID,
		id,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(404, gin.H{"message": "Employee not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Employee updated"})
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	query := `DELETE FROM employees WHERE id=$1`

	result, err := config.DB.Exec(query, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(404, gin.H{"message": "Employee not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Employee deleted"})
}
