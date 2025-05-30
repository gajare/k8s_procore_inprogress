package main

import (
	"fmt"
	"log"
	"os"
	"procore-call-logs/handlers"

	// "procore-call_logs/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Gin router
	router := gin.Default()
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3002"
	}
	// Configure CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Origin", frontendURL)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Routes
	router.POST("/api/auth/token", handlers.GetAuthToken)
	router.GET("/api/call_logs", handlers.GetcallLogs)
	router.GET("/api/call_logs/filter", handlers.GetFilteredCallLogs)
	router.POST("/api/call_logs", handlers.CreateCallLog)
	router.PUT("/api/call_logs/:id", handlers.UpdateCallLog)
	router.DELETE("/api/call_logs/:id", handlers.DeleteCallLog)
	router.GET("/api/call_logs/:id", handlers.GetcallLogDetails)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}
	router.Run(":" + port)
	fmt.Println("server is starting")
}
