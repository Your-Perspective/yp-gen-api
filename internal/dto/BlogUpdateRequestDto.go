package dto

import (
	"github.com/go-playground/validator/v10"
)

// BlogUpdateRequestDto corresponds to the Java BlogUpdateRequestDto class
type BlogUpdateRequestDto struct {
	ID          int    `json:"id"`
	BlogTitle   string `json:"blogTitle" validate:"required,max=255"`
	Published   bool   `json:"published" validate:"required"`
	BlogContent string `json:"blogContent" validate:"required"`
	IsPin       bool   `json:"isPin" validate:"required"`
	Thumbnail   string `json:"thumbnail" validate:"omitempty,max=255"`
	Summary     string `json:"summary" validate:"omitempty,max=500"`
	MinRead     int    `json:"minRead" validate:"required,min=1"`
}

// Validate function to validate the BlogUpdateRequestDto struct
func (b *BlogUpdateRequestDto) Validate() error {
	validate := validator.New()
	return validate.Struct(b)
}
