package handlers

import (
	"backend-quotation/config"
	"backend-quotation/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCompany(c *gin.Context) {
	var company models.Company

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
	INSERT INTO companies (name, legal_name, address, logo)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`

	err := config.DB.QueryRow(
		query,
		company.Name,
		company.LegalName,
		company.Address,
		company.Logo,
	).Scan(&company.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, company)
}

func GetCompanyByID(c *gin.Context) {
	id := c.Param("id")

	var company models.Company

	query := `
	SELECT id, name, legal_name, address, logo
	FROM companies
	WHERE id=$1
	`

	err := config.DB.QueryRow(query, id).Scan(
		&company.ID,
		&company.Name,
		&company.LegalName,
		&company.Address,
		&company.Logo,
	)

	if err != nil {
		c.JSON(404, gin.H{"message": "Company not found"})
		return
	}

	c.JSON(200, company)
}

func GetCompanies(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, name, legal_name, address, logo FROM companies")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var companies []models.Company

	for rows.Next() {
		var company models.Company
		err := rows.Scan(
			&company.ID,
			&company.Name,
			&company.LegalName,
			&company.Address,
			&company.Logo,
		)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		companies = append(companies, company)
	}

	c.JSON(200, companies)
}

func UpdateCompany(c *gin.Context) {
	id := c.Param("id")

	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	query := `
	UPDATE companies
	SET name=$1, legal_name=$2, address=$3, logo=$4
	WHERE id=$5
	`

	result, err := config.DB.Exec(
		query,
		company.Name,
		company.LegalName,
		company.Address,
		company.Logo,
		id,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(404, gin.H{"message": "Company not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Company updated"})
}

func DeleteCompany(c *gin.Context) {
	id := c.Param("id")

	query := `DELETE FROM companies WHERE id=$1`

	result, err := config.DB.Exec(query, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(404, gin.H{"message": "Company not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Company deleted"})
}
