package models

import "time"

type Donor struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"size:255;not null" json:"name"`
	Mobile  string `gorm:"size:15" json:"mobile"`
	Address string `gorm:"type:text" json:"address"`

	VillageID uint    `json:"village_id"`
	Village   Village `gorm:"foreignKey:VillageID" json:"village"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
