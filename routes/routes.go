package routes

import (
	"github.com/iragsraghu/renuka-trust-app/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/donations", handlers.CreateDonation)
	r.GET("/donations", handlers.GetDonations)
}
