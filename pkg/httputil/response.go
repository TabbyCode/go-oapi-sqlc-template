package httputil

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Printf("Error encoding response: %v", err)
		}
	}
}

func WriteError(w http.ResponseWriter, statusCode int, message string) {
	errResp := ErrorResponse{
		Code:    statusCode,
		Message: message,
	}

	WriteJSON(w, statusCode, errResp)
}
