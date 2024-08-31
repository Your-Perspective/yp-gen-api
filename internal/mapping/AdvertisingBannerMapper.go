package mapper

import (
	"yp-blog-api/internal/dto"
	"yp-blog-api/internal/models"
)

type AdvertisingBannerMapper interface {
	AdvertisingBannerToDto(advertisingBanner models.AdvertisingBanner) dto.AdvertisingBannerDto
	AdvertisingBannerListToDtoList(advertisingBanners []models.AdvertisingBanner) []dto.AdvertisingBannerDto
}

type advertisingBannerMapperImpl struct{}

func NewAdvertisingBannerMapper() AdvertisingBannerMapper {
	return &advertisingBannerMapperImpl{}
}

func (m *advertisingBannerMapperImpl) AdvertisingBannerToDto(advertisingBanner models.AdvertisingBanner) dto.AdvertisingBannerDto {
	return dto.AdvertisingBannerDto{
		ID:       advertisingBanner.ID,
		Title:    advertisingBanner.Title,
		ImageURL: advertisingBanner.ImageURL,
		Link:     advertisingBanner.Link,
	}
}

func (m *advertisingBannerMapperImpl) AdvertisingBannerListToDtoList(advertisingBanners []models.AdvertisingBanner) []dto.AdvertisingBannerDto {
	var dtos []dto.AdvertisingBannerDto
	for _, banner := range advertisingBanners {
		dtos = append(dtos, m.AdvertisingBannerToDto(banner))
	}
	return dtos
}
