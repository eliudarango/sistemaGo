package routes

import (
	"encoding/json"
	"net/http"
	"sistemabackend/models"
	"sistemabackend/utils"
)

func LoadRouters() {
	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		db := utils.IniciarConexion()

		var users []models.User
		result := db.Find(&users)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})
}