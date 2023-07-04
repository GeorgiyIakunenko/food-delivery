package response

import "net/http"

type Basic struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserResponse struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func SendOK(w http.ResponseWriter, data any) {
	SendJson(w, http.StatusOK, data)
}
