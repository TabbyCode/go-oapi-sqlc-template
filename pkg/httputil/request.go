package httputil

import (
	"encoding/json"
	"net/http"
)

func ReadJSON(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid request payload")
		return false
	}
	return true
}
