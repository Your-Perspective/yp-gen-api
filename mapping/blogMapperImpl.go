package mapper

import (
	"fmt"
	"strconv"
	"time"
	"yp-blog-api/dto"
	"yp-blog-api/models"
)

type blogMapperImpl struct{}

func NewBlogMapper() BlogMapper {
	return &blogMapperImpl{}
}

// BlogToBlogCardDto Map multiple Blog entities to BlogCardDto
func (m *blogMapperImpl) BlogToBlogCardDto(blogs []models.Blog) []dto.BlogCardDto {
	var dtos []dto.BlogCardDto
	for _, blog := range blogs {
		dtos = append(dtos, m.BlogToBlogCardDtoSingle(blog))
	}
	return dtos
}

// BlogToBlogCardDtoSingle Map a single Blog entity to BlogCardDto
func (m *blogMapperImpl) BlogToBlogCardDtoSingle(blog models.Blog) dto.BlogCardDto {
	return dto.BlogCardDto{
		Slug:                 blog.Slug,
		Thumbnail:            blog.Thumbnail,
		Summary:              blog.Summary,
		BlogTitle:            blog.BlogTitle,
		FormattedCountViewer: m.formatCountViewer(blog.CountViewer),
		MinRead:              blog.MinRead,
		Published:            blog.Published,
		Author:               dto.AuthorCardDto{UserName: blog.Author.UserName},
		CreatedAt:            GetTimeAgo(blog.CreatedAt),
	}
}

// BlogToBlogDetailDto Map a single Blog entity to BlogDetailDto
func (m *blogMapperImpl) BlogToBlogDetailDto(blog models.Blog) dto.BlogDetailDto {
	return dto.BlogDetailDto{
		Slug:                 blog.Slug,
		BlogContent:          blog.BlogContent,
		Summary:              blog.Summary,
		Thumbnail:            blog.Thumbnail,
		BlogTitle:            blog.BlogTitle,
		FormattedCountViewer: m.formatCountViewer(blog.CountViewer),
		MinRead:              blog.MinRead,
		Published:            blog.Published,
		Author:               dto.AuthorCardDetailDto{UserName: blog.Author.UserName},
		CreatedAt:            GetTimeAgo(blog.CreatedAt),
		LastModifiedTimeAgo:  GetTimeAgo(blog.UpdatedAt),
		Categories:           mapCategories(blog.Categories),
		Tags:                 mapTags(blog.Tags),
	}
}

// BlogToRecentPostBlogDto Map a single Blog entity to RecentPostBlogDto
func (m *blogMapperImpl) BlogToRecentPostBlogDto(blog models.Blog) dto.RecentPostBlogDto {
	return dto.RecentPostBlogDto{
		BlogTitle: blog.BlogTitle,
		Slug:      blog.Slug,
		TimeAgo:   GetTimeAgo(blog.CreatedAt),
	}
}

// ToTopAuthorDTO Map an array of results to TopAuthorDto
func (m *blogMapperImpl) ToTopAuthorDTO(results []interface{}) *dto.TopAuthorDto {
	if len(results) != 4 {
		return nil
	}
	username, _ := results[0].(string)
	bio, _ := results[1].(string)
	totalViews, _ := results[2].(int64)
	profileImage, _ := results[3].(string)

	return &dto.TopAuthorDto{
		Username:            username,
		Bio:                 bio,
		FormattedTotalViews: m.formatTotalCountViewer(totalViews),
		ProfileImage:        profileImage,
	}
}

// ToTopAuthorDTOList Map a list of results to a list of TopAuthorDto
func (m *blogMapperImpl) ToTopAuthorDTOList(results [][]interface{}) []*dto.TopAuthorDto {
	var dtos []*dto.TopAuthorDto
	for _, result := range results {
		dtos = append(dtos, m.ToTopAuthorDTO(result))
	}
	return dtos
}

// CreateBlogDtoToBlog Map BlogCreateRequestDto to Blog entity
func (m *blogMapperImpl) CreateBlogDtoToBlog(dto dto.BlogCreateRequestDto) models.Blog {
	return models.Blog{
		BlogTitle:   dto.BlogTitle,
		Published:   dto.Published,
		BlogContent: dto.BlogContent,
		Slug:        dto.Slug,
		IsPin:       dto.IsPin,
		Thumbnail:   dto.Thumbnail,
		Summary:     dto.Summary,
		MinRead:     dto.MinRead,
		// Additional fields can be mapped as needed
	}
}

// UpdateBlog Update an existing Blog entity with BlogUpdateRequestDto
func (m *blogMapperImpl) UpdateBlog(blog *models.Blog, dto dto.BlogUpdateRequestDto) {
	if dto.BlogTitle != "" {
		blog.BlogTitle = dto.BlogTitle
	}
	if dto.Published {
		blog.Published = dto.Published
	}
	if dto.BlogContent != "" {
		blog.BlogContent = dto.BlogContent
	}
	// Handle other fields similarly
}

// BlogDtoToBlogAdminDto Map a list of Blog entities to BlogAdminDto
func (m *blogMapperImpl) BlogDtoToBlogAdminDto(blogs []models.Blog) []dto.BlogAdminDto {
	var dtos []dto.BlogAdminDto
	for range blogs {
		// Assuming you have defined BlogAdminDto struct elsewhere
		dtos = append(dtos, dto.BlogAdminDto{
			// Map fields accordingly
		})
	}
	return dtos
}

// Helper function to format the count viewer
func (m *blogMapperImpl) formatCountViewer(countViewer int) string {
	if countViewer >= 10000 {
		return fmt.Sprintf("%.1fk", float64(countViewer)/1000.0)
	}
	return strconv.Itoa(countViewer)
}

// Helper function to format the total count viewer
func (m *blogMapperImpl) formatTotalCountViewer(countViewer int64) string {
	if countViewer >= 10000 {
		return fmt.Sprintf("%.1fk", float64(countViewer)/1000.0)
	}
	return strconv.FormatInt(countViewer, 10)
}

// Mapping functions for categories and tags
func mapCategories(categories []models.Category) []dto.CategoryDto {
	var dtos []dto.CategoryDto
	for _, category := range categories {
		dtos = append(dtos, dto.CategoryDto{
			ID:    int64(category.ID),
			Slug:  category.Slug,
			Title: category.Title,
		})
	}
	return dtos
}

func mapTags(tags []models.Tag) []dto.TagDto {
	var dtos []dto.TagDto
	for _, tag := range tags {
		dtos = append(dtos, dto.TagDto{
			ID:    int64(tag.ID),
			Title: tag.Title,
		})
	}
	return dtos
}

// GetTimeAgo TimeAgo utility functions
func GetTimeAgo(t time.Time) string {
	duration := time.Since(t)
	switch {
	case duration.Hours() < 24:
		return fmt.Sprintf("%d hours ago", int(duration.Hours()))
	case duration.Hours() < 48:
		return "Yesterday"
	default:
		return fmt.Sprintf("%d days ago", int(duration.Hours()/24))
	}
}
