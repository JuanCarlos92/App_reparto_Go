package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"jornada-backend/services"
)

// GetEventsHandler maneja la petición GET /events y devuelve varios eventos desde Dolibarr
func GetEventsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtén el token (usa tus credenciales reales)
		token, _, err := services.GetDolibarrToken("admin", "12345678")
		if err != nil {
			http.Error(w, "Error obteniendo token de Dolibarr: "+err.Error(), http.StatusInternalServerError)
			return
		}

		service := services.NewAgendaService()

		events, err := service.GetEvents(token)
		if err != nil {
			http.Error(w, "Error obteniendo eventos: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Mapea los eventos para la respuesta JSON, con formato de fecha legible
		type EventResponse struct {
			ID          string `json:"id"`
			ClientID    string `json:"client_id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			MeetingDate string `json:"meeting_date"`
		}

		var response []EventResponse
		for _, e := range events {
			response = append(response, EventResponse{
				ClientID:    e.ClientID,
				MeetingDate: time.Unix(e.DateUnix, 0).UTC().Format(time.RFC3339),
			})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
