package handlers

import (
	"net/http"
	"strconv"

	"github.com/iragsraghu/renuka-trust-app/config"
	"github.com/iragsraghu/renuka-trust-app/models"

	"github.com/gin-gonic/gin"
)

func GetVillageReport(c *gin.Context) {

	villageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid village id",
		})
		return
	}

	var village models.Village

	if err := config.DB.First(&village, villageID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Village not found",
		})
		return
	}

	// Total donors
	var totalDonors int64

	config.DB.Model(&models.Donor{}).
		Where("village_id = ?", villageID).
		Count(&totalDonors)

	// Total donations
	var totalDonations int64

	config.DB.Model(&models.Donation{}).
		Joins("JOIN donors ON donors.id = donations.donor_id").
		Where("donors.village_id = ?", villageID).
		Count(&totalDonations)

	// Total collection
	var totalCollection float64

	config.DB.Model(&models.Donation{}).
		Select("COALESCE(SUM(donations.amount),0)").
		Joins("JOIN donors ON donors.id = donations.donor_id").
		Where("donors.village_id = ?", villageID).
		Scan(&totalCollection)

	c.JSON(http.StatusOK, gin.H{
		"village_id":       village.ID,
		"village_name":     village.Name,
		"total_donors":     totalDonors,
		"total_donations":  totalDonations,
		"total_collection": totalCollection,
	})
}

func GetAllVillageReports(c *gin.Context) {

	var reports []models.VillageCollectionReport

	err := config.DB.Table("villages").
		Select(`
			villages.id as village_id,
			villages.name as village_name,
			COUNT(DISTINCT donors.id) as total_donors,
			COALESCE(SUM(donations.amount),0) as total_collection
		`).
		Joins("LEFT JOIN donors ON donors.village_id = villages.id").
		Joins("LEFT JOIN donations ON donations.donor_id = donors.id").
		Group("villages.id, villages.name").
		Scan(&reports).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reports)
}

func GetDashboard(c *gin.Context) {

	var response models.DashboardResponse

	// Total Villages
	config.DB.Model(&models.Village{}).
		Count(&response.TotalVillages)

	// Total Donors
	config.DB.Model(&models.Donor{}).
		Count(&response.TotalDonors)

	// Total Donations
	config.DB.Model(&models.Donation{}).
		Count(&response.TotalDonations)

	// Total Collection
	config.DB.Model(&models.Donation{}).
		Select("COALESCE(SUM(amount),0)").
		Scan(&response.TotalCollection)

	c.JSON(http.StatusOK, response)
}
