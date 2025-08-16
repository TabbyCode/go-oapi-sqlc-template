package httputil

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/xurvan/go-oapi-sqlc-template/internal/gen/oapi"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(statusCode)

		_, err = w.Write(jsonData)
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
	} else {
		w.WriteHeader(statusCode)
	}
}

func WriteError(w http.ResponseWriter, statusCode int, message string) {
	errResp := oapi.Error{
		Code:    statusCode,
		Message: message,
	}

	WriteJSON(w, statusCode, errResp)
}
