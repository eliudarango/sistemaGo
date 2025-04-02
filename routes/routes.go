package routes

import (
	"encoding/json"
	"net/http"
	"sistemabackend/controllers"
	"sistemabackend/models"
	"sistemabackend/utils"
)

func LoadRouters() {
	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet: // Manejar solicitudes GET (obtener todos los usuarios o uno por ID)
			id := r.URL.Query().Get("id")
			if id != "" {
				controllers.GetUser(w, r) // Obtener un usuario por ID
			} else {
				db := utils.IniciarConexion()

				var users []models.User
				result := db.Find(&users)
				if result.Error != nil {
					http.Error(w, result.Error.Error(), http.StatusInternalServerError)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(users)
			}

		case http.MethodPost: // Manejar solicitudes POST (crear un usuario)
			controllers.CreateUser(w, r)

		case http.MethodPut: // Manejar solicitudes PUT (actualizar un usuario)
			controllers.UpdateUser(w, r)

		case http.MethodDelete: // Manejar solicitudes DELETE (eliminar un usuario)
			controllers.DeleteUser(w, r)

		default: // Método no permitido
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
}
