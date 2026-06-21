package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/iragsraghu/renuka-trust-app/config"
	"github.com/iragsraghu/renuka-trust-app/models"

	"github.com/gin-gonic/gin"
)

func CreateDonation(c *gin.Context) {

	var donation models.Donation

	if err := c.ShouldBindJSON(&donation); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// Check donor exists

	var donor models.Donor

	if err := config.DB.First(&donor, donation.DonorID).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Donor not found",
		})

		return
	}

	// Generate receipt number

	donation.ReceiptNumber = fmt.Sprintf(
		"RYD%d",
		time.Now().Unix(),
	)

	result := config.DB.Create(&donation)

	if result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Donation created successfully",
		"donation": donation,
	})
}

func GetDonations(c *gin.Context) {

	var donations []models.Donation

	result := config.DB.
		Preload("Donor").
		Preload("Donor.Village").
		Order("created_at DESC").
		Find(&donations)

	if result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, donations)

}
