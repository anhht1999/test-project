package models

type Order struct {
	Id          int64      `json:"id"`
	UserId      int64      `json:"user_id"`
	PhoneNumber string     `json:"phone_number"`
	Address     string     `json:"address"`
	TotalPrice  int64      `json:"total_price"`
	CreatedAt   string     `json:"created_at"`
	OrderNote   string     `json:"order_note"`
	Status      bool       `json:"status"`
	OrderItems  []OderItem `json:"order_items"`
}
