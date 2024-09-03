package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"yp-blog-api/internal/service"
)

type AdminController struct {
	blogService service.BlogService
}

// NewAdminController NewBlogController creates a new BlogController
func NewAdminController(blogService service.BlogService) *AdminController {
	return &AdminController{
		blogService: blogService,
	}
}

// GetAllBlogsForAdmin godoc
// @Summary Get all blogs for admin
// @Description Retrieve a list of all blogs for administrative purposes
// @Tags Admin
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.BlogAdminDto
// @Router /api/admin/blogs [get]
func (ctrl *AdminController) GetAllBlogsForAdmin(c *gin.Context) {
	log.Println("Handling request to get all blogs for admin")

	blogs, err := ctrl.blogService.FindAllBlogForAdmin()
	if err != nil {
		log.Printf("Error occurred while getting blogs for admin: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Successfully retrieved %d blogs for admin", len(blogs))
	c.JSON(http.StatusOK, blogs)
}
