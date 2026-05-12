package models

type Company struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LegalName string `json:"legal_name"`
	Address   string `json:"address"`
	Logo      string `json:"logo"`
	// Field tambahan baru:
    Email     string `json:"email"`
    Website   string `json:"website"`
    Phone     string `json:"phone"`
}
