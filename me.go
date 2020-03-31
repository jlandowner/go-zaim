package gozaim

// Me ユーザ情報
type Me struct {
	ID              int    `json:"id"`
	Login           string `json:"login"`
	Name            string `json:"name"`
	InputCount      int    `json:"input_count"`
	DayCount        int    `json:"day_count"`
	RepeatCount     int    `json:"repeat_count"`
	Day             int    `json:"day"`
	Week            int    `json:"week"`
	Month           int    `json:"month"`
	CurrencyCode    string `json:"currency_code"`
	ProfileImageURL string `json:"profile_image_url"`
	CoverImageURL   string `json:"cover_image_url"`
	ProfileModified string `json:"profile_modified"`
}
