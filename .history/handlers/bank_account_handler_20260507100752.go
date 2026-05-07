package handlers

import (
	"backend-quotation/config"
	"backend-quotation/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 1. READ ALL (Mengambil semua data bank)
func GetBankAccounts(c *gin.Context) {
	var accounts []models.BankAccount

	// Menggunakan GORM .Find() untuk menggantikan SELECT dan perulangan rows.Next() yang panjang
	if err := config.DB.Find(&accounts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

// 2. READ BY ID (Mengambil data bank berdasarkan ID tertentu)
func GetBankAccountByID(c *gin.Context) {
	id := c.Param("id")
	var a models.BankAccount

	// Menggunakan GORM .First() untuk mencari data berdasarkan ID (kunci primer)
	if err := config.DB.First(&a, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, a)
}

// 3. UPDATE (Memperbarui data bank)
func UpdateBankAccount(c *gin.Context) {
	id := c.Param("id")
	var a models.BankAccount

	// Cari datanya dulu di database untuk memastikan ID tersebut beneran ada
	if err := config.DB.First(&a, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan untuk diperbarui"})
		return
	}

	// Ikat data baru yang dikirim dari Postman
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Menggunakan GORM .Save() untuk langsung meng-update seluruh kolom struct ke database
	if err := config.DB.Save(&a).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

// 4. CREATE (Menambahkan data bank baru)
func CreateBankAccount(c *gin.Context) {
	var a models.BankAccount
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Menggunakan GORM .Create() untuk menyimpan objek data baru. 
	// ID baru akan otomatis terisi ke dalam objek 'a' setelah berhasil disimpan.
	if err := config.DB.Create(&a).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, a)
}

// 5. DELETE (Menghapus data bank berdasarkan ID)
func DeleteBankAccount(c *gin.Context) {
	id := c.Param("id")
	var a models.BankAccount

	// 1. Cari datanya dulu di database untuk memastikan ID tersebut beneran ada
	if err := config.DB.First(&a, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan untuk dihapus"})
		return
	}

	// 2. Jika data ada, jalankan fungsi .Delete() dari GORM
	if err := config.DB.Delete(&a).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}