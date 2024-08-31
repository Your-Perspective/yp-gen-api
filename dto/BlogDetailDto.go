package dto

// BlogDetailDto corresponds to the Java BlogDetailDto record
type BlogDetailDto struct {
	Slug                 string              `json:"slug"`
	BlogContent          string              `json:"blogContent"`
	Summary              string              `json:"summary"`
	Thumbnail            string              `json:"thumbnail"`
	BlogTitle            string              `json:"blogTitle"`
	FormattedCountViewer string              `json:"formattedCountViewer"`
	MinRead              int                 `json:"minRead"`
	Published            bool                `json:"published"`
	Author               AuthorCardDetailDto `json:"author"`
	CreatedAt            string              `json:"createdAt"`
	LastModifiedTimeAgo  string              `json:"lastModifiedTimeAgo"`
	Categories           []CategoryDto       `json:"categories"`
	Tags                 []TagDto            `json:"tags"`
}
