package response

import "net/http"

type Basic struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserResponse struct {
	ID        int64           `json:"id"`
	FirstName string          `json:"firstName"`
	LastName  string          `json:"lastName"`
	Username  string          `json:"username"`
	Age       int             `json:"age"`
	Email     string          `json:"email"`
	Phone     string          `json:"phone"`
	Address   AddressResponse `json:"address"`
}

type AddressResponse struct {
	City         string `json:"city"`
	PostalCode   int    `json:"postal_code"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	Country      string `json:"country"`
}

func SendOK(w http.ResponseWriter, data any) {
	SendJson(w, http.StatusOK, data)
}
