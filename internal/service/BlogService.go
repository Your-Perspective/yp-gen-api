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
	DeleteById(id uint) error

	FindBlogCardByCategoriesSlug(slug string) []interface{}
	FindAllBlogForAdmin() []dto2.BlogAdminDto
	FindBlogDetailByAuthorAndSlug(author string, slug string) dto2.BlogDetailDto
	Find6BlogsByUsernameAndCountViewer(username string) []dto2.BlogCardDto
	Find6BlogsByCategoriesSlug(slug string) []dto2.BlogCardDto
	RecentPost() []dto2.RecentPostBlogDto
	CreateBlog(blogCreateRequestDto dto2.BlogCreateRequestDto) error
	UpdateBlog(blogUpdateRequestDto dto2.BlogUpdateRequestDto, id int)
	DeleteBlogById(id int)
	DeleteBlogByChangeStatus(id int)
}
