package response

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Supplier    string  `json:"supplier"`
	Category    string  `json:"category"`
	SupplierID  int64   `json:"supplier_id"`
	CategoryID  int64   `json:"category_id"`
}
