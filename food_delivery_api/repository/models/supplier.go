package models

type Supplier struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	CategoryID int64  `json:"category_id"`
	Image      []byte `json:"image"`
	AddressID  int64  `json:"address_id"`
	OpenTime   string `json:"open_time"`
	CloseTime  string `json:"close_time"`
}
