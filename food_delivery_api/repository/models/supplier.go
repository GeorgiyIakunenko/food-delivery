package models

type Supplier struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Image     string `json:"image"`
	Address   string `json:"address"`
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
}
