package helper

// HitungTotalZerra menghitung harga + PPN 11% sesuai dokumen Zerra
func HitungTotalZerra(harga float64) float64 {
	pajak := 0.11 * harga
	return harga + pajak
}