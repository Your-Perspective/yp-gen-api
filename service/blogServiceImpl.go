package service

import (
	"yp-blog-api/dto"
	"yp-blog-api/mapping"
	"yp-blog-api/models"
	_ "yp-blog-api/repository"
	repositories "yp-blog-api/repository"
)

// blogServiceImpl implements the BlogService interface.
type blogServiceImpl struct {
	repositories.BlogRepository
}

// NewBlogService creates a new instance of BlogService.
func NewBlogService(blogRepo repositories.BlogRepository) BlogService {
	return &blogServiceImpl{
		BlogRepository: blogRepo,
	}
}

// Save saves a new blog to the repository.
func (s *blogServiceImpl) Save(blog models.Blog) (models.Blog, error) {
	return s.BlogRepository.Save(blog)
}

// FindById retrieves a blog by its ID.
func (s *blogServiceImpl) FindById(id uint) (models.Blog, error) {
	return s.BlogRepository.FindById(id)
}

// FindAll retrieves all blogs.
func (s *blogServiceImpl) FindAll() ([]models.Blog, error) {
	return s.BlogRepository.FindAll()
}

// Update updates an existing blog.
func (s *blogServiceImpl) Update(blog models.Blog) (models.Blog, error) {
	return s.BlogRepository.Update(blog)
}

// DeleteById deletes a blog by its ID.
func (s *blogServiceImpl) DeleteById(id uint) error {
	return s.BlogRepository.DeleteById(id)
}

// GetBlogById retrieves a blog by its ID and maps it to a BlogDetailDto.
func (s *blogServiceImpl) GetBlogById(id uint) (dto.BlogDetailDto, error) {
	blog, err := s.BlogRepository.FindById(id)
	if err != nil {
		return dto.BlogDetailDto{}, err
	}
	return mapper.ToBlogDetailDto(blog), nil
}

// GetAllBlogs retrieves all blogs and maps them to a list of BlogDto.
func (s *blogServiceImpl) GetAllBlogs() ([]dto.BlogDto, error) {
	blogs, err := s.BlogRepository.FindAll()
	if err != nil {
		return nil, err
	}
	blogDtos := make([]dto.BlogDto, len(blogs))
	for i, blog := range blogs {
		blogDtos[i] = mapper.ToBlogDto(blog)
	}
	return blogDtos, nil
}
