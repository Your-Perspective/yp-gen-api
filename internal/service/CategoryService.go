package service

import (
	"yp-blog-api/internal/dto"
	"yp-blog-api/internal/mapping"
	"yp-blog-api/internal/repository"
)

type CategoryService interface {
	FindAllCategory() ([]dto.CategoryListDto, error)
	FindTopCategoriesByBlogCount() ([]dto.CategoryListDto, error)
}

type categoryServiceImpl struct {
	categoryRepo   repositories.CategoryRepository
	categoryMapper mapper.CategoryMapper
}

func NewCategoryService(categoryRepo repositories.CategoryRepository, categoryMapper mapper.CategoryMapper) CategoryService {
	return &categoryServiceImpl{
		categoryRepo:   categoryRepo,
		categoryMapper: categoryMapper,
	}
}

func (s *categoryServiceImpl) FindAllCategory() ([]dto.CategoryListDto, error) {
	categories, err := s.categoryRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return s.categoryMapper.IterableCategoryToIterableDto(categories), nil
}

func (s *categoryServiceImpl) FindTopCategoriesByBlogCount() ([]dto.CategoryListDto, error) {
	categories, err := s.categoryRepo.FindTopCategoriesByBlogCount()
	if err != nil {
		return nil, err
	}
	return s.categoryMapper.ListCategoryToListDto(categories), nil
}
