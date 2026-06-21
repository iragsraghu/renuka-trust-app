package models

type VillageCollectionReport struct {
	VillageID       uint    `json:"village_id"`
	VillageName     string  `json:"village_name"`
	TotalDonors     int64   `json:"total_donors"`
	TotalCollection float64 `json:"total_collection"`
}
