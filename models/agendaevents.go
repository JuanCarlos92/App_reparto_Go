package models

type AgendaEvent struct {
	ClientID string `json:"socid"`  // ID del cliente
	DateUnix int64  `json:"datep"`  // Fecha de la reunión (timestamp Unix)
}

type EventResponse struct {
	ClientID    string `json:"client_id"`
	MeetingDate string `json:"meeting_date"`
}