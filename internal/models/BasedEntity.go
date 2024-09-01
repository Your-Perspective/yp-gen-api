package models

import (
	"time"

	"gorm.io/gorm"
)

// BasedEntity corresponds to the Java BasedEntity class
type BasedEntity struct {
	IsDeleted      bool      `gorm:"default:false" json:"isDeleted"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"createdAt"`
	LastModifiedAt time.Time `gorm:"autoUpdateTime" json:"lastModifiedAt"`
}

// BeforeCreate is a GORM hook that automatically sets the CreatedAt field
func (e *BasedEntity) BeforeCreate(*gorm.DB) (err error) {
	if e.CreatedAt.IsZero() {
		e.CreatedAt = time.Now()
	}
	return
}

// BeforeUpdate is a GORM hook that automatically sets the LastModifiedAt field
func (e *BasedEntity) BeforeUpdate(*gorm.DB) (err error) {
	e.LastModifiedAt = time.Now()
	return
}
