package handlers

import (
	"backend-quotation/config"
	"backend-quotation/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 1. CREATE (Menambahkan Company Baru)
func CreateCompany(c *gin.Context) {
	var company models.Company

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Menggunakan GORM .Create() untuk menyimpan data company baru ke database
	if err := config.DB.Create(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, company)
}

// 2. READ BY ID (Mengambil Satu Data Company Berdasarkan ID)
func GetCompanyByID(c *gin.Context) {
	id := c.Param("id")
	var company models.Company

	// Menggunakan GORM .First() untuk mencari data berdasarkan Primary Key (ID)
	if err := config.DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, company)
}

// 3. READ ALL (Mengambil Semua Data Company)
func GetCompanies(c *gin.Context) {
	var companies []models.Company

	// Menggunakan GORM .Find() untuk langsung menarik semua record tanpa loop rows.Next()
	if err := config.DB.Find(&companies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, companies)
}

// 4. UPDATE (Memperbarui Data Company)
func UpdateCompany(c *gin.Context) {
	id := c.Param("id")
	var company models.Company

	// Cari datanya dulu untuk memastikan id tersebut beneran ada
	if err := config.DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Company not found"})
		return
	}

	// Ikat data perubahan baru dari JSON input
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Menggunakan GORM .Save() untuk memperbarui seluruh field objek ke database
	if err := config.DB.Save(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company updated"})
}

// 5. DELETE (Menghapus Data Company)
func DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	var company models.Company

	// Cari datanya dulu sebelum dihapus agar valid
	if err := config.DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Company not found"})
		return
	}

	// Menggunakan GORM .Delete() untuk menghapus record berdasarkan struct yang ditemukan
	if err := config.DB.Delete(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted"})
}