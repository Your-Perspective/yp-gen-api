package service

import "yp-blog-api/models"

// BlogService defines the interface for blog-related operations.
type BlogService interface {
	Save(blog models.Blog) (models.Blog, error)
	FindById(id uint) (models.Blog, error)
	FindAll() ([]models.Blog, error)
	Update(blog models.Blog) (models.Blog, error)
	DeleteById(id uint) error
}
