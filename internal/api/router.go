package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"yp-blog-api/internal/controller"
	"yp-blog-api/internal/service"
)

// SetupRouter initializes the Gin router with all the routes and dependencies
func SetupRouter(blogService service.BlogService) *gin.Engine {
	// Set up the Gin router
	router := gin.Default()
	//add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Initialize the controller with the service
	blogController := controller.NewBlogController(blogService)
	//authorController := controller.NewAuthorController(blogService)
	adminController := controller.NewAdminController(blogService)

	// Define the routes
	router.GET("/api/blogs-admin", blogController.GetAllBlogs)
	router.GET("/api/blogs-admin/:id", blogController.GetBlogById)
	router.PUT("/api/blogs/:slug", blogController.UpdateBlog)
	router.DELETE("/api/blogs-admin/:id", blogController.DeleteBlog)

	// project api blog
	router.GET("/api/blogs/:categoriesSlug", blogController.ListAllByCategoriesSlug)
	router.GET("/api/blogs/", blogController.ListAllByCategoriesSlug)
	router.GET("/api/blogs/@:author/:slug", blogController.GetBlogDetailByAuthorAndSlug) // Updated route
	router.POST("/api/blogs", blogController.CreateBlog)
	router.GET("/api/blogs/recent-posts", blogController.GetRecentPosts)
	router.GET("/api/blogs/category/:slug/top6", blogController.Find6BlogsByCategoriesSlug)
	router.GET("/api/blogs/user/:username/top6", blogController.Find6BlogsByUsernameAndCountViewer)
	router.GET("/api/admin/blogs", adminController.GetAllBlogsForAdmin)

	return router
}
