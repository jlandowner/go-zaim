package gozaim

// Account 口座情報
type Account struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Mode            string `json:"mode"`
	Sort            int    `json:"sort"`
	ParentAccountID int    `json:"parent_category_id"`
	Active          int    `json:"active"`
	Modified        string `json:"modified"`
}
