package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/leonardosm2/Weather-By-CEP/internal/entity"
)

type LocationClient struct {
	BaseURL string
}

type LocationResponse struct {
	Localidade string `json:"localidade"`
}

func NewLocationClient(url string) *LocationClient {
	return &LocationClient{BaseURL: url}
}

func (l *LocationClient) GetLocation(cep entity.CEP) (string, error) {
	url := strings.ReplaceAll(l.BaseURL, "@CEP", string(cep))

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var location LocationResponse
	err = json.Unmarshal(body, &location)
	if err != nil {
		return "", err
	}

	return location.Localidade, nil
}
