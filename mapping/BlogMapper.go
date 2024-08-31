package mapper

import (
	"yp-blog-api/dto"
	"yp-blog-api/models"
)

type BlogMapper interface {
	BlogToBlogCardDto(blogs []models.Blog) []dto.BlogCardDto
	BlogToBlogCardDtoSingle(blog models.Blog) dto.BlogCardDto
	BlogToBlogDetailDto(blog models.Blog) dto.BlogDetailDto
	BlogToRecentPostBlogDto(blog models.Blog) dto.RecentPostBlogDto
	ToTopAuthorDTO(results []interface{}) *dto.TopAuthorDto
	ToTopAuthorDTOList(results [][]interface{}) []*dto.TopAuthorDto
	CreateBlogDtoToBlog(dto dto.BlogCreateRequestDto) models.Blog
	UpdateBlog(blog *models.Blog, dto dto.BlogUpdateRequestDto)
	BlogDtoToBlogAdminDto(blogs []models.Blog) []dto.BlogAdminDto
}
