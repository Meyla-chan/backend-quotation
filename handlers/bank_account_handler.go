package handlers

import (
	"backend-quotation/config"
	"backend-quotation/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// READ ALL (Sudah ada sebelumnya)
func GetBankAccounts(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, bank_name, account_name, account_number, company_id FROM bank_accounts")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var accounts []models.BankAccount
	for rows.Next() {
		var a models.BankAccount
		rows.Scan(&a.ID, &a.BankName, &a.AccountName, &a.AccountNumber, &a.CompanyID)
		accounts = append(accounts, a)
	}
	c.JSON(http.StatusOK, accounts)
}

// READ BY ID (Baru)
func GetBankAccountByID(c *gin.Context) {
	id := c.Param("id")
	var a models.BankAccount
	err := config.DB.QueryRow("SELECT id, bank_name, account_name, account_number, company_id FROM bank_accounts WHERE id = $1", id).
		Scan(&a.ID, &a.BankName, &a.AccountName, &a.AccountNumber, &a.CompanyID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, a)
}

// UPDATE (Baru)
func UpdateBankAccount(c *gin.Context) {
	id := c.Param("id")
	var a models.BankAccount
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec("UPDATE bank_accounts SET bank_name=$1, account_name=$2, account_number=$3, company_id=$4 WHERE id=$5",
		a.BankName, a.AccountName, a.AccountNumber, a.CompanyID, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

// CREATE (Sudah ada sebelumnya)
func CreateBankAccount(c *gin.Context) {
	var a models.BankAccount
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO bank_accounts (bank_name, account_name, account_number, company_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := config.DB.QueryRow(query, a.BankName, a.AccountName, a.AccountNumber, a.CompanyID).Scan(&a.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, a)
}