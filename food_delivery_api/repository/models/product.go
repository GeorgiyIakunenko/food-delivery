package models

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	CategoryID  int64     `json:"category_id"`
	Category    *Category `json:"category"`
	SupplierID  int64     `json:"supplier_id"`
	Supplier    *Supplier `json:"supplier"`
	Image       []byte    `json:"image"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
}
