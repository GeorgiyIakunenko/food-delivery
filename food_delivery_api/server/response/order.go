package response

type Order struct {
	ID           int64     `json:"id"`
	Customer     User      `json:"customer"`
	Supplier     Supplier  `json:"supplier"`
	Products     []Product `json:"products"`
	OrderStatus  string    `json:"order_status"`
	OrderDate    string    `json:"order_date"`
	OrderAddress string    `json:"order_address"`
}
