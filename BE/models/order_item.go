package models

type OderItem struct {
	// OderId    int64 `json:"order_id"`
	ProductId int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}
