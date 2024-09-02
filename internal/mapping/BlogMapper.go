package mapper

import (
	dto2 "yp-blog-api/internal/dto"
	"yp-blog-api/internal/models"
)

type BlogMapper interface {
	BlogToBlogCardDto(blogs []models.Blog) []dto2.BlogCardDto
	BlogToBlogCardDtoSingle(blog models.Blog) dto2.BlogCardDto
	BlogToBlogDetailDto(blog models.Blog) dto2.BlogDetailDto
	ToTopAuthorDTO(results []interface{}) *dto2.TopAuthorDto
	ToTopAuthorDTOList(results [][]interface{}) []*dto2.TopAuthorDto
	CreateBlogDtoToBlog(dto dto2.BlogCreateRequestDto) models.Blog
	UpdateBlog(blog *models.Blog, dto dto2.BlogUpdateRequestDto)
	BlogDtoToBlogAdminDto(blogs []models.Blog) []dto2.BlogAdminDto
	BlogToRecentPostBlogDto(blog models.Blog) dto2.RecentPostBlogDto
}
