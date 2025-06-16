package handlers

import (
	"encoding/json"
	"net/http"
	"jornada-backend/services"
)

func GetBasicTicketInfoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, _, err := services.GetDolibarrToken("admin", "12345678")
		if err != nil {
			http.Error(w, "Error obteniendo token de Dolibarr: "+err.Error(), http.StatusInternalServerError)
			return
		}

		service := services.NewTicketService()
		tickets, err := service.GetAllTicketsBasicInfo(token)
		if err != nil {
			http.Error(w, "Error obteniendo tickets: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tickets)
	}
}
