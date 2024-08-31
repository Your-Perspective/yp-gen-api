package service

import (
	"math/rand"
	"time"
	"yp-blog-api/dto"
	mapper "yp-blog-api/mapping"
	"yp-blog-api/models"
	_ "yp-blog-api/repository"
	repositories "yp-blog-api/repository"
)

// blogServiceImpl implements the BlogService interface.
type blogServiceImpl struct {
	blogRepo     repositories.BlogRepository
	bannerRepo   *repositories.AdvertisingBannerRepository
	blogMapper   mapper.BlogMapper
	bannerMapper mapper.AdvertisingBannerMapper
}

// NewBlogService creates a new instance of blogServiceImpl
func NewBlogService(blogRepo repositories.BlogRepository, bannerRepo *repositories.AdvertisingBannerRepository, blogMapper mapper.BlogMapper, bannerMapper mapper.AdvertisingBannerMapper) *blogServiceImpl {
	return &blogServiceImpl{
		blogRepo:     blogRepo,
		bannerRepo:   bannerRepo,
		blogMapper:   blogMapper,
		bannerMapper: bannerMapper,
	}
}

func (s *blogServiceImpl) FindBlogCardByCategoriesSlug(slug string) []interface{} {
	var blogs []models.Blog
	var err error

	if slug == "" || slug == "ALL" {
		blogs, err = s.blogRepo.FindAllByPublishedAndNotDeletedOrderByCountViewerDescCreatedAtDesc()
	} else {
		blogs, err = s.blogRepo.FindBlogsByCategorySlug(slug)
	}

	if err != nil {
		// Handle the error, possibly log it and return an empty list
		return []interface{}{}
	}

	blogCardDtos := s.blogMapper.BlogToBlogCardDto(blogs)

	banners, err := s.bannerRepo.FindAllByIsDeletedIsFalse()
	if err != nil {
		// Handle the error, possibly log it and return blog cards only
		return s.convertBlogCardsToInterface(blogCardDtos)
	}

	bannerDtos := s.bannerMapper.AdvertisingBannerListToDtoList(banners)

	// Shuffle banners
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(bannerDtos), func(i, j int) {
		bannerDtos[i], bannerDtos[j] = bannerDtos[j], bannerDtos[i]
	})

	return interleaveBlogsAndBanners(blogCardDtos, bannerDtos)
}

// convertBlogCardsToInterface converts a slice of BlogCardDto to a slice of empty interfaces
func (s *blogServiceImpl) convertBlogCardsToInterface(blogCards []dto.BlogCardDto) []interface{} {
	result := make([]interface{}, len(blogCards))
	for i, blogCard := range blogCards {
		result[i] = blogCard
	}
	return result
}

// interleaveBlogsAndBanners interleaves blog cards and banners
func interleaveBlogsAndBanners(blogs []dto.BlogCardDto, banners []dto.AdvertisingBannerDto) []interface{} {
	var result []interface{}
	bannerIndex := 0
	bannerCount := len(banners)

	for i := 0; i < len(blogs); i++ {
		result = append(result, blogs[i])
		// Add a banner after every few blogs (e.g., every 3 blogs)
		if (i+1)%3 == 0 && bannerIndex < bannerCount {
			result = append(result, banners[bannerIndex])
			bannerIndex++
		}
	}

	// Add any remaining banners if necessary
	for bannerIndex < bannerCount {
		result = append(result, banners[bannerIndex])
		bannerIndex++
	}

	return result
}
func (s *blogServiceImpl) FindAllBlogForAdmin() []dto.BlogAdminDto {
	//TODO implement me
	panic("implement me")
}

func (s *blogServiceImpl) FindBlogDetailByAuthorAndSlug(author string, slug string) dto.BlogDetailDto {
	//TODO implement me
	panic("implement me")
}

func (s *blogServiceImpl) Find6BlogsByUsernameAndCountViewer(username string) []dto.BlogCardDto {
	//TODO implement me
	panic("implement me")
}

func (s *blogServiceImpl) Find6BlogsByCategoriesSlug(slug string) []dto.BlogCardDto {
	//TODO implement me
	panic("implement me")
}

func (s *blogServiceImpl) RecentPost() []dto.RecentPostBlogDto {
	//TODO implement me
	panic("implement me")
}

func (s *blogServiceImpl) CreateBlog(blogCreateRequestDto dto.BlogCreateRequestDto) {
	//TODO implement me
	panic("implement me")
}

func (s *blogServiceImpl) UpdateBlog(blogUpdateRequestDto dto.BlogUpdateRequestDto, id int) {
	//TODO implement me
	panic("implement me")
}

func (s *blogServiceImpl) DeleteBlogById(id int) {
	//TODO implement me
	panic("implement me")
}

func (s *blogServiceImpl) DeleteBlogByChangeStatus(id int) {
	//TODO implement me
	panic("implement me")
}

// NewBlogService creates a new instance of BlogService.

// Save saves a new blog to the repository.
func (s *blogServiceImpl) Save(blog models.Blog) (models.Blog, error) {
	return s.blogRepo.Save(blog)
}

// FindById retrieves a blog by its ID.
func (s *blogServiceImpl) FindById(id uint) (models.Blog, error) {
	return s.blogRepo.FindById(id)
}

// FindAll retrieves all blogs.
func (s *blogServiceImpl) FindAll() ([]models.Blog, error) {
	return s.blogRepo.FindAll()
}

// Update updates an existing blog.
func (s *blogServiceImpl) Update(blog models.Blog) (models.Blog, error) {
	return s.blogRepo.Update(blog)
}

// DeleteById deletes a blog by its ID.
func (s *blogServiceImpl) DeleteById(id uint) error {
	return s.blogRepo.DeleteById(id)
}
