package main

import (
	"log"
	"net/http"
	"time"
	"github.com/rs/cors"
	"sistemabackend/routes"
	"sistemabackend/utils"
)

func main() {
	// Inicializar la conexión a la base de datos (para verificar que funciona)
	db := utils.IniciarConexion()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error al obtener la instancia de DB:", err)
	}
	defer sqlDB.Close()

	// Cargar las rutas
	routes.LoadRouters()

	// Configuración CORS mejorada
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200", "http://127.0.0.1:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:           true, // Habilita logs detallados de CORS
	})

	// Configuración del servidor HTTP
	server := &http.Server{
		Addr:         ":8080",
		Handler:      corsHandler.Handler(http.DefaultServeMux),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("🚀 Servidor iniciado en http://localhost:8080")
	log.Println("📌 Endpoint de usuarios: GET http://localhost:8080/api/users")
	log.Fatal(server.ListenAndServe())
}