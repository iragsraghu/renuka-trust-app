package utils

import (
	"fmt"
	"time"

	"github.com/iragsraghu/renuka-trust-app/config"
	"github.com/iragsraghu/renuka-trust-app/models"
)

func GenerateReceiptNumber() string {

	year := time.Now().Year()

	var count int64

	config.DB.Model(&models.Donation{}).Count(&count)

	count++

	return fmt.Sprintf(
		"RYD-%d-%06d",
		year,
		count,
	)
}
