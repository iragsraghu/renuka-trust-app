package models

type DashboardResponse struct {
	TotalVillages   int64   `json:"total_villages"`
	TotalDonors     int64   `json:"total_donors"`
	TotalDonations  int64   `json:"total_donations"`
	TotalCollection float64 `json:"total_collection"`
}
