package models

type DetailOrder struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int    `json:"quantity"`
	Status   bool   `json:"status"`
}

type DetailUser struct {
	Id          int64   `json:"id"`
	CreatedAt   string  `json:"created_at"`
	Last_name   string  `json:"last_name"`
	First_name  string  `json:"first_name"`
	Email       string  `json:"email"`
	Address     string  `json:"address"`
	PhoneNumber string  `json:"phone_number"`
	OrderNote   string  `json:"order_note"`
	TotalPrice  float64 `json:"total_price"`
	Status      bool    `json:"status"`
}
