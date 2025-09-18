package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"perfume-quiz-backend/config"
	"perfume-quiz-backend/handlers"
	"perfume-quiz-backend/models"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}

	// Setup database
	db := config.SetupDatabase()
	
	// Auto migrate the schema
	db.AutoMigrate(&models.Question{}, &models.CustomerResult{})

	// Seed questions if not exists
	models.SeedQuestions(db)

	// Setup Gin
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// Routes
	api := r.Group("/api")
	{
		api.GET("/questions", handlers.GetQuestions(db))
		api.POST("/submit-quiz", handlers.SubmitQuiz(db))
		
		// Admin routes
		admin := api.Group("/admin")
		{
			admin.GET("/results", handlers.GetAllResults(db))
			admin.GET("/stats", handlers.GetStats(db))
			admin.DELETE("/results/:id", handlers.DeleteResult(db))
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	r.Run(":" + port)
}