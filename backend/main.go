package main

import (
	"github.com/Itsjoses/sebook-be/database"
	"github.com/Itsjoses/sebook-be/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	database.ConnectDB()
}

func main() {
	r := gin.Default()
	r.Use(corsMiddleware())
	routes.SetupRoutes(r)
	r.Run()
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Izinkan akses dari origin yang berbeda (ganti "*")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// Izinkan metode HTTP yang diijinkan
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Izinkan header HTTP yang diijinkan
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Tangani permintaan OPTIONS (preflight request)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
