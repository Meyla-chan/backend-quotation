package models

type BankAccount struct {
    ID            int    `json:"id"`
    BankName      string `json:"bank_name"`      // Contoh: BCA, Mandiri
    AccountName   string `json:"account_name"`   // Contoh: PT. Zerra Teknologi Integrasi
    AccountNumber string `json:"account_number"` // Contoh: 1230000000
    CompanyID     int    `json:"company_id"`     // Relasi ke tabel companies
}