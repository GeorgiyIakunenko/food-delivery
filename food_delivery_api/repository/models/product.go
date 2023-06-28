package models

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	SupplierID  int64   `json:"supplier_id"`
	CategoryID  int64   `json:"category_id"`
	Price       float64 `json:"price"`
	Image       []byte  `json:"image"`
	Description string  `json:"description"`
}
