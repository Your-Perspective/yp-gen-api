package controller

import "yp-blog-api/internal/service"

type AuthorController struct {
	blogService service.BlogService
}

// NewAuthorController NewBlogController creates a new BlogController
func NewAuthorController(blogService service.BlogService) *AuthorController {
	return &AuthorController{
		blogService: blogService,
	}
}
