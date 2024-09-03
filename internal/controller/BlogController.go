package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
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

// UpdateBlog handles PUT requests to update an existing blog by its slug
// @Summary Update an existing blog
// @Description Update a blog by its slug
// @Tags Blog
// @Accept  json
// @Produce  json
// @Param slug path string true "Blog Slug"
// @Param blog body dto.BlogUpdateRequestDto true "Blog update data"
// @Router /api/blogs/{slug} [put]
func (ctrl *BlogController) UpdateBlog(ctx *gin.Context) {
	var blogUpdateRequestDto dto.BlogUpdateRequestDto
	slug := ctx.Param("slug")

	// Bind the request body to the DTO
	if err := ctx.ShouldBindJSON(&blogUpdateRequestDto); err != nil {
		ctx.JSON(http.StatusBadRequest, handler.ErrorResponse{
			Error:   "Invalid Request",
			Message: err.Error(),
		})
		return
	}

	// Validate the DTO
	validate := validator.New()
	if err := validate.Struct(blogUpdateRequestDto); err != nil {
		validationErrors := handler.FormatValidationErrors(err)
		ctx.JSON(http.StatusBadRequest, handler.ErrorResponse{
			Error:   "Validation Failed",
			Message: "Some fields did not pass validation",
			Fields:  validationErrors,
		})
		return
	}

	// Call the service to update the blog
	if err := ctrl.blogService.UpdateBlog(blogUpdateRequestDto, slug); err != nil {
		if err.Error() == "blog not found" {
			ctx.JSON(http.StatusNotFound, handler.ErrorResponse{
				Error:   "Not Found",
				Message: "Blog not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, handler.ErrorResponse{
			Error:   "Internal Server Error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, handler.SuccessResponse{
		Message: "Blog updated successfully",
	})
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

// GetBlogDetailByAuthorAndSlug
// @Tags Blog
// @Accept  json
// @Produce  json
// @Router /api/blogs/:author/:slug [get]
// @Param author path string true "Author Name"
// @Param slug path string true "Blog Slug"
// @Success 200 {object} dto.BlogDetailDto
// @Failure 404 {object} handler.ErrorResponse
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
// @Router /api/blogs/recent-posts [get]
func (ctrl *BlogController) GetRecentPosts(c *gin.Context) {
	// Log the start of the GetRecentPosts request
	log.Println("Handling request to get recent posts")

	recentPosts, err := ctrl.blogService.FindRecentPosts()
	if err != nil {
		// Log the error encountered during the service call
		log.Printf("Error occurred while getting recent posts: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Log the successful retrieval of recent posts
	log.Printf("Successfully retrieved %d recent posts", len(recentPosts))

	c.JSON(http.StatusOK, recentPosts)
}

// Find6BlogsByCategoriesSlug godoc
// @Summary Get top 6 blogs by category slug
// @Description Retrieve top 6 blogs by category slug, ordered randomly.
// @Tags Blog
// @Param slug path string true "Category Slug"
// @Produce  json
// @Success 200 {array} dto.BlogCardDto
// @Router /api/blogs/category/{slug}/top6 [get]
func (ctrl *BlogController) Find6BlogsByCategoriesSlug(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	blogCardDtos := ctrl.blogService.Find6BlogsByCategoriesSlug(slug)
	if len(blogCardDtos) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"message": "No blogs found"})
		return
	}

	c.JSON(http.StatusOK, blogCardDtos)
}

// Find6BlogsByUsernameAndCountViewer godoc
// @Summary Get top 6 blogs by username and count viewer
// @Description Retrieve top 6 blogs by username, ordered randomly, and including viewer count.
// @Tags Blog
// @Param username path string true "Username"
// @Produce  json
// @Success 200 {array} dto.BlogCardDto
// @Router /api/blogs/user/{username}/top6 [get]
func (ctrl *BlogController) Find6BlogsByUsernameAndCountViewer(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	blogCardDtos := ctrl.blogService.Find6BlogsByUsernameAndCountViewer(username)
	if len(blogCardDtos) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"message": "No blogs found"})
		return
	}

	c.JSON(http.StatusOK, blogCardDtos)
}

// DeleteBlogByChangeStatus
// @Tags Blog
// @Summary Mark a blog as deleted by changing its status
// @Description This endpoint sets the IsDeleted field of a blog to true based on the blog ID.
// @Param id path uint true "Blog ID"
// @Success 200 {string} string "Blog marked as deleted"
// @Failure 404 {object} handler.ErrorResponse
// @Router /api/blogs/{id} [delete]
func (ctrl *BlogController) DeleteBlogByChangeStatus(c *gin.Context) {
	// Extract the 'id' parameter from the URL and convert it to uint
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponse{Error: "Invalid ID", Message: "Blog ID must be a valid number"})
		return
	}

	// Call the service method to mark the blog as deleted
	err = ctrl.blogService.DeleteBlogByChangeStatus(uint(id))
	if err != nil {
		// Handle the error, respond with 404 if the blog is not found
		c.JSON(http.StatusNotFound, handler.ErrorResponse{Error: "Not Found", Message: "Blog not found"})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, "Blog marked as deleted")
}
