package repositories

import (
	"errors"
	"gorm.io/gorm"
	"yp-blog-api/internal/models"
)

// AdvertisingBannerRepository handles database operations for AdvertisingBanner
type AdvertisingBannerRepository struct {
	db *gorm.DB
}

// NewAdvertisingBannerRepository creates a new instance of AdvertisingBannerRepository
func NewAdvertisingBannerRepository(db *gorm.DB) *AdvertisingBannerRepository {
	return &AdvertisingBannerRepository{db: db}
}

// FindAllByIsDeletedIsFalse retrieves all non-deleted advertising banners
func (r *AdvertisingBannerRepository) FindAllByIsDeletedIsFalse() ([]models.AdvertisingBanner, error) {
	var banners []models.AdvertisingBanner
	err := r.db.Where("is_deleted = ?", false).Find(&banners).Error
	if err != nil {
		return nil, err
	}
	return banners, nil
}

// FindById retrieves an advertising banner by ID
func (r *AdvertisingBannerRepository) FindById(id int64) (*models.AdvertisingBanner, error) {
	var banner models.AdvertisingBanner
	err := r.db.First(&banner, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &banner, err
}

// Save saves or updates an advertising banner
func (r *AdvertisingBannerRepository) Save(banner *models.AdvertisingBanner) error {
	return r.db.Save(banner).Error
}

// Delete removes an advertising banner from the database
func (r *AdvertisingBannerRepository) Delete(banner *models.AdvertisingBanner) error {
	return r.db.Delete(banner).Error
}
