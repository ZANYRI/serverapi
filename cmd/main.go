package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"ai-api/router"
)

func main() {

	godotenv.Load("./env/config.env")

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	router.RegisterRoutes(r)

	r.Run("0.0.0.0:8080")
}
