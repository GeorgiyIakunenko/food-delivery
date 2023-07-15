package models

import "time"

type Order struct {
	ID            int64     `json:"id"`
	CustomerID    int64     `json:"customer_id"`
	TotalPrice    float64   `json:"total_price"`
	OrderStatus   string    `json:"order_status"`
	PaymentMethod string    `json:"payment_method"`
	CreatedAt     time.Time `json:"created_at"`
	Address       string    `json:"address"`
}
