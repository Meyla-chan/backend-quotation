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

<<<<<<< HEAD
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
=======
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

