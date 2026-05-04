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
}
