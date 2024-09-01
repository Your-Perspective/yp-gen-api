package mapper

import (
	"fmt"
	"strconv"
	"time"
	dto2 "yp-blog-api/internal/dto"
	models2 "yp-blog-api/internal/models"
)

type blogMapperImpl struct{}

func NewBlogMapper() BlogMapper {
	return &blogMapperImpl{}
}

// BlogToBlogCardDto Map multiple Blog entities to BlogCardDto
func (m *blogMapperImpl) BlogToBlogCardDto(blogs []models2.Blog) []dto2.BlogCardDto {
	var dtos []dto2.BlogCardDto
	for _, blog := range blogs {
		dtos = append(dtos, m.BlogToBlogCardDtoSingle(blog))
	}
	return dtos
}

// BlogToBlogCardDtoSingle Map a single Blog entity to BlogCardDto
func (m *blogMapperImpl) BlogToBlogCardDtoSingle(blog models2.Blog) dto2.BlogCardDto {
	return dto2.BlogCardDto{
		Slug:                 blog.Slug,
		Thumbnail:            blog.Thumbnail,
		Summary:              blog.Summary,
		BlogTitle:            blog.BlogTitle,
		FormattedCountViewer: m.formatCountViewer(blog.CountViewer),
		MinRead:              blog.MinRead,
		Published:            blog.Published,
		Author:               dto2.AuthorCardDto{UserName: blog.Author.UserName},
		CreatedAt:            GetTimeAgo(blog.CreatedAt),
	}
}

// BlogToBlogDetailDto Map a single Blog entity to BlogDetailDto
func (m *blogMapperImpl) BlogToBlogDetailDto(blog models2.Blog) dto2.BlogDetailDto {
	return dto2.BlogDetailDto{
		Slug:                 blog.Slug,
		BlogContent:          blog.BlogContent,
		Summary:              blog.Summary,
		Thumbnail:            blog.Thumbnail,
		BlogTitle:            blog.BlogTitle,
		FormattedCountViewer: m.formatCountViewer(blog.CountViewer),
		MinRead:              blog.MinRead,
		Published:            blog.Published,
		Author: dto2.AuthorCardDetailDto{
			ProfileImage: blog.Author.ProfileImage,
			UserName:     blog.Author.UserName,
			Bio:          blog.Author.Bio,
		},
		CreatedAt:           GetTimeAgo(blog.CreatedAt),
		LastModifiedTimeAgo: GetTimeAgo(blog.UpdatedAt),
		Categories:          mapCategories(blog.Categories),
		Tags:                mapTags(blog.Tags),
	}
}

// BlogToRecentPostBlogDto Map a single Blog entity to RecentPostBlogDto
func (m *blogMapperImpl) BlogToRecentPostBlogDto(blog models2.Blog) dto2.RecentPostBlogDto {
	return dto2.RecentPostBlogDto{
		BlogTitle: blog.BlogTitle,
		Slug:      blog.Slug,
		TimeAgo:   GetTimeAgo(blog.CreatedAt),
	}
}

// ToTopAuthorDTO Map an array of results to TopAuthorDto
func (m *blogMapperImpl) ToTopAuthorDTO(results []interface{}) *dto2.TopAuthorDto {
	if len(results) != 4 {
		return nil
	}
	username, _ := results[0].(string)
	bio, _ := results[1].(string)
	totalViews, _ := results[2].(int64)
	profileImage, _ := results[3].(string)

	return &dto2.TopAuthorDto{
		Username:            username,
		Bio:                 bio,
		FormattedTotalViews: m.formatTotalCountViewer(totalViews),
		ProfileImage:        profileImage,
	}
}

// ToTopAuthorDTOList Map a list of results to a list of TopAuthorDto
func (m *blogMapperImpl) ToTopAuthorDTOList(results [][]interface{}) []*dto2.TopAuthorDto {
	var dtos []*dto2.TopAuthorDto
	for _, result := range results {
		dtos = append(dtos, m.ToTopAuthorDTO(result))
	}
	return dtos
}

// CreateBlogDtoToBlog Map BlogCreateRequestDto to Blog entity
func (m *blogMapperImpl) CreateBlogDtoToBlog(dto dto2.BlogCreateRequestDto) models2.Blog {
	return models2.Blog{
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
func (m *blogMapperImpl) UpdateBlog(blog *models2.Blog, dto dto2.BlogUpdateRequestDto) {
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
func (m *blogMapperImpl) BlogDtoToBlogAdminDto(blogs []models2.Blog) []dto2.BlogAdminDto {
	var dtos []dto2.BlogAdminDto
	for range blogs {
		// Assuming you have defined BlogAdminDto struct elsewhere
		dtos = append(dtos, dto2.BlogAdminDto{
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
func mapCategories(categories []models2.Category) []dto2.CategoryDto {
	var dtos []dto2.CategoryDto
	for _, category := range categories {
		dtos = append(dtos, dto2.CategoryDto{
			ID:    int64(category.ID),
			Slug:  category.Slug,
			Title: category.Title,
		})
	}
	return dtos
}

func mapTags(tags []models2.Tag) []dto2.TagDto {
	var dtos []dto2.TagDto
	for _, tag := range tags {
		dtos = append(dtos, dto2.TagDto{
			ID:    int64(tag.ID),
			Title: tag.Title,
		})
	}
	return dtos
}

// GetTimeAgo TimeAgo utility functions
// GetTimeAgo returns a human-readable string representing the time ago from the given time.
func GetTimeAgo(t time.Time) string {
	now := time.Now()
	duration := now.Sub(t)

	switch {
	case duration.Hours() > 24*365:
		// More than a year ago: show year, month, and day
		return t.Format("2006-01-02")
	case duration.Hours() > 24*7:
		// More than 7 days ago: show month and day
		return t.Format("01-02")
	case duration.Hours() >= 24:
		// More than 1 day ago: show the number of days ago
		return fmt.Sprintf("%d days ago", int(duration.Hours()/24))
	case duration.Hours() >= 1:
		// More than 1 hour ago: show the number of hours ago
		return fmt.Sprintf("%d hours ago", int(duration.Hours()))
	case duration.Minutes() >= 1:
		// More than 1 minute ago: show the number of minutes ago
		return fmt.Sprintf("%d minutes ago", int(duration.Minutes()))
	default:
		// Less than 1 minute ago: show the number of seconds ago
		return fmt.Sprintf("%d seconds ago", int(duration.Seconds()))
	}
}

// GetLastModifiedTimeAgo returns a human-readable string representing the time ago from the given time.
// This is a wrapper function around GetTimeAgo.
func GetLastModifiedTimeAgo(t time.Time) string {
	return GetTimeAgo(t)
}
