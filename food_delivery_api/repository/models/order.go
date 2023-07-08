package models

import "time"

type Order struct {
	ID            int64     `json:"id"`
	CustomerID    int64     `json:"customer_id"`
	SupplierID    int64     `json:"supplier_id"`
	OrderStatus   string    `json:"order_status"`
	OrderDatetime time.Time `json:"order_datetime"`
	Address       string    `json:"address"`
}
