package routes

import (
	"github.com/iragsraghu/renuka-trust-app/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/villages", handlers.CreateVillage)
	r.GET("/villages", handlers.GetVillages)

	r.POST("/donors", handlers.CreateDonor)
	r.GET("/donors", handlers.GetDonors)

	r.POST("/donations", handlers.CreateDonation)
	r.GET("/donations", handlers.GetDonations)

	r.GET("/donations/search", handlers.SearchDonations)

	r.GET("/reports/village/:id", handlers.GetVillageReport)
	r.GET("/reports/villages", handlers.GetAllVillageReports)

	r.GET("/dashboard", handlers.GetDashboard)

	r.GET("/receipts/:receipt_number", handlers.GenerateReceipt)

	r.GET("/donors/:id/donations", handlers.GetDonorHistory)
}
