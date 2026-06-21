package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iragsraghu/renuka-trust-app/config"
	"github.com/iragsraghu/renuka-trust-app/models"
)

func CreateDonation(c *gin.Context) {
	var donation models.Donation

	if err := c.ShouldBindJSON(&donation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	donation.ReceiptNumber = fmt.Sprintf("RYDT%d", time.Now().Unix())

	if err := config.DB.Create(&donation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, donation)
}

func GetDonations(c *gin.Context) {

	var donations []models.Donation

	result := config.DB.Order("created_at DESC").Find(&donations)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, donations)
}
