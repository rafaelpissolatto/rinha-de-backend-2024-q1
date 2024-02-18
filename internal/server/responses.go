package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON sends a JSON response to the client
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// to avoid error when update user
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, fmt.Sprintf("Error to encode JSON: %s", err.Error()), http.StatusInternalServerError)
		}
	}
}

// Error sends a JSON error response to the client
func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
