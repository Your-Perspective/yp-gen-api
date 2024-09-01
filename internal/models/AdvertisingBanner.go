package models

import (
	"time"
)

// AdvertisingBanner corresponds to the Java AdvertisingBanner entity
type AdvertisingBanner struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"type:varchar(255)" json:"title"`
	ImageURL  string    `gorm:"type:varchar(255)" json:"imageUrl"`
	Link      string    `gorm:"type:varchar(255)" json:"link"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	IsDeleted bool      `gorm:"default:false" json:"isDeleted"`
}
