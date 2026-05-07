package helper

import "testing"

// Nama fungsi tes harus diawali dengan kata "Test"
func TestHitungTotalZerra(t *testing.T) {
	// 1. Tentukan input (seperti di PDF Zerra)
	hargaInput := 1000000.0 // Rp 1.000.000
	
	// 2. Tentukan ekspektasi (hasil yang seharusnya benar)
	ekspektasi := 1110000.0 // Rp 1.110.000
	
	// 3. Jalankan fungsi yang sudah kita buat tadi
	hasil := HitungTotalZerra(hargaInput)

	// 4. Bandingkan: Apakah hasil hitungan kodingan kita sesuai dengan ekspektasi?
	if hasil != ekspektasi {
		// Jika tidak sama, maka tes GAGAL (FAIL)
		t.Errorf("HITUNGAN SALAH: Harusnya %.2f, tapi kodinganmu menghasilkan %.2f", ekspektasi, hasil)
	}
}