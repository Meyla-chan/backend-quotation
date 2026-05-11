package handlers

import (
	"backend-quotation/config"
	"backend-quotation/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CREATE
func CreateQuotation(c *gin.Context) {
	var input models.Quotation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var totalItem float64
	for i := range input.Items {
		input.Items[i].ItemNo = i + 1
		input.Items[i].Total = float64(input.Items[i].Qty) * input.Items[i].Price
		totalItem += input.Items[i].Total
	}

	input.Subtotal = totalItem
	input.DiscountAmt = totalItem * (input.DiscountPerc / 100)
	input.Tax = (totalItem - input.DiscountAmt) * 0.11
	input.GrandTotal = (totalItem - input.DiscountAmt) + input.Tax
	input.Date = time.Now()

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": input})
}

// GET ALL
func GetAllQuotations(c *gin.Context) {
	var quotations []models.Quotation
	config.DB.Preload("Items").Preload("Employee").Preload("BankAccount").Find(&quotations)
	c.JSON(http.StatusOK, gin.H{"data": quotations})
}

// GET BY ID (Ini yang tadi error/undefined)
func GetQuotationByID(c *gin.Context) {
	id := c.Param("id")
	var quotation models.Quotation
	if err := config.DB.Preload("Items").Preload("Employee").Preload("BankAccount").First(&quotation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": quotation})
}

// UPDATE
func UpdateQuotation(c *gin.Context) {
	id := c.Param("id")
	var quotation models.Quotation
	if err := config.DB.First(&quotation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	var input models.Quotation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Where("quotation_id = ?", id).Delete(&models.QuotationItem{})

	var totalItem float64
	for i := range input.Items {
		input.Items[i].QuotationID = quotation.ID
		input.Items[i].ItemNo = i + 1
		input.Items[i].Total = float64(input.Items[i].Qty) * input.Items[i].Price
		totalItem += input.Items[i].Total
	}

	input.Subtotal = totalItem
	input.DiscountAmt = totalItem * (input.DiscountPerc / 100)
	input.Tax = (totalItem - input.DiscountAmt) * 0.11
	input.GrandTotal = (totalItem - input.DiscountAmt) + input.Tax

	config.DB.Model(&quotation).Updates(input)
	if len(input.Items) > 0 {
		config.DB.Create(&input.Items)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update Berhasil", "data": input})
}

// DELETE
func DeleteQuotation(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Quotation{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}