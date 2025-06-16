package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"jornada-backend/models"
)

type AgendaService struct {
	client      *http.Client
	dolibarrURL string
}

func NewAgendaService() *AgendaService {
	return &AgendaService{
		client:      &http.Client{Timeout: 10 * time.Second},
		dolibarrURL: "http://localhost:8080/dolibarr/api/index.php/agendaevents", // Ajusta al endpoint real
	}
}

type dolibarrEvent struct {
	ClientID    string `json:"fk_soc"`
	DateUnix    int64  `json:"date_creation"` // o la propiedad que represente fecha (timestamp)
}

func (s *AgendaService) GetEvents(token string) ([]models.AgendaEvent, error) {
	req, err := http.NewRequest("GET", s.dolibarrURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("DOLAPIKEY", token)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error al obtener eventos desde Dolibarr: status %d, body %s", resp.StatusCode, string(bodyBytes))
	}

	var raw []dolibarrEvent
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	var events []models.AgendaEvent
	for _, e := range raw {
		events = append(events, models.AgendaEvent{
			ClientID: e.ClientID,
			DateUnix: e.DateUnix,
		})
	}

	return events, nil
}

