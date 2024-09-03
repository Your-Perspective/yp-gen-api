package service

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	_ "sync"
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
	viewCountMap sync.Map // Thread-safe map for storing view counts

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
		viewCountMap: sync.Map{}, // Initialize sync.Map

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

func (s *blogServiceImpl) FindBlogDetailByAuthorAndSlug(author string, slug string) (dto2.BlogDetailDto, error) {
	// Fetch the blog by author and slug
	blog, err := s.blogRepo.FindByUsernameAndSlug(author, slug)
	if err != nil {
		return dto2.BlogDetailDto{}, fmt.Errorf("blog by author '%s' and slug '%s' could not be found: %w", author, slug, err)
	}

	// Increment the view count
	err = s.IncrementViewCount(int(blog.ID))
	if err != nil {
		return dto2.BlogDetailDto{}, fmt.Errorf("failed to increment view count: %w", err)
	}

	// Map the Blog entity to BlogDetailDto
	blogDetail := s.blogMapper.BlogToBlogDetailDto(blog)
	return blogDetail, nil
}

// IncrementViewCount increments the view count for the given blog ID
func (s *blogServiceImpl) IncrementViewCount(id int) error {
	// Load or store the initial value if it doesn't exist
	value, _ := s.viewCountMap.LoadOrStore(id, 1)

	// Increment the value
	if currentCount, ok := value.(int); ok {
		s.viewCountMap.Store(id, currentCount+1)
		return nil
	}
	return fmt.Errorf("failed to increment view count for blog ID: %d", id)
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

	// Generate a 9-digit unique identifier
	uniqueIdentifier := utils.GenerateUniqueIdentifier()

	// Prepare the blog title for slug generation
	nameBlog := blogCreateRequestDto.BlogTitle

	// Check if the blog title contains Khmer characters
	if utils.ContainsKhmer(nameBlog) {
		nameBlog = utils.RemoveKhmerCharacters(nameBlog)
	}

	// Concatenate title and categories for the slug
	titleAndCategories := nameBlog + "-" + strings.Join(categoryNames, "-")

	// Generate a descriptive slug using the slug package and append the unique identifier
	slug := utils.Init(titleAndCategories + "-" + uniqueIdentifier)
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

func (s *blogServiceImpl) FindRecentPosts() ([]dto2.RecentPostBlogDto, error) {
	posts, err := s.blogRepo.FindRecentPosts()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

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

func (s *blogServiceImpl) FindAllBlogForAdmin() ([]dto2.BlogAdminDto, error) {
	blogs, err := s.blogRepo.FindAll()
	if err != nil {
		return nil, err
	}

	blogDtos := s.blogMapper.BlogDtoToBlogAdminDto(blogs)
	return blogDtos, nil
}

func (s *blogServiceImpl) Find6BlogsByUsernameAndCountViewer(username string) []dto2.BlogCardDto {
	blogs, err := s.blogRepo.FindRandom6ByUsername(username)
	if err != nil {
		// Handle the error, possibly log it and return an empty list
		return []dto2.BlogCardDto{}
	}

	// Use the mapper to convert the blogs to BlogCardDto
	blogCardDtos := s.blogMapper.BlogToBlogCardDto(blogs)

	return blogCardDtos
}

func (s *blogServiceImpl) Find6BlogsByCategoriesSlug(slug string) []dto2.BlogCardDto {
	blogs, err := s.blogRepo.FindTop6ByCategorySlug(slug)
	if err != nil {
		// Handle the error, possibly log it and return an empty list
		return []dto2.BlogCardDto{}
	}

	// Use the mapper to convert the blogs to BlogCardDto
	blogCardDtos := s.blogMapper.BlogToBlogCardDto(blogs)

	return blogCardDtos
}
func (s *blogServiceImpl) UpdateBlog(blogUpdateRequestDto dto2.BlogUpdateRequestDto, slug string) error {
	// Validate the DTO
	if err := blogUpdateRequestDto.Validate(); err != nil {
		return err
	}

	// Fetch the existing blog by slug
	blog, err := s.blogRepo.FindBySlug(slug)
	if err != nil {
		return errors.New("blog not found")
	}

	// Map the updated fields from the DTO to the Blog entity
	s.blogMapper.UpdateBlog(&blog, blogUpdateRequestDto)

	// Save the updated blog
	_, err = s.blogRepo.Save(blog)
	if err != nil {
		return err
	}

	return nil
}

func (s *blogServiceImpl) DeleteBlogByChangeStatus(id uint) error {
	// Find the blog by ID using the FindById method
	blog, err := s.FindById(id)
	if err != nil {
		return err // Return an error if the blog is not found
	}

	// Set IsDeleted to true
	blog.IsDeleted = true

	// Save the changes to the database using the Save method
	if _, err := s.Save(blog); err != nil {
		return err // Return an error if the save fails
	}

	return nil // Return nil if the operation was successful
}
