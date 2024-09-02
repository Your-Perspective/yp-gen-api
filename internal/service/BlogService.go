package service

import (
	dto2 "yp-blog-api/internal/dto"
	"yp-blog-api/internal/models"
)

// BlogService defines the interface for blog-related operations.
type BlogService interface {
	Save(blog models.Blog) (models.Blog, error)
	FindById(id uint) (models.Blog, error)
	FindAll() ([]models.Blog, error)
	Update(blog models.Blog) (models.Blog, error)

	FindBlogCardByCategoriesSlug(slug string) []interface{}
	FindBlogDetailByAuthorAndSlug(author string, slug string) (dto2.BlogDetailDto, error)
	Find6BlogsByUsernameAndCountViewer(username string) []dto2.BlogCardDto
	Find6BlogsByCategoriesSlug(slug string) []dto2.BlogCardDto
	CreateBlog(blogCreateRequestDto dto2.BlogCreateRequestDto) error
	DeleteById(id uint) error

	UpdateBlog(blogUpdateRequestDto dto2.BlogUpdateRequestDto, id int)
	DeleteBlogByChangeStatus(id int)
	FindAllBlogForAdmin() []dto2.BlogAdminDto

	// RecentPost have issue
	RecentPost() ([]dto2.RecentPostBlogDto, error)
}
