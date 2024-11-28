package main

import (
	"fmt"
	"net/http"

	"github.com/leonardosm2/Weather-By-CEP/internal/adapters/api"
	"github.com/leonardosm2/Weather-By-CEP/internal/infra/web"
	"github.com/leonardosm2/Weather-By-CEP/internal/infra/web/webserver"
)

func main() {
	locationClient := api.NewLocationClient("https://viacep.com.br/ws/@CEP/json/")
	weatherClient := api.NewWeatherClient("http://api.weatherapi.com/v1/current.json?key=@APIKEY&q=bulk", "1ba50f84ca424deba5f233704242711")

	webserver := webserver.NewWebServer(":8080")
	webTempHandler := web.NewWebTempHandler(locationClient, weatherClient)
	webserver.AddHandler("/temp", http.MethodGet, webTempHandler.Get)
	fmt.Println("Starting web server on port 8080")
	webserver.Start()
}
