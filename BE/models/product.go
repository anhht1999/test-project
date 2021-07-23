package models

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       int64   `json:"price"`
	Trademark   string  `json:"trade_mark"`
	Image       []Image `json:"image"`
	CategoryId  int     `json:"category_id"`
	Category    string  `json:"category"`
}
