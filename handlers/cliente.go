package handlers

import (
	"encoding/json"
	"net/http"
	"jornada-backend/services"
)

func GetBasicClientInfoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Paso 1: Obtener token
		token, _, err := services.GetDolibarrToken("admin", "12345678") // <-- usa credenciales reales
		if err != nil {
			http.Error(w, "Error obteniendo token de Dolibarr: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Paso 2: Llamar al servicio con el token
		service := services.NewClientService()
		clientInfo, err := service.GetAllClientsBasicInfo(token)
		if err != nil {
			http.Error(w, "Error obteniendo info cliente: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Paso 3: Responder en JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(clientInfo)
	}
}
