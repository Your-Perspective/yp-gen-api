package service

import (
	"yp-blog-api/dto"
	"yp-blog-api/models"
)

// BlogService defines the interface for blog-related operations.
type BlogService interface {
	Save(blog models.Blog) (models.Blog, error)
	FindById(id uint) (models.Blog, error)
	FindAll() ([]models.Blog, error)
	Update(blog models.Blog) (models.Blog, error)
	DeleteById(id uint) error

	FindBlogCardByCategoriesSlug(slug string) []interface{}
	FindAllBlogForAdmin() []dto.BlogAdminDto
	FindBlogDetailByAuthorAndSlug(author string, slug string) dto.BlogDetailDto
	Find6BlogsByUsernameAndCountViewer(username string) []dto.BlogCardDto
	Find6BlogsByCategoriesSlug(slug string) []dto.BlogCardDto
	RecentPost() []dto.RecentPostBlogDto
	CreateBlog(blogCreateRequestDto dto.BlogCreateRequestDto) error
	UpdateBlog(blogUpdateRequestDto dto.BlogUpdateRequestDto, id int)
	DeleteBlogById(id int)
	DeleteBlogByChangeStatus(id int)
}
