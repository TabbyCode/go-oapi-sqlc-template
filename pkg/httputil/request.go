package httputil

import (
	"encoding/json"
	"net/http"
)

func ReadJSON(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid request payload")
		return false
	}

	return true
}
