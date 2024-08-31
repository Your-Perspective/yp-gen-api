package main

import (
	"log"
	"yp-blog-api/controller"
	database "yp-blog-api/db"
	"yp-blog-api/mapping"
	"yp-blog-api/models"
	repositories "yp-blog-api/repository"
	"yp-blog-api/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	database.InitDatabase()

	// Ensure the database is closed when the main function ends
	defer database.CloseDatabase()

	// AutoMigrate to create/update the schema
	err := database.DB.AutoMigrate(&models.Blog{}, &models.User{}, &models.Tag{}, &models.Category{}, &models.AdvertisingBanner{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize the repositories
	blogRepo := repositories.NewBlogRepository(database.DB)
	bannerRepo := repositories.NewAdvertisingBannerRepository(database.DB)

	// Initialize the mappers
	blogMapper := mapper.NewBlogMapper()
	bannerMapper := mapper.NewAdvertisingBannerMapper()

	// Initialize the service with all required dependencies
	blogService := service.NewBlogService(blogRepo, bannerRepo, blogMapper, bannerMapper)

	// Set up the Gin router
	router := gin.Default()

	// Initialize the controller with the service
	blogController := controller.NewBlogController(blogService)

	// Define the routes
	router.GET("/blogs", blogController.GetAllBlogs)
	router.GET("/blogs/:id", blogController.GetBlogById)
	router.POST("/blogs", blogController.CreateBlog)
	router.PUT("/blogs/:id", blogController.UpdateBlog)
	router.DELETE("/blogs/:id", blogController.DeleteBlog)
	router.GET("/blogs/categories/:categoriesSlug", blogController.ListAllByCategoriesSlug)
	router.GET("/blogs/categories/", blogController.ListAllByCategoriesSlug)
	// Start the server
	err = router.Run(":9090")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
