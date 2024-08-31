package api

import (
	"github.com/gin-gonic/gin"
	"yp-blog-api/internal/controller"
	"yp-blog-api/internal/service"
)

// SetupRouter initializes the Gin router with all the routes and dependencies
func SetupRouter(blogService service.BlogService) *gin.Engine {
	// Set up the Gin router
	router := gin.Default()

	// Initialize the controller with the service
	blogController := controller.NewBlogController(blogService)

	// Define the routes
	router.GET("/api/blogs", blogController.GetAllBlogs)
	router.GET("/api/blogs/:id", blogController.GetBlogById)
	router.POST("/api/blogs-admin", blogController.CreateBlog)
	router.PUT("/api/blogs/:id", blogController.UpdateBlog)
	router.DELETE("/api/blogs/:id", blogController.DeleteBlog)
	router.GET("/api/blogs/categories/:categoriesSlug", blogController.ListAllByCategoriesSlug)
	router.GET("/api/blogs/", blogController.ListAllByCategoriesSlug)
	router.POST("/api/blogs", blogController.CreateBlog)

	return router
}
