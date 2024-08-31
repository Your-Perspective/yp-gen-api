package models

import "time"

type User struct {
	Email             string `gorm:"size:64;not null;unique"`
	ID                uint   `gorm:"primaryKey;autoIncrement"`
	UserName          string `gorm:"type:text;not null;default:'Unknown'"`         // Provide a default value
	Password          string `gorm:"size:256;not null;default:'default_password'"` // Provide a default value
	ResetToken        string `gorm:"size:256"`
	Top3Count         byte
	Bio               string    `gorm:"size:256"`
	About             string    `gorm:"size:500"`
	ConfirmationToken string    `gorm:"size:256"`
	IsVerified        bool      `gorm:"default:false"`
	VerifiedByAdmin   bool      `gorm:"default:false"`
	ProfileImage      string    `gorm:"size:256"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (User) TableName() string {
	return "users"
}
