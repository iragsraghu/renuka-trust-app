package handlers

import (
	"net/http"
	"strconv"

	"github.com/iragsraghu/renuka-trust-app/config"
	"github.com/iragsraghu/renuka-trust-app/models"

	"github.com/gin-gonic/gin"
)

func GetDonorHistory(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid donor id",
		})

		return
	}

	var donor models.Donor

	err = config.DB.
		Preload("Village").
		First(&donor, id).Error

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Donor not found",
		})

		return
	}

	var donations []models.Donation

	config.DB.
		Where("donor_id = ?", id).
		Order("created_at DESC").
		Find(&donations)

	var totalAmount float64

	for _, donation := range donations {

		totalAmount += donation.Amount
	}

	response := models.DonorHistoryResponse{}

	response.Donor.ID = donor.ID
	response.Donor.Name = donor.Name
	response.Donor.Mobile = donor.Mobile
	response.Donor.Village = donor.Village.Name

	response.TotalDonations = int64(len(donations))

	response.TotalAmount = totalAmount

	for _, d := range donations {

		response.Donations = append(
			response.Donations,
			models.DonationHistoryItem{
				ReceiptNumber: d.ReceiptNumber,
				Amount:        d.Amount,
				PaymentMode:   d.PaymentMode,
				Purpose:       d.Purpose,
				CreatedAt:     d.CreatedAt,
			},
		)
	}

	c.JSON(http.StatusOK, response)
}
