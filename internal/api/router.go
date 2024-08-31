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
	router.GET("/api/blogs-admin", blogController.GetAllBlogs)
	router.GET("/api/blogs-admin/:id", blogController.GetBlogById)
	router.POST("/api/blogs-admin", blogController.CreateBlog)
	router.PUT("/api/blogs-admin/:id", blogController.UpdateBlog)
	router.DELETE("/api/blogs-admin/:id", blogController.DeleteBlog)

	// project api blog
	router.GET("/api/blogs/:categoriesSlug", blogController.ListAllByCategoriesSlug)
	router.GET("/api/blogs/", blogController.ListAllByCategoriesSlug)
	router.GET("/api/blogs/@:author/:slug", blogController.GetBlogDetailByAuthorAndSlug) // Updated route

	return router
}
