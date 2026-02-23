package main

import (
	"log"
	"os"

	"swjtu-ctf-oj/internal/api"
	"swjtu-ctf-oj/internal/middleware"
	"swjtu-ctf-oj/internal/model"
	"swjtu-ctf-oj/internal/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load database configuration from environment
	dbConfig := repository.LoadConfigFromEnv()

	// Connect to database
	db, err := repository.NewDatabase(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Auto migrate database tables
	err = db.AutoMigrate(&model.User{}, &model.Challenge{}, &model.Container{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	log.Println("Database migration completed successfully")

	// Initialize Gin router
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Initialize APIs
	authAPI := api.NewAuthAPI(db)
	challengeAPI := api.NewChallengeAPI(db)

	// API routes
	api := r.Group("/api")
	{
		// Auth routes (public)
		auth := api.Group("/auth")
		{
			auth.POST("/login", authAPI.Login)
			auth.POST("/register", authAPI.Register)
			// Protected route
			auth.GET("/me", middleware.JWTAuth(), authAPI.GetCurrentUser)
		}

		// Challenge routes (public for listing, protected for details)
		challenges := api.Group("/challenges")
		{
			challenges.GET("", challengeAPI.GetChallenges)
			challenges.GET("/categories", challengeAPI.GetChallengesByCategory)
			challenges.GET("/search", challengeAPI.SearchChallenges)
			challenges.GET("/:id", challengeAPI.GetChallenge)
		}
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
