package utils

import (
	"encoding/json"
	"net/http"
)

// Función para enviar respuestas JSON
func RespondJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
