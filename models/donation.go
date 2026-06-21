package models

import "time"

type Donation struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	DonorID       uint      `json:"donor_id"`
	Donor         Donor     `gorm:"foreignKey:DonorID" json:"donor"`
	Amount        float64   `json:"amount"`
	PaymentMode   string    `gorm:"size:50" json:"payment_mode"`
	Purpose       string    `gorm:"size:255" json:"purpose"`
	ReceiptNumber string    `gorm:"size:50;uniqueIndex" json:"receipt_number"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type DonationResponse struct {
	ID            uint    `json:"id"`
	DonorName     string  `json:"donor_name"`
	VillageName   string  `json:"village_name"`
	Amount        float64 `json:"amount"`
	PaymentMode   string  `json:"payment_mode"`
	Purpose       string  `json:"purpose"`
	ReceiptNumber string  `json:"receipt_number"`
}
