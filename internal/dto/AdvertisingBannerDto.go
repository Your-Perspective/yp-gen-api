package dto

type AdvertisingBannerDto struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	ImageURL string `json:"imageUrl"`
	Link     string `json:"link"`
}
