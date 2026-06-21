package handlers

import (
	"net/http"

	"github.com/iragsraghu/renuka-trust-app/config"
	"github.com/iragsraghu/renuka-trust-app/models"

	"github.com/gin-gonic/gin"
)

func CreateVillage(c *gin.Context) {

	var village models.Village

	if err := c.ShouldBindJSON(&village); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := config.DB.Create(&village)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, village)
}

func GetVillages(c *gin.Context) {

	var villages []models.Village

	result := config.DB.Find(&villages)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, villages)
}
