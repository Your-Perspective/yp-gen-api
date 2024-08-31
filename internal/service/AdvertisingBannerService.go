package service

import (
	"errors"
	"fmt"
	"yp-blog-api/internal/models"
	"yp-blog-api/internal/repository"
)

// AdvertisingBannerService handles the business logic for advertising banners
type AdvertisingBannerService struct {
	repo *repositories.AdvertisingBannerRepository
}

// NewAdvertisingBannerService creates a new instance of AdvertisingBannerService
func NewAdvertisingBannerService(repo *repositories.AdvertisingBannerRepository) *AdvertisingBannerService {
	return &AdvertisingBannerService{repo: repo}
}

// GetAllBanners retrieves all non-deleted advertising banners
func (s *AdvertisingBannerService) GetAllBanners() ([]models.AdvertisingBanner, error) {
	return s.repo.FindAllByIsDeletedIsFalse()
}

// GetBannerById retrieves an advertising banner by ID
func (s *AdvertisingBannerService) GetBannerById(id int64) (*models.AdvertisingBanner, error) {
	banner, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	if banner == nil {
		return nil, errors.New(fmt.Sprintf("Banner not found with ID: %d", id))
	}
	return banner, nil
}

// CreateBanner creates a new advertising banner
func (s *AdvertisingBannerService) CreateBanner(banner *models.AdvertisingBanner) (*models.AdvertisingBanner, error) {
	err := s.repo.Save(banner)
	return banner, err
}

// UpdateBanner updates an existing advertising banner
func (s *AdvertisingBannerService) UpdateBanner(id int64, bannerDetails *models.AdvertisingBanner) (*models.AdvertisingBanner, error) {
	banner, err := s.GetBannerById(id)
	if err != nil {
		return nil, err
	}

	banner.Title = bannerDetails.Title
	banner.ImageURL = bannerDetails.ImageURL
	banner.Link = bannerDetails.Link

	err = s.repo.Save(banner)
	return banner, err
}

// DeleteBanner deletes an advertising banner by ID
func (s *AdvertisingBannerService) DeleteBanner(id int64) error {
	banner, err := s.GetBannerById(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(banner)
}
