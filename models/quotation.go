package models

import "time"

type Quotation struct {
	ID            uint            `gorm:"primaryKey" json:"id"`
	CustomerName  string          `json:"customer_name"`
	CompanyID     uint            `json:"company_id"`     
	EmployeeID    uint            `json:"employee_id"`    
	BankAccountID uint            `json:"bank_account_id"` 
	Date          time.Time       `json:"date"`
	Subtotal      float64         `json:"subtotal"`
	DiscountPerc  float64         `json:"discount_percent"` 
	DiscountAmt   float64         `json:"discount_amount"`  
	Tax           float64         `json:"tax"`              
	GrandTotal    float64         `json:"grand_total"`
	Items         []QuotationItem `gorm:"foreignKey:QuotationID;constraint:OnDelete:CASCADE" json:"items"`
}

type QuotationItem struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	QuotationID uint    `json:"quotation_id"`
	ItemNo      int     `json:"item_no"` // Nomor urut 1, 2, 3...
	Description string  `json:"description"`
	Qty         int     `json:"qty"`
	Price       float64 `json:"price"`
	Total       float64 `json:"total"`
}