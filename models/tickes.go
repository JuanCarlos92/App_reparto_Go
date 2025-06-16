package models

type TicketBasicInfo struct {
	ID        string `json:"id"`
	Ref       string `json:"ref"`
	Subject   string `json:"subject"`
	Status    string `json:"status"`
	DateOpen  int64  `json:"date_creation"` // o usa time.Time si haces parseo
}
