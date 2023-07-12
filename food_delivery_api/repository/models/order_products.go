package models

type OrderProduct struct {
	ID        int64    `json:"id"`
	OrderID   int64    `json:"order_id"`
	ProductID int64    `json:"product_id"`
	Product   *Product `json:"product"`
	Quantity  int64    `json:"quantity"`
}
