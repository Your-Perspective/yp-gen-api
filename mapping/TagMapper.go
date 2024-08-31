package mapper

import (
	"yp-blog-api/dto"
	"yp-blog-api/models"
)

type TagMapper interface {
	TagToTagDto(tag *models.Tag) *dto.TagDto
	TagDtoToTag(tagDto dto.TagDto) *models.Tag
	TagsToTagDtos(tags []models.Tag) []*dto.TagDto
}

type tagMapperImpl struct{}

func NewTagMapper() TagMapper {
	return &tagMapperImpl{}
}

func (m *tagMapperImpl) TagToTagDto(tag *models.Tag) *dto.TagDto {
	return &dto.TagDto{
		ID:    int64(tag.ID),
		Title: tag.Title,
	}
}

func (m *tagMapperImpl) TagDtoToTag(tagDto dto.TagDto) *models.Tag {
	return &models.Tag{
		ID:    uint(tagDto.ID),
		Title: tagDto.Title,
	}
}

func (m *tagMapperImpl) TagsToTagDtos(tags []models.Tag) []*dto.TagDto {
	var tagDtos []*dto.TagDto
	for _, tag := range tags {
		tagDtos = append(tagDtos, m.TagToTagDto(&tag))
	}
	return tagDtos
}
