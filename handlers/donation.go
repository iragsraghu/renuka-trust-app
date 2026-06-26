package handlers

import (
	"net/http"

	"github.com/iragsraghu/renuka-trust-app/config"
	"github.com/iragsraghu/renuka-trust-app/models"
	"github.com/iragsraghu/renuka-trust-app/utils"

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
	donation.ReceiptNumber = utils.GenerateReceiptNumber()

	result := config.DB.Create(&donation)

	if result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})

		return
	}

	config.DB.
		Preload("Donor").
		Preload("Donor.Village").
		First(&donation, donation.ID)

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
		Order("id DESC").
		Find(&donations)

	if result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, donations)

}

func SearchDonations(c *gin.Context) {

	var donations []models.Donation

	name := c.Query("name")
	mobile := c.Query("mobile")

	// At least one search parameter is required
	if name == "" && mobile == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide either name or mobile number",
		})
		return
	}

	query := config.DB.
		Model(&models.Donation{}).
		Preload("Donor").
		Preload("Donor.Village").
		Joins("JOIN donors ON donors.id = donations.donor_id")

	if name != "" {
		query = query.Where("donors.name LIKE ?", "%"+name+"%")
	} else {
		query = query.Where("donors.mobile = ?", mobile)
	}

	result := query.Find(&donations)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	if len(donations) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No donations found",
		})
		return
	}

	c.JSON(http.StatusOK, donations)
}
