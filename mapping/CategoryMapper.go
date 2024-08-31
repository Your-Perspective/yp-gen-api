package mapper

import (
	"yp-blog-api/dto"
	"yp-blog-api/models"
)

type CategoryMapper interface {
	IterableCategoryToIterableDto(categories []models.Category) []dto.CategoryListDto
	ListCategoryToListDto(categories []models.Category) []dto.CategoryListDto
}

type categoryMapperImpl struct{}

func NewCategoryMapper() CategoryMapper {
	return &categoryMapperImpl{}
}

func (m *categoryMapperImpl) IterableCategoryToIterableDto(categories []models.Category) []dto.CategoryListDto {
	var categoryDtos []dto.CategoryListDto
	for _, category := range categories {
		categoryDtos = append(categoryDtos, dto.CategoryListDto{
			ID:    int64(category.ID),
			Slug:  category.Slug,
			Title: category.Title,
		})
	}
	return categoryDtos
}

func (m *categoryMapperImpl) ListCategoryToListDto(categories []models.Category) []dto.CategoryListDto {
	return m.IterableCategoryToIterableDto(categories)
}
