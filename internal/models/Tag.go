package models

import "time"

type Tag struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	//Title     string    `gorm:"size:100;not null;unique"`
	Title     string    `gorm:"size:100;unique"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	IsDeleted bool      `gorm:"default:false" json:"isDeleted"`
	Blogs     []Blog    `gorm:"many2many:blog_tags;"`
}

func (Tag) TableName() string {
	return "tags"
}
