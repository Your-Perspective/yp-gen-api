package models

import "time"

type Blog struct {
	ID          uint       `gorm:"primaryKey"`
	BlogTitle   string     `gorm:"type:varchar(256);not null"`
	Published   bool       `gorm:"default:false"`
	BlogContent string     `gorm:"type:text;not null"`
	Slug        string     `gorm:"type:varchar(256);not null;unique"`
	IsPin       bool       `gorm:"default:false"`
	Thumbnail   string     `gorm:"type:varchar(256)"`
	CountViewer int        `gorm:"type:int"`
	Summary     string     `gorm:"type:text"`
	MinRead     int        `gorm:"type:tinyint"`
	ParentID    *uint      `gorm:"index"`
	Parent      *Blog      `gorm:"foreignKey:ParentID"`
	AuthorID    uint       `gorm:"index"`
	Author      User       `gorm:"foreignKey:AuthorID"`
	Tags        []Tag      `gorm:"many2many:blog_tags;"`
	Categories  []Category `gorm:"many2many:blog_categories;"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (Blog) TableName() string {
	return "blogs"
}
