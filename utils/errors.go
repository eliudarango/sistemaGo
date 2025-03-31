package utils

import (
	"log"
	"net/http"
)

// Función para manejar errores y responder con el código y el mensaje
func HandleError(w http.ResponseWriter, message string, statusCode int) {
	log.Println(message)
	http.Error(w, message, statusCode)
}
