package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"jornada-backend/models"
)

type ProductService struct {
	client      *http.Client
	dolibarrURL string
}

func NewProductService() *ProductService {
	return &ProductService{
		client:      &http.Client{Timeout: 10 * time.Second},
		dolibarrURL: "http://localhost:8080/dolibarr/api/index.php/products",
	}
}

type dolibarrProduct struct {
	ID     string `json:"id"`
	Ref    string `json:"ref"`
	Label  string `json:"label"`
	Price  string `json:"price"`
	Weight string `json:"weight"`
}

func (s *ProductService) GetAllProductsBasicInfo(token string) ([]models.ProductBasicInfo, error) {
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
		return nil, errors.New("error al obtener productos desde Dolibarr")
	}

	var raw []dolibarrProduct
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	var result []models.ProductBasicInfo
	for _, p := range raw {
		result = append(result, models.ProductBasicInfo{
			ID:     p.ID,
			Ref:    p.Ref,
			Label:  p.Label,
			Price:  p.Price,
			Weight: p.Weight,
		})
	}
	return result, nil
}
