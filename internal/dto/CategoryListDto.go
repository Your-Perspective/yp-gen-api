package dto

// CategoryListDto represents the data transfer object for a category
type CategoryListDto struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
}
