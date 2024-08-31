package dto

type BlogCardDto struct {
	Slug                 string        `json:"slug"`
	Thumbnail            string        `json:"thumbnail"`
	Summary              string        `json:"summary"`
	BlogTitle            string        `json:"blogTitle"`
	FormattedCountViewer string        `json:"formattedCountViewer"`
	MinRead              int           `json:"minRead"`
	Published            bool          `json:"published"`
	Author               AuthorCardDto `json:"author"`
	CreatedAt            string        `json:"createdAt"`
}
