package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/iragsraghu/renuka-trust-app/config"
	"github.com/iragsraghu/renuka-trust-app/routes"
)

func main() {

	_ = godotenv.Load()

	config.ConnectDB()

	r := gin.Default()

	r.Use(cors.Default())

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
