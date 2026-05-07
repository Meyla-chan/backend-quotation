package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB sekarang bertipe *gorm.DB, bukan *sql.DB lagi
var DB *gorm.DB

func ConnectDB() {
	// 1. Ambil string koneksi database milikmu kemarin
	dsn := "host=localhost port=4321 user=postgres password=4321 dbname=quotation_db sslmode=disable"

	// 2. Buka koneksi menggunakan GORM dan driver postgres bawaan GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database menggunakan GORM:", err)
	}

	log.Println("Hore! Koneksi database menggunakan GORM Berhasil!")

	// 3. Simpan koneksi ke variabel global DB agar bisa dipakai di handler/repository
	DB = db
}