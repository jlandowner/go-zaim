package gozaim

// Currency 通貨
type Currency struct {
	CurrencyCode string `json:"currency_code"`
	Unit         string `json:"unit"`
	Name         string `json:"name"`
	Point        int    `json:"point"`
}
