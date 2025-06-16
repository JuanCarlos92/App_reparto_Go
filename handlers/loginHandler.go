package handlers

import (
	"encoding/json"
	"fmt"
	"jornada-backend/models"
	"jornada-backend/services"
	"net/http"
)

// LoginHandler maneja las solicitudes de login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var loginData models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, "Error al procesar los datos", http.StatusBadRequest)
		return
	}

    token, _, err := services.GetDolibarrToken(loginData.Login, loginData.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener el token: %v", err), http.StatusUnauthorized)
		return
	}

	response := models.LoginResponse{
		DOLAPIKEY: token,
		Login:     loginData.Login,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

