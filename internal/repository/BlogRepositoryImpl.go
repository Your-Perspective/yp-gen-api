package repositories

import (
	"errors"
	"gorm.io/gorm"
	"time"
	"yp-blog-api/internal/models"
)

type blogRepositoryImpl struct {
	db *gorm.DB
}

// NewBlogRepository creates a new instance of BlogRepositoryImpl.
func NewBlogRepository(db *gorm.DB) BlogRepository {
	return &blogRepositoryImpl{db: db}
}

func (r *blogRepositoryImpl) FindBlogsByCategorySlug(categorySlug string) ([]models.Blog, error) {
	var blogs []models.Blog
	err := r.db.Joins("JOIN blog_categories bc ON bc.blog_id = blogs.id").
		Joins("JOIN categories c ON bc.category_id = c.id").
		Where("c.slug = ? AND blogs.published = ? AND blogs.is_deleted IS FALSE", categorySlug, true).
		Order("blogs.created_at DESC, blogs.count_viewer DESC").
		Find(&blogs).Error
	return blogs, err
}
func (r *blogRepositoryImpl) FindAllByPublishedAndNotDeletedOrderByCountViewerDescCreatedAtDesc() ([]models.Blog, error) {
	var blogs []models.Blog
	err := r.db.Where("published = ? AND is_deleted IS FALSE", true).
		Order("count_viewer DESC, created_at DESC").
		Find(&blogs).Error
	return blogs, err
}

func (r *blogRepositoryImpl) FindAllByPublishedAndNotDeletedOrderByCreatedAtDesc() ([]models.Blog, error) {
	var blogs []models.Blog
	err := r.db.Where("published = ? AND is_deleted IS FALSE", true).
		Order("created_at DESC").
		Find(&blogs).Error
	return blogs, err
}

func (r *blogRepositoryImpl) FindRandom6ByUsername(username string) ([]models.Blog, error) {
	var blogs []models.Blog
	err := r.db.Joins("JOIN users u ON u.id = blogs.author_id").
		Where("u.user_name = ? AND blogs.published = ? AND blogs.deleted_at IS NULL", username, true).
		Order("RANDOM()").
		Limit(6).
		Find(&blogs).Error
	return blogs, err
}

func (r *blogRepositoryImpl) FindTop6ByCategorySlug(categorySlug string) ([]models.Blog, error) {
	var blogs []models.Blog
	err := r.db.Joins("JOIN blog_categories bc ON bc.blog_id = blogs.id").
		Joins("JOIN categories c ON bc.category_id = c.id").
		Where("c.slug = ? AND blogs.published = ? AND blogs.deleted_at IS NULL", categorySlug, true).
		Order("RANDOM()").
		Limit(6).
		Find(&blogs).Error
	return blogs, err
}
func (r *blogRepositoryImpl) FindByUsernameAndSlug(username, slug string) (models.Blog, error) {
	var blog models.Blog
	err := r.db.Preload("Author"). // Preloads the related Author (User) model
					Joins("Author").
					Where("Author.user_name = ? AND blogs.slug = ?", username, slug).
					First(&blog).Error
	return blog, err
}

func (r *blogRepositoryImpl) FindTopAuthors(startDate time.Time, limit int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := r.db.Table("blogs b").
		Select("u.user_name AS username, u.bio AS bio, SUM(b.count_viewer) AS total_views, u.profile_image AS profile_image").
		Joins("JOIN users u ON u.id = b.author_id").
		Where("b.created_at >= ?", startDate).
		Group("u.user_name, u.bio, u.profile_image").
		Order("total_views DESC").
		Limit(limit).
		Scan(&results).Error
	return results, err
}

func (r *blogRepositoryImpl) CountPinnedBlogsByAuthorId(authorId uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Blog{}).
		Where("author_id = ? AND is_pin = ?", authorId, true).
		Count(&count).Error
	return count, err
}

func (r *blogRepositoryImpl) FindAllByAuthorNameOrderByPinnedAndCreatedAtAndCountViewer(authorName string) ([]models.Blog, error) {
	var blogs []models.Blog
	err := r.db.Joins("JOIN users u ON u.id = blogs.author_id").
		Where("u.user_name = ? AND blogs.published = ? AND blogs.deleted_at IS NULL", authorName, true).
		Order("blogs.is_pin DESC, blogs.created_at DESC, blogs.count_viewer DESC").
		Find(&blogs).Error
	return blogs, err
}

func (r *blogRepositoryImpl) CountByAuthorEmailIgnoreCase(authorEmail string) (int64, error) {
	var count int64
	err := r.db.Model(&models.Blog{}).
		Joins("JOIN users u ON u.id = blogs.author_id").
		Where("LOWER(u.email) = LOWER(?)", authorEmail).
		Count(&count).Error
	return count, err
}

func (r *blogRepositoryImpl) Save(blog models.Blog) (models.Blog, error) {
	if err := r.db.Create(&blog).Error; err != nil {
		return models.Blog{}, err
	}
	return blog, nil
}

func (r *blogRepositoryImpl) FindById(id uint) (models.Blog, error) {
	var blog models.Blog
	if err := r.db.First(&blog, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Blog{}, errors.New("blog not found")
		}
		return models.Blog{}, err
	}
	return blog, nil
}

func (r *blogRepositoryImpl) FindAll() ([]models.Blog, error) {
	var blogs []models.Blog
	if err := r.db.Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func (r *blogRepositoryImpl) FindBySlug(slug string) (models.Blog, error) {
	var blog models.Blog
	if err := r.db.Where("slug = ?", slug).First(&blog).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Blog{}, errors.New("blog not found")
		}
		return models.Blog{}, err
	}
	return blog, nil
}

func (r *blogRepositoryImpl) Update(blog models.Blog) (models.Blog, error) {
	if err := r.db.Save(&blog).Error; err != nil {
		return models.Blog{}, err
	}
	return blog, nil
}

func (r *blogRepositoryImpl) DeleteById(id uint) error {
	if err := r.db.Delete(&models.Blog{}, id).Error; err != nil {
		return err
	}
	return nil
}
