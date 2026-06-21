package handlers

import (
	"net/http"

	"github.com/iragsraghu/renuka-trust-app/config"
	"github.com/iragsraghu/renuka-trust-app/models"

	"github.com/gin-gonic/gin"
)

func CreateDonor(c *gin.Context) {

	var donor models.Donor

	if err := c.ShouldBindJSON(&donor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// check village exists
	var village models.Village

	if err := config.DB.First(&village, donor.VillageID).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Village not found",
		})

		return
	}

	result := config.DB.Create(&donor)

	if result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, donor)
}

func GetDonors(c *gin.Context) {

	var donors []models.Donor

	result := config.DB.Preload("Village").Find(&donors)

	if result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, donors)
}
