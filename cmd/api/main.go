package main

import (
	"log"
	"yp-blog-api/internal/api"
	"yp-blog-api/internal/config"
	_ "yp-blog-api/internal/mapping"
	mapper "yp-blog-api/internal/mapping"
	"yp-blog-api/internal/models"
	_ "yp-blog-api/internal/repository"
	repositories "yp-blog-api/internal/repository"
	"yp-blog-api/internal/service"
)

func main() {
	// Initialize the database
	config.InitDatabase()

	// Ensure the database is closed when the main function ends
	defer config.CloseDatabase()

	// AutoMigrate to create/update the schema
	err := config.DB.AutoMigrate(&models.Blog{}, &models.User{}, &models.Tag{}, &models.Category{}, &models.AdvertisingBanner{})
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

	// Start the server
	err = router.Run(":9090")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
