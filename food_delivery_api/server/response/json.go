package response

import (
	"encoding/json"
	"net/http"
)

func SendJson(w http.ResponseWriter, code int, data any) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		SendServerError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonData)
}
