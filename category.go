package gozaim

// Category カテゴリ
type Category struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Mode             string `json:"mode"`
	Sort             int    `json:"sort"`
	ParentCategoryID int    `json:"parent_category_id"`
	Active           int    `json:"active"`
	Modified         string `json:"modified"`
}
