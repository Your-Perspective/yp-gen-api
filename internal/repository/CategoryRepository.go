package repositories

import (
	"gorm.io/gorm"
	"yp-blog-api/internal/models"
)

type CategoryRepository interface {
	FindAll() ([]models.Category, error)
	FindAllById(ids []int) ([]models.Category, error)
	FindTopCategoriesByBlogCount() ([]models.Category, error)
	FindBySlug(slug string) (*models.Category, error)
}

type categoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepositoryImpl{db: db}
}

func (r *categoryRepositoryImpl) FindAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepositoryImpl) FindAllById(ids []int) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Where("id IN ?", ids).Find(&categories).Error
	return categories, err
}

func (r *categoryRepositoryImpl) FindTopCategoriesByBlogCount() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.
		Select("categories.*, COUNT(blogs.id) as blog_count").
		Joins("JOIN blog_categories ON categories.id = blog_categories.category_id").
		Joins("JOIN blogs ON blogs.id = blog_categories.blog_id").
		Group("categories.id").
		Order("blog_count DESC").
		Find(&categories).Error
	return categories, err
}

func (r *categoryRepositoryImpl) FindBySlug(slug string) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("slug = ?", slug).First(&category).Error
	return &category, err
}
