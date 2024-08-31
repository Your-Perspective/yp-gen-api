package service

import (
	"errors"
	"yp-blog-api/internal/dto"
	"yp-blog-api/internal/mapping"
	_ "yp-blog-api/internal/mapping"
	"yp-blog-api/internal/repository"
)

type TagService interface {
	CreateTag(tagDto dto.TagDto) (*dto.TagDto, error)
	UpdateTag(id int, tagDetails dto.TagDto) (*dto.TagDto, error)
	GetAllTags() ([]*dto.TagDto, error)
	GetTagById(id int) (*dto.TagDto, error)
	DeleteTag(id int) error
}

type tagServiceImpl struct {
	tagRepo   repositories.TagRepository
	tagMapper mapper.TagMapper
}

func NewTagService(tagRepo repositories.TagRepository, tagMapper mapper.TagMapper) TagService {
	return &tagServiceImpl{
		tagRepo:   tagRepo,
		tagMapper: tagMapper,
	}
}

func (s *tagServiceImpl) CreateTag(tagDto dto.TagDto) (*dto.TagDto, error) {
	if tagDto.Title == "" {
		return nil, errors.New("tag name cannot be null or empty")
	}

	existingTag, err := s.tagRepo.FindByTitle(tagDto.Title)
	if err == nil && existingTag != nil {
		return nil, errors.New("tag with name '" + tagDto.Title + "' already exists")
	}

	tag := s.tagMapper.TagDtoToTag(tagDto)
	err = s.tagRepo.Create(tag)
	if err != nil {
		return nil, err
	}

	return s.tagMapper.TagToTagDto(tag), nil
}

func (s *tagServiceImpl) UpdateTag(id int, tagDetails dto.TagDto) (*dto.TagDto, error) {
	if tagDetails.Title == "" {
		return nil, errors.New("tag name cannot be null or empty")
	}

	tag, err := s.tagRepo.FindById(id)
	if err != nil {
		return nil, errors.New("tag not found: " + string(rune(id)))
	}

	existingTag, err := s.tagRepo.FindByTitle(tagDetails.Title)
	if err == nil && existingTag != nil && existingTag.ID != uint(id) { // Convert id to uint
		return nil, errors.New("tag with name '" + tagDetails.Title + "' already exists")
	}

	tag.Title = tagDetails.Title
	err = s.tagRepo.Update(tag)
	if err != nil {
		return nil, err
	}

	return s.tagMapper.TagToTagDto(tag), nil
}

func (s *tagServiceImpl) GetAllTags() ([]*dto.TagDto, error) {
	tags, err := s.tagRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return s.tagMapper.TagsToTagDtos(tags), nil
}

func (s *tagServiceImpl) GetTagById(id int) (*dto.TagDto, error) {
	tag, err := s.tagRepo.FindById(id)
	if err != nil {
		return nil, errors.New("tag not found: " + string(rune(id)))
	}
	return s.tagMapper.TagToTagDto(tag), nil
}

func (s *tagServiceImpl) DeleteTag(id int) error {
	tag, err := s.tagRepo.FindById(id)
	if err != nil {
		return errors.New("tag not found: " + string(rune(id)))
	}
	return s.tagRepo.Delete(tag)
}
