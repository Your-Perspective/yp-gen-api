package dto

import "github.com/go-playground/validator/v10"

// TagDto corresponds to the Java TagDto record
type TagDto struct {
	Title string `json:"title" validate:"required"` // Equivalent to @NotBlank
	ID    int64  `json:"id"`
}

// Validate function for TagDto to ensure Title is not blank
func (t *TagDto) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
