package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"jornada-backend/models"
)

type TicketService struct {
	client      *http.Client
	dolibarrURL string
}

func NewTicketService() *TicketService {
	return &TicketService{
		client:      &http.Client{Timeout: 10 * time.Second},
		dolibarrURL: "http://localhost:8080/dolibarr/api/index.php/tickets",
	}
}

type dolibarrTicket struct {
	ID           string `json:"id"`
	Ref          string `json:"ref"`
	Subject      string `json:"subject"`
	Status       string `json:"status"`
	DateCreation int64  `json:"date_creation"`
}

func (s *TicketService) GetAllTicketsBasicInfo(token string) ([]models.TicketBasicInfo, error) {
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
		return nil, errors.New("error al obtener datos de tickets desde Dolibarr")
	}

	var rawTickets []dolibarrTicket
	if err := json.NewDecoder(resp.Body).Decode(&rawTickets); err != nil {
		return nil, err
	}

	var result []models.TicketBasicInfo
	for _, t := range rawTickets {
		result = append(result, models.TicketBasicInfo{
			ID:       t.ID,
			Ref:      t.Ref,
			Subject:  t.Subject,
			Status:   t.Status,
			DateOpen: t.DateCreation,
		})
	}

	return result, nil
}
