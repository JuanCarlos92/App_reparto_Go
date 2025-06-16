package models

type ProductBasicInfo struct {
	ID     string  `json:"id"`
	Ref    string  `json:"ref"`
	Label  string  `json:"label"`
	Price  string  `json:"price"`
	Weight string  `json:"weight"`
}
