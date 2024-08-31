package service

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strings"
	"time"
	dto2 "yp-blog-api/internal/dto"
	mapper2 "yp-blog-api/internal/mapping"
	"yp-blog-api/internal/models"
	_ "yp-blog-api/internal/repository"
	repositories2 "yp-blog-api/internal/repository"
	"yp-blog-api/internal/utils"
)

// blogServiceImpl implements the BlogService interface.
type blogServiceImpl struct {
	blogRepo     repositories2.BlogRepository
	tagRepo      repositories2.TagRepository
	categoryRepo repositories2.CategoryRepository
	bannerRepo   *repositories2.AdvertisingBannerRepository
	blogMapper   mapper2.BlogMapper
	bannerMapper mapper2.AdvertisingBannerMapper
}

// NewBlogService creates a new instance of blogServiceImpl
func NewBlogService(blogRepo repositories2.BlogRepository, bannerRepo *repositories2.AdvertisingBannerRepository, blogMapper mapper2.BlogMapper, bannerMapper mapper2.AdvertisingBannerMapper, categoryRepo repositories2.CategoryRepository, TagRepo repositories2.TagRepository) *blogServiceImpl {
	return &blogServiceImpl{
		blogRepo:     blogRepo,
		bannerRepo:   bannerRepo,
		blogMapper:   blogMapper,
		bannerMapper: bannerMapper,
		categoryRepo: categoryRepo,
		tagRepo:      TagRepo,
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
func (s *blogServiceImpl) convertBlogCardsToInterface(blogCards []dto2.BlogCardDto) []interface{} {
	result := make([]interface{}, len(blogCards))
	for i, blogCard := range blogCards {
		result[i] = blogCard
	}
	return result
}

// interleaveBlogsAndBanners interleaves blog cards and banners
func interleaveBlogsAndBanners(blogs []dto2.BlogCardDto, banners []dto2.AdvertisingBannerDto) []interface{} {
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
func (s *blogServiceImpl) FindAllBlogForAdmin() []dto2.BlogAdminDto {
	//TODO implement me
	panic("implement me")
}

func (s *blogServiceImpl) FindBlogDetailByAuthorAndSlug(author string, slug string) dto2.BlogDetailDto {
	//TODO implement me
	panic("implement me")
}

func (s *blogServiceImpl) Find6BlogsByUsernameAndCountViewer(username string) []dto2.BlogCardDto {
	//TODO implement me
	panic("implement me")
}

func (s *blogServiceImpl) Find6BlogsByCategoriesSlug(slug string) []dto2.BlogCardDto {
	//TODO implement me
	panic("implement me")
}

func (s *blogServiceImpl) RecentPost() []dto2.RecentPostBlogDto {
	//TODO implement me
	panic("implement me")
}

func (s *blogServiceImpl) CreateBlog(blogCreateRequestDto dto2.BlogCreateRequestDto) error {
	// Validate the incoming DTO
	if err := blogCreateRequestDto.Validate(); err != nil {
		return fmt.Errorf("validation error: %v", err)
	}

	// Map the DTO to the Blog entity
	blog := s.blogMapper.CreateBlogDtoToBlog(blogCreateRequestDto)

	// Retrieve categories by IDs
	categories, err := s.categoryRepo.FindAllById(blogCreateRequestDto.CategoryIds)
	if err != nil {
		return fmt.Errorf("error retrieving categories: %v", err)
	}
	blog.Categories = categories

	// Concatenate category slugs for the slug generation
	var categoryNames []string
	for _, category := range categories {
		categoryNames = append(categoryNames, strings.ReplaceAll(strings.ToLower(category.Slug), " ", "-"))
	}

	// Generate a unique identifier (UUID)
	uniqueIdentifier := uuid.New().String()

	// Prepare the blog title for slug generation
	nameBlog := blogCreateRequestDto.BlogTitle

	// Check if the blog title contains Khmer characters
	if utils.ContainsKhmer(nameBlog) {
		nameBlog = utils.RemoveKhmerCharacters(nameBlog)
	}

	// Concatenate title and categories for the slug
	titleAndCategories := strings.ToLower(strings.ReplaceAll(nameBlog, " ", "-")) + "-" + strings.Join(categoryNames, "-")

	// Generate a descriptive slug and append the UUID
	slug := utils.Init("-" + titleAndCategories + "-" + uniqueIdentifier)
	blog.Slug = slug

	// Check the pinned blogs limit
	if err := s.checkPinnedBlogsLimit(int(blog.Author.ID), blog.IsPin); err != nil {
		return err
	}

	// Retrieve tags by IDs
	tags, err := s.tagRepo.FindAllById(blogCreateRequestDto.Tags)
	if err != nil {
		return fmt.Errorf("error retrieving tags: %v", err)
	}
	blog.Tags = tags

	// Save the blog and check for errors
	_, err = s.blogRepo.Save(blog)
	if err != nil {
		return fmt.Errorf("error saving blog: %v", err)
	}

	return nil
}

func (s *blogServiceImpl) checkPinnedBlogsLimit(authorID int, isPin bool) error {
	// Example check: Limit to 3 pinned blogs per author
	if isPin {
		pinnedCount, err := s.blogRepo.CountPinnedBlogsByAuthorId(uint(authorID))
		if err != nil {
			return err
		}
		if pinnedCount >= 3 {
			return fmt.Errorf("pinned blogs limit exceeded")
		}
	}
	return nil
}
func (s *blogServiceImpl) UpdateBlog(blogUpdateRequestDto dto2.BlogUpdateRequestDto, id int) {
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
