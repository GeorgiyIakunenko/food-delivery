package utils

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func MustGetIDFromVars(r *http.Request) (int, bool) {
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		return 0, false
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, false
	}

	return id, true
}
