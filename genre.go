package gozaim

// Genre ジャンル
type Genre struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Sort          int    `json:"sort"`
	Active        int    `json:"active"`
	CategoryID    int    `json:"category_id"`
	ParentGenreID int    `json:"parent_genre_id"`
	Modified      string `json:"modified"`
	LocalID       int    `json:"local_id"`
}
