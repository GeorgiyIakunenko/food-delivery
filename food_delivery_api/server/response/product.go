package response

type Product struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Supplier string `json:"supplier"`
	Category string `json:"category"`
}
