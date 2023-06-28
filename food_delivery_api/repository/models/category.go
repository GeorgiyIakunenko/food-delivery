package models

type Category struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Image       []byte `json:"image"`
	Description string `json:"description"`
}
