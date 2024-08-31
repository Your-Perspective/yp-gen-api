package dto

import (
	"github.com/go-playground/validator/v10"
)

// BlogCreateRequestDto corresponds to the Java BlogCreateRequestDto class
type BlogCreateRequestDto struct {
	ID          int    `json:"id"`
	BlogTitle   string `json:"blogTitle" validate:"required,max=500"`
	Published   bool   `json:"published" validate:"required"`
	BlogContent string `json:"blogContent" validate:"required"`
	Slug        string `json:"slug"`
	IsPin       bool   `json:"isPin" validate:"required"`
	Thumbnail   string `json:"thumbnail" validate:"omitempty,max=255"`
	Summary     string `json:"summary" validate:"omitempty,max=500"`
	MinRead     int    `json:"minRead" validate:"required,min=1"`
	CategoryIds []int  `json:"categoryIds"`
	Tags        []int  `json:"tags"`
}

// Validate function to validate the BlogCreateRequestDto struct
func (b *BlogCreateRequestDto) Validate() error {
	validate := validator.New()
	return validate.Struct(b)
}
