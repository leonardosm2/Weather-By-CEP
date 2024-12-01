package usecase

import (
	"errors"
	"testing"

	"github.com/leonardosm2/Weather-By-CEP/internal/adapters/api"
	"github.com/leonardosm2/Weather-By-CEP/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type LocationClientMock struct {
	mock.Mock
}

func (m *LocationClientMock) GetLocation(cep entity.CEP) (string, error) {
	args := m.Called(cep)
	if cep == "07085310" {
		return "Guarulhos", nil
	} else {
		return "", args.Error(1)
	}
}

type WeatherClientMock struct {
}

func (m *WeatherClientMock) GetWeather(city string) (float64, error) {
	return 28.5, nil
}

func TestGetTempUseCase(t *testing.T) {
	locationClient := &LocationClientMock{}
	locationClient.On("GetLocation", entity.CEP("07085310")).Return("Guarulhos", nil)
	locationClient.On("GetLocation", entity.CEP("07085311")).Return("", api.ErrNotFoundZipcode)
	locationClient.On("GetLocation", entity.CEP("0708531")).Return("", errors.New("bad request"))

	weatherClient := &WeatherClientMock{}

	getTemp := NewGetTempUseCase(locationClient, weatherClient)

	expected := TempOutputDTO{
		TempC: 28.5,
		TempF: 83.3,
		TempK: 301.5,
	}

	dto, err := getTemp.Execute("07085310")
	assert.Equal(t, expected, dto)
	assert.Nil(t, err)

	_, err = getTemp.Execute("07085311")
	assert.Equal(t, api.ErrNotFoundZipcode, err)

	_, err = getTemp.Execute("0708531")
	assert.Equal(t, entity.ErrInvalidZipcode, err)
}
