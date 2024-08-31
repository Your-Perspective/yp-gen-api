package repositories

import (
	"yp-blog-api/models"

	"gorm.io/gorm"
)

type TagRepository interface {
	Create(tag *models.Tag) error
	Update(tag *models.Tag) error
	FindById(id int) (*models.Tag, error)
	FindByTitle(title string) (*models.Tag, error)
	FindAll() ([]models.Tag, error)
	Delete(tag *models.Tag) error
	FindAllById(ids []int) ([]models.Tag, error)
}

type tagRepositoryImpl struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepositoryImpl{db: db}
}

func (r *tagRepositoryImpl) Create(tag *models.Tag) error {
	return r.db.Create(tag).Error
}

func (r *tagRepositoryImpl) Update(tag *models.Tag) error {
	return r.db.Save(tag).Error
}

func (r *tagRepositoryImpl) FindById(id int) (*models.Tag, error) {
	var tag models.Tag
	err := r.db.First(&tag, id).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *tagRepositoryImpl) FindByTitle(title string) (*models.Tag, error) {
	var tag models.Tag
	err := r.db.Where("title = ?", title).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *tagRepositoryImpl) FindAll() ([]models.Tag, error) {
	var tags []models.Tag
	err := r.db.Find(&tags).Error
	return tags, err
}

func (r *tagRepositoryImpl) Delete(tag *models.Tag) error {
	return r.db.Delete(tag).Error
}

func (r *tagRepositoryImpl) FindAllById(ids []int) ([]models.Tag, error) {
	var tags []models.Tag
	err := r.db.Where("id IN ?", ids).Find(&tags).Error
	return tags, err
}
