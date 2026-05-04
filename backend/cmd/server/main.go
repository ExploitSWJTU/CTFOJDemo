package main

import (
	"log"
	"os"
	"time"

	"swjtu-ctf-oj/internal/api"
	"swjtu-ctf-oj/internal/middleware"
	"swjtu-ctf-oj/internal/model"
	"swjtu-ctf-oj/internal/repository"
	"swjtu-ctf-oj/internal/service"

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
	err = db.AutoMigrate(&model.User{}, &model.Challenge{}, &model.Container{}, &model.Submission{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	log.Println("Database migration completed successfully")

	if err := service.EnsureDefaultAdmin(db); err != nil {
		log.Fatalf("Failed to ensure default admin: %v", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Initialize APIs
	authAPI := api.NewAuthAPI(db)
	challengeService := service.NewChallengeService(db)
	challengeAPI := api.NewChallengeAPI(db, challengeService)
	submissionAPI := api.NewSubmissionAPI(db)

	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			if err := challengeService.CleanupExpiredContainers(); err != nil {
				log.Printf("Container cleanup failed: %v", err)
			}
		}
	}()

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
			challenges.POST("/:id/container", middleware.JWTAuth(), challengeAPI.StartContainer)
			challenges.GET("/:id/container", middleware.JWTAuth(), challengeAPI.GetContainer)
			challenges.POST("/:id/container/extend", middleware.JWTAuth(), challengeAPI.ExtendContainer)
			challenges.DELETE("/:id/container", middleware.JWTAuth(), challengeAPI.DestroyContainer)
			challenges.POST("/:id/flag", middleware.JWTAuth(), challengeAPI.SubmitFlag)
		}

		submissions := api.Group("/submissions")
		submissions.Use(middleware.JWTAuth())
		{
			submissions.GET("/me", submissionAPI.GetMySubmissions)
		}

		solves := api.Group("/solves")
		solves.Use(middleware.JWTAuth())
		{
			solves.GET("/me", submissionAPI.GetMySolveRecords)
		}

		admin := api.Group("/admin")
		admin.Use(middleware.JWTAuth(), middleware.AdminOnly())
		{
			admin.GET("/challenges", challengeAPI.AdminListChallenges)
			admin.GET("/challenges/deleted", challengeAPI.AdminListDeletedChallenges)
			admin.POST("/challenges", challengeAPI.AdminCreateChallenge)
			admin.PUT("/challenges/:id", challengeAPI.AdminUpdateChallenge)
			admin.PATCH("/challenges/:id/visibility", challengeAPI.AdminSetChallengeVisibility)
			admin.DELETE("/challenges/:id", challengeAPI.AdminDeleteChallenge)
			admin.POST("/challenges/:id/restore", challengeAPI.AdminRestoreChallenge)
			admin.GET("/instances", challengeAPI.AdminListInstances)
			admin.DELETE("/instances/:id", challengeAPI.AdminDestroyInstance)
			admin.GET("/submissions", submissionAPI.AdminListSubmissions)
		}
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("BACKEND_PORT")
	}
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
