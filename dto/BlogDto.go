package dto

import "time"

type BlogDto struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Slug        string    `json:"slug"`
	IsPinned    bool      `json:"is_pinned"`
	Thumbnail   string    `json:"thumbnail"`
	CountViewer int       `json:"count_viewer"`
	Summary     string    `json:"summary"`
	MinRead     int       `json:"min_read"`
	AuthorName  string    `json:"author_name"`
	Tags        []string  `json:"tags"`
	Categories  []string  `json:"categories"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BlogDetailDto struct {
	BlogDto
	ParentBlogTitle string `json:"parent_blog_title"`
}
