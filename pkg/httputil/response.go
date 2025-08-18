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
	errResp := ErrorResponse{
		Code:    statusCode,
		Message: message,
	}

	WriteJSON(w, statusCode, errResp)
}
