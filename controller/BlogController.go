package controller

import (
	"net/http"
	"strconv"
	"yp-blog-api/models"
	"yp-blog-api/service"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	blogService service.BlogService
}

// NewBlogController creates a new BlogController
func NewBlogController(blogService service.BlogService) *BlogController {
	return &BlogController{
		blogService: blogService,
	}
}

// GetAllBlogs handles GET requests to fetch all blogs
func (ctrl *BlogController) GetAllBlogs(c *gin.Context) {
	blogs, err := ctrl.blogService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blogs"})
		return
	}
	c.JSON(http.StatusOK, blogs)
}

// GetBlogById handles GET requests to fetch a blog by ID
func (ctrl *BlogController) GetBlogById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	blog, err := ctrl.blogService.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}
	c.JSON(http.StatusOK, blog)
}

// CreateBlog handles POST requests to create a new blog
func (ctrl *BlogController) CreateBlog(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	savedBlog, err := ctrl.blogService.Save(blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog"})
		return
	}

	c.JSON(http.StatusOK, savedBlog)
}

// UpdateBlog handles PUT requests to update an existing blog
func (ctrl *BlogController) UpdateBlog(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	blog.ID = uint(id)
	updatedBlog, err := ctrl.blogService.Update(blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog"})
		return
	}

	c.JSON(http.StatusOK, updatedBlog)
}

// DeleteBlog handles DELETE requests to delete a blog by ID
func (ctrl *BlogController) DeleteBlog(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	if err := ctrl.blogService.DeleteById(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

// ListAllByCategoriesSlug handles GET requests to list blog cards by category slug
func (c *BlogController) ListAllByCategoriesSlug(ctx *gin.Context) {
	// Get the categoriesSlug from the URL parameters
	slug := ctx.Param("categoriesSlug")

	// Call the service to get the list of blog cards
	blogCards := c.blogService.FindBlogCardByCategoriesSlug(slug)

	// Respond with the result in JSON format
	ctx.JSON(http.StatusOK, blogCards)
}
