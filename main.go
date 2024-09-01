package main

import (
	"log"
	"os"
	"strings"
	"yp-blog-api/docs"

	"github.com/joho/godotenv"
	_ "yp-blog-api/docs"
	"yp-blog-api/internal/api"
	"yp-blog-api/internal/config"
	"yp-blog-api/internal/mapping"
	"yp-blog-api/internal/models"
	"yp-blog-api/internal/repository"
	"yp-blog-api/internal/service"
)

// @title backend service for blog api
// @version 1.0
// @description backend service api restfull using Gin framework
func main() {
	// Load environment variables from the .env file
	err := godotenv.Load(".env." + os.Getenv("APP_ENV"))
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	docs.SwaggerInfo.Title = os.Getenv("SWAGGER_TITLE")
	docs.SwaggerInfo.Description = os.Getenv("SWAGGER_DESCRIPTION")
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	// Determine the scheme based on the SWAGGER_HOST
	if strings.HasPrefix(os.Getenv("SWAGGER_HOST"), "https://") {
		docs.SwaggerInfo.Schemes = []string{"https"}
	} else {
		docs.SwaggerInfo.Schemes = []string{"http"}
	}
	// Initialize the database
	config.InitDatabase()

	// Ensure the database is closed when the main function ends
	defer config.CloseDatabase()

	// AutoMigrate to create/update the schema
	err = config.DB.AutoMigrate(&models.Blog{}, &models.User{}, &models.Tag{}, &models.Category{}, &models.AdvertisingBanner{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize the repositories
	blogRepo := repositories.NewBlogRepository(config.DB)
	bannerRepo := repositories.NewAdvertisingBannerRepository(config.DB)
	tagRepo := repositories.NewTagRepository(config.DB)
	categoryRepo := repositories.NewCategoryRepository(config.DB)

	// Initialize the mappers
	blogMapper := mapper.NewBlogMapper()
	bannerMapper := mapper.NewAdvertisingBannerMapper()

	// Initialize the service with all required dependencies
	blogService := service.NewBlogService(blogRepo, bannerRepo, blogMapper, bannerMapper, categoryRepo, tagRepo)

	// Set up the router with the initialized service
	router := api.SetupRouter(blogService)

	// Get the port from the environment variables
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "9090" // Default port if not set
	}

	// Start the server
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
