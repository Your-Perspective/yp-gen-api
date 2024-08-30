package mapper

import (
	"yp-blog-api/dto"
	"yp-blog-api/models"
)

func ToBlogDto(blog models.Blog) dto.BlogDto {
	return dto.BlogDto{
		ID:          blog.ID,
		Title:       blog.BlogTitle,
		Content:     blog.BlogContent,
		Slug:        blog.Slug,
		IsPinned:    blog.IsPin,
		Thumbnail:   blog.Thumbnail,
		CountViewer: blog.CountViewer,
		Summary:     blog.Summary,
		MinRead:     blog.MinRead,
		AuthorName:  blog.Author.UserName, // Assuming User struct has a Name field
		Tags:        mapTagsToNames(blog.Tags),
		Categories:  mapCategoriesToNames(blog.Categories),
		CreatedAt:   blog.CreatedAt,
		UpdatedAt:   blog.UpdatedAt,
	}
}

func ToBlogDetailDto(blog models.Blog) dto.BlogDetailDto {
	blogDto := ToBlogDto(blog)
	var parentTitle string
	if blog.Parent != nil {
		parentTitle = blog.Parent.BlogTitle
	}
	return dto.BlogDetailDto{
		BlogDto:         blogDto,
		ParentBlogTitle: parentTitle,
	}
}

func mapTagsToNames(tags []models.Tag) []string {
	names := make([]string, len(tags))
	for i, tag := range tags {
		names[i] = tag.Title // Access the Title field directly
	}
	return names
}

func mapCategoriesToNames(categories []models.Category) []string {
	names := make([]string, len(categories))
	for i, category := range categories {
		names[i] = category.Title // Assuming Category struct has a Name field
	}
	return names
}
