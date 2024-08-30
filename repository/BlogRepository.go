package repositories

import (
	"time"
	"yp-blog-api/models"
)

type BlogRepository interface {
	FindBlogsByCategorySlug(categorySlug string) ([]models.Blog, error)
	FindAllByPublishedAndNotDeletedOrderByCountViewerDescCreatedAtDesc() ([]models.Blog, error)
	FindAllByPublishedAndNotDeletedOrderByCreatedAtDesc() ([]models.Blog, error)
	FindRandom6ByUsername(username string) ([]models.Blog, error)
	FindTop6ByCategorySlug(categorySlug string) ([]models.Blog, error)
	FindByUsernameAndSlug(username, slug string) (models.Blog, error)
	FindTopAuthors(startDate time.Time, limit int) ([]map[string]interface{}, error)
	CountPinnedBlogsByAuthorId(authorId uint) (int64, error)
	FindAllByAuthorNameOrderByPinnedAndCreatedAtAndCountViewer(authorName string) ([]models.Blog, error)
	CountByAuthorEmailIgnoreCase(authorEmail string) (int64, error)

	Save(blog models.Blog) (models.Blog, error)
	FindById(id uint) (models.Blog, error)
	FindAll() ([]models.Blog, error)
	FindBySlug(slug string) (models.Blog, error)
	Update(blog models.Blog) (models.Blog, error)
	DeleteById(id uint) error
}
