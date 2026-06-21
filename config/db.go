package config

import (
	"fmt"

	"github.com/iragsraghu/renuka-trust-app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "root:Wsxokn@123@tcp(localhost:3306)/renuka_trust?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database Connection Failed")
	}

	db.AutoMigrate(
		&models.Village{},
		&models.Donor{},
		&models.Donation{},
	)

	fmt.Println("Database Connected Successfully")

	DB = db
}
