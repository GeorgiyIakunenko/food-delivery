package response

type OrderProduct struct {
	OrderID   int64    `json:"order_id"`
	ProductID int64    `json:"product_id"`
	Product   *Product `json:"product"`
	Quantity  int64    `json:"quantity"`
}
