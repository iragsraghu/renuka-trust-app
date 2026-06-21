package models

import "time"

type DonorHistoryResponse struct {
	Donor struct {
		ID      uint   `json:"id"`
		Name    string `json:"name"`
		Mobile  string `json:"mobile"`
		Village string `json:"village"`
	} `json:"donor"`

	TotalDonations int64 `json:"total_donations"`

	TotalAmount float64 `json:"total_amount"`

	Donations []DonationHistoryItem `json:"donations"`
}

type DonationHistoryItem struct {
	ReceiptNumber string `json:"receipt_number"`

	Amount float64 `json:"amount"`

	PaymentMode string `json:"payment_mode"`

	Purpose string `json:"purpose"`

	CreatedAt time.Time `json:"created_at"`
}
