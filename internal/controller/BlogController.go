package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"yp-blog-api/internal/dto"
	"yp-blog-api/internal/handler"
	"yp-blog-api/internal/models"
	"yp-blog-api/internal/service"
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
// @Summary Get all blogs
// @Description Get the list of all blogs
// @Tags Blog
// @Produce  json
// @Success 200 {array} models.Blog
// @Failure 500 {object} handler.ErrorResponse
// @Router /api/blogs-admin [get]
func (ctrl *BlogController) GetAllBlogs(c *gin.Context) {
	blogs, err := ctrl.blogService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Error: "Internal Server Error", Message: "Failed to fetch blogs"})
		return
	}
	c.JSON(http.StatusOK, blogs)
}

// GetBlogById handles GET requests to fetch a blog by ID
// @Summary Get a blog by ID
// @Description Get details of a blog by its ID
// @Tags Blog
// @Produce  json
// @Param id path int true "Blog ID"
// @Success 200 {object} models.Blog
// @Failure 400 {object} handler.ErrorResponse
// @Failure 404 {object} handler.ErrorResponse
// @Router /api/blogs-admin/{id} [get]
func (ctrl *BlogController) GetBlogById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponse{Error: "Bad Request", Message: "Invalid blog ID"})
		return
	}

	blog, err := ctrl.blogService.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, handler.ErrorResponse{Error: "Not Found", Message: "Blog not found"})
		return
	}
	c.JSON(http.StatusOK, blog)
}

// CreateBlogAdmin handles POST requests to create a new blog
// @Summary Create a new blog
// @Description Create a new blog with the provided details
// @Tags Blog
// @Accept  json
// @Produce  json
// @Param blog body models.Blog true "Blog data"
// @Success 200 {object} models.Blog
// @Failure 400 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /api/blogs-admin [post]
func (ctrl *BlogController) CreateBlogAdmin(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		var errorDetails handler.ErrorResponse
		errorDetails.Error = "Invalid input"
		errorDetails.Message = "Failed to parse blog data"

		c.JSON(http.StatusBadRequest, errorDetails)
		return
	}

	savedBlog, err := ctrl.blogService.Save(blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Error: "Internal Server Error", Message: "Failed to create blog"})
		return
	}

	c.JSON(http.StatusOK, savedBlog)
}

// UpdateBlog handles PUT requests to update an existing blog
// @Summary Update an existing blog
// @Description Update a blog by its ID
// @Tags Blog
// @Accept  json
// @Produce  json
// @Param id path int true "Blog ID"
// @Param blog body models.Blog true "Blog data"
// @Success 200 {object} models.Blog
// @Failure 400 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /api/blogs-admin/{id} [put]
func (ctrl *BlogController) UpdateBlog(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponse{Error: "Bad Request", Message: "Invalid blog ID"})
		return
	}

	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponse{Error: "Invalid input", Message: "Failed to parse blog data"})
		return
	}

	blog.ID = uint(id)
	updatedBlog, err := ctrl.blogService.Update(blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Error: "Internal Server Error", Message: "Failed to update blog"})
		return
	}

	c.JSON(http.StatusOK, updatedBlog)
}

// DeleteBlog handles DELETE requests to delete a blog by ID
// @Summary Delete a blog
// @Description Delete a blog by its ID
// @Tags Blog
// @Param id path int true "Blog ID"
// @Success 200 {object} handler.SuccessResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /api/blogs-admin/{id} [delete]
func (ctrl *BlogController) DeleteBlog(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponse{Error: "Bad Request", Message: "Invalid blog ID"})
		return
	}

	if err := ctrl.blogService.DeleteById(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Error: "Internal Server Error", Message: "Failed to delete blog"})
		return
	}

	c.JSON(http.StatusOK, handler.SuccessResponse{Message: "Blog deleted successfully"})
}

// ListAllByCategoriesSlug handles GET requests to list blog cards by category slug
// @Summary List blogs by category slug
// @Description List all blogs under a specific category identified by its slug
// @Tags Blog
// @Produce  json
// @Param categoriesSlug path string true "Category Slug"
// @Success 200 {array} models.Blog
// @Router /api/blogs/{categoriesSlug} [get]
// @Router /api/blogs/ [get]
func (ctrl *BlogController) ListAllByCategoriesSlug(c *gin.Context) {
	// Get the categoriesSlug from the URL parameters
	slug := c.Param("categoriesSlug")

	// Call the service to get the list of blog cards
	blogCards := ctrl.blogService.FindBlogCardByCategoriesSlug(slug)

	// Respond with the result in JSON format
	c.JSON(http.StatusOK, blogCards)
}

// CreateBlog handles POST requests to create a new blog
// @Summary Create a new blog
// @Description Create a new blog with the provided details
// @Tags Blog
// @Accept  json
// @Produce  json
// @Param blog body dto.BlogCreateRequestDto true "Blog data"
// @Success 200 {object} handler.SuccessResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /api/blogs [post]
func (ctrl *BlogController) CreateBlog(c *gin.Context) {
	var blogCreateRequestDto dto.BlogCreateRequestDto

	// Bind JSON input to the DTO
	if err := c.ShouldBindJSON(&blogCreateRequestDto); err != nil {
		var errorDetails handler.ErrorResponse
		errorDetails.Error = "Invalid input"
		errorDetails.Message = "Failed to parse blog data"

		// Collect detailed error information
		var errs validator.ValidationErrors
		errors.As(err, &errs)
		errorDetails.Fields = make(map[string]string)
		for _, e := range errs {
			errorDetails.Fields[e.Field()] = e.Error() // or any custom error message format
		}

		c.JSON(http.StatusBadRequest, errorDetails)
		return
	}

	// Call the service layer to create the blog
	if err := ctrl.blogService.CreateBlog(blogCreateRequestDto); err != nil {
		c.JSON(http.StatusInternalServerError, handler.ErrorResponse{
			Error:   "Internal Server Error",
			Message: err.Error(),
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, handler.SuccessResponse{Message: "Blog created successfully"})
}

func (ctrl *BlogController) GetBlogDetailByAuthorAndSlug(c *gin.Context) {
	// Extract the 'author' and 'slug' parameters from the URL
	author := c.Param("author")
	slug := c.Param("slug")

	// Call the service method to find the blog detail
	blogDetail, err := ctrl.blogService.FindBlogDetailByAuthorAndSlug(author, slug)
	if err != nil {
		// Handle the error, respond with 404 if the blog is not found
		c.JSON(http.StatusNotFound, handler.ErrorResponse{Error: "Not Found", Message: "Blog not found"})
		return
	}

	// Respond with the blog detail
	c.JSON(http.StatusOK, blogDetail)
}

// GetRecentPosts handles GET requests to fetch recent blog posts
// @Summary blog recent post
// @Description Get the most recent and popular blog posts
// @Tags Blog
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.RecentPostBlogDto
// @Failure 400 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /api/blogs/recent-post [get]
func (ctrl *BlogController) GetRecentPosts(c *gin.Context) {
	// Call the service to get the recent posts
	recentPosts := ctrl.blogService.RecentPost()

	// Respond with the result in JSON format
	c.JSON(http.StatusOK, recentPosts)
}
