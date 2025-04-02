package controllers

import (
	"sistemabackend/models"
	"sistemabackend/utils"
	"net/http"
	"encoding/json"
	"log"
)
// Metodo para obtener todos los usuarios (GET /api/users)
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := utils.IniciarConexion()
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		utils.HandleError(w, "Error al obtener los usuarios", http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(w, users, http.StatusOK)
}

// Metodo para obtener un usuario por su ID (GET /api/users/{id})
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	db := utils.IniciarConexion()
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		utils.HandleError(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}
	utils.RespondJSON(w, user, http.StatusOK)
}

// Metodo para crear un nuevo usuario (POST /api/users)
func CreateUser(w http.ResponseWriter, r *http.Request) {
    log.Println("Solicitud recibida en CreateUser")

    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        log.Printf("Error al parsear el JSON: %v\n", err)
        utils.HandleError(w, "Error al parsear los datos", http.StatusBadRequest)
        return
    }

    log.Printf("Datos recibidos: %+v\n", user)

    db := utils.IniciarConexion()
    if err := db.Create(&user).Error; err != nil {
        log.Printf("Error al insertar el usuario en la base de datos: %v\n", err)
        utils.HandleError(w, "Error al insertar el usuario", http.StatusInternalServerError)
        return
    }

    log.Printf("Usuario creado exitosamente: %+v\n", user)
    utils.RespondJSON(w, user, http.StatusCreated)
}

// Metodo para actualizar un usuario (PUT /api/users/{id})
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	db := utils.IniciarConexion()
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		utils.HandleError(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	var updatedUser models.User
	// Parsear los datos de la solicitud
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		utils.HandleError(w, "Error al parsear los datos", http.StatusBadRequest)
		return
	}

	// Actualizar los datos
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email

	if err := db.Save(&user).Error; err != nil {
		utils.HandleError(w, "Error al actualizar el usuario", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, user, http.StatusOK)
}

// Metodo para eliminar un usuario (DELETE /api/users/{id})
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	db := utils.IniciarConexion()
	if err := db.Delete(&models.User{}, id).Error; err != nil {
		utils.HandleError(w, "Error al borrar el usuario", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Respuesta vacía para indicar éxito
}