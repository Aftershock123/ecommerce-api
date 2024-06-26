package main

import (
	"github.com/Aftershock123/ecommerce/src/config"
	"github.com/Aftershock123/ecommerce/src/database"
	"github.com/Aftershock123/ecommerce/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	// Connect to MySQL database using configuration
	database.ConnectDB(cfg.DB)

	// Initialize Gin with server configuration
	r := routes.SetupRouter()

	// Middleware to set a default Host header
	r.Use(func(c *gin.Context) {
		c.Request.Header.Set("Host", "localhost:"+cfg.Server.Port)
		c.Next()
	})

	// Run Gin server
	r.Run(":" + cfg.Server.Port)
}
