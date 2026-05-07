package handlers

import (
	"backend-quotation/config"
	"backend-quotation/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 1. CREATE (Menambahkan Employee Baru)
func CreateEmployee(c *gin.Context) {
	var emp models.Employee

	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Menggunakan GORM .Create() untuk menyimpan data karyawan baru
	if err := config.DB.Create(&emp).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, emp)
}

// 2. READ ALL (Mengambil Semua Data Employee)
func GetEmployees(c *gin.Context) {
	var employees []models.Employee

	// Menggunakan GORM .Find() langsung memetakan semua isi tabel ke slice struct
	if err := config.DB.Find(&employees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}

// 3. READ BY ID (Mengambil Satu Data Employee Berdasarkan ID)
func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	var emp models.Employee

	// Menggunakan GORM .First() untuk mencari record berdasarkan Primary Key (ID)
	if err := config.DB.First(&emp, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, emp)
}

// 4. UPDATE (Memperbarui Data Employee)
func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var emp models.Employee

	// Cari datanya dulu untuk memastikan ID karyawan tersebut beneran ada
	if err := config.DB.First(&emp, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
		return
	}

	// Ikat data baru yang dikirim dari request body JSON
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Menggunakan GORM .Save() untuk memperbarui seluruh field objek ke database
	if err := config.DB.Save(&emp).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee updated"})
}

// 5. DELETE (Menghapus Data Employee)
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	var emp models.Employee

	// Cari datanya dulu sebelum dihapus agar valid
	if err := config.DB.First(&emp, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
		return
	}

	// Menggunakan GORM .Delete() untuk menghapus record
	if err := config.DB.Delete(&emp).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}
