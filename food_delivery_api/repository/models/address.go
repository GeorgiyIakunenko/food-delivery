package models

type Address struct {
	ID           int64  `json:"id"`
	City         string `json:"city"`
	PostalCode   int    `json:"postal_code"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	Country      string `json:"country"`
}
