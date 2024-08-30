package main

import (
	"log"
	"yp-blog-api/controller"
	database "yp-blog-api/db"
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
	err := database.DB.AutoMigrate(&models.Blog{}, &models.User{}, &models.Tag{}, &models.Category{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize the repository, service, and controller
	blogRepo := repositories.NewBlogRepository(database.DB)
	blogService := service.NewBlogService(blogRepo)

	// Set up the Gin router
	router := gin.Default()

	blogController := controller.NewBlogController(blogService)

	router.GET("/blogs", blogController.GetAllBlogs)
	router.GET("/blogs/:id", blogController.GetBlogById)
	router.POST("/blogs", blogController.CreateBlog)
	router.PUT("/blogs/:id", blogController.UpdateBlog)
	router.DELETE("/blogs/:id", blogController.DeleteBlog)

	// Start the server
	err = router.Run(":9090")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
