package response

type Supplier struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	Image     []byte `json:"image"`
	Address   string `json:"address"`
	OpenTime  string `json:"openTime"`
	CloseTime string `json:"closeTime"`
}
