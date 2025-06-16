package services

import (
    "encoding/json"
    "errors"
    "io"
    "log"
    "net/http"
    "time"
    "jornada-backend/models"
)

type ClientService struct {
    client      *http.Client
    dolibarrURL string
}

func NewClientService() *ClientService {
    return &ClientService{
        client:      &http.Client{Timeout: 10 * time.Second},
        dolibarrURL: "http://localhost:8080/dolibarr/api/index.php/thirdparties",
    }
}

type dolibarrClient struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
    Address string `json:"address"`
    Zip     string `json:"zip"`
    Town    string `json:"town"`
    Region  string `json:"region_id"`
    Phone   string `json:"phone"`
    Email   string `json:"email"`
}

func (s *ClientService) GetAllClientsBasicInfo(token string) ([]models.ClientBasicInfo, error) {
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

    log.Printf("Código de respuesta Dolibarr: %d", resp.StatusCode)

    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    log.Printf("Cuerpo de respuesta Dolibarr: %s", string(bodyBytes))

    if resp.StatusCode != http.StatusOK {
        return nil, errors.New("error al obtener datos de clientes desde Dolibarr")
    }

    var clients []dolibarrClient
    if err := json.Unmarshal(bodyBytes, &clients); err != nil {
        return nil, err
    }

    var result []models.ClientBasicInfo
    for _, c := range clients {
        client := models.ClientBasicInfo{
            ID:       c.ID,
            Name:     c.Name,
            Address:  c.Address,
            Zip:      &c.Zip,
            Town:     &c.Town,
            RegionID: &c.Region,
            Phone:    &c.Phone,
            Email:    &c.Email,
        }
        result = append(result, client)
    }

    return result, nil
}
