package response

import "net/http"

type Basic struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func SendOK(w http.ResponseWriter, data any) {
	SendJson(w, http.StatusOK, data)
}
