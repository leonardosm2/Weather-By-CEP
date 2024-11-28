package web

import (
	"encoding/json"
	"net/http"

	"github.com/leonardosm2/Weather-By-CEP/internal/entity"
	"github.com/leonardosm2/Weather-By-CEP/internal/usecase"
)

type WebTempHandler struct {
	LocationClient entity.LocationClientInterface
	WeatherClient  entity.WeatherClientInterface
}

func NewWebTempHandler(
	locationClient entity.LocationClientInterface,
	weatherClient entity.WeatherClientInterface,
) *WebTempHandler {
	return &WebTempHandler{
		LocationClient: locationClient,
		WeatherClient:  weatherClient,
	}
}

func (h *WebTempHandler) Get(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	input := queryParams.Get("CEP")
	getTemp := usecase.NewGetTempUseCase(h.LocationClient, h.WeatherClient)
	output, err := getTemp.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}