package dto

// BlogAdminDto corresponds to the Java BlogAdminDto record
type BlogAdminDto struct {
	ID          int           `json:"id"`
	BlogTitle   string        `json:"blogTitle"`
	Published   bool          `json:"published"`
	Slug        string        `json:"slug"`
	IsPin       bool          `json:"isPin"`
	Thumbnail   string        `json:"thumbnail"`
	CountViewer int           `json:"countViewer"`
	Summary     string        `json:"summary"`
	MinRead     int           `json:"minRead"`
	Author      UserDto       `json:"author"`
	Tags        []TagDto      `json:"tags"`
	Categories  []CategoryDto `json:"categories"`
}
