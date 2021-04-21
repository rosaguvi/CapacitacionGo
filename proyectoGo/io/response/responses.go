package response

import (
	"net/http"
	"encoding/json"
)

func Json(response interface{}, status int, w http.ResponseWriter) {
	sendResponse(response, status, w)
}

func Error(message string, status int, w http.ResponseWriter) {
	sendResponse(message, status, w)
}

func sendResponse(response interface{}, status int, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}