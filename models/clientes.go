package models

// clientes 
type ClientBasicInfo struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Address  string  `json:"address"`
	Zip      *string `json:"zip,omitempty"`
	Town     *string `json:"town,omitempty"`
	RegionID *string `json:"region_id,omitempty"`
	Phone    *string `json:"phone,omitempty"`
	Email    *string `json:"email,omitempty"`
}
