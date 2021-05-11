package response

import (
	"encoding/json"
	"net/http"
)

// SuccessResponse attempts to send a success response
func SuccessResponse(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(body)
}

// ErrorResponse attempts to send a error response
func ErrorResponse(w http.ResponseWriter, code int, message string) {
	var error Error
	error.Error = message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(error)
}
