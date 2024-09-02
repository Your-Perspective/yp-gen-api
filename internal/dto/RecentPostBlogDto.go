package dto

type RecentPostBlogDto struct {
	BlogTitle string  `json:"blogTitle"`
	Slug      string  `json:"slug"`
	TimeAgo   string  `json:"timeAgo"`
	Author    UserDto `json:"author"`
}
