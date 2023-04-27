package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	// used for retrieval in the noaa-client
	weatherDotGovAPIEndpoint = "https://api.weather.gov"
)

type WeatherClient interface {
	GetForecast(lat, lng float64) (*ForecastResponse, error)
}

type NoaaClient struct {
	APIEndpoint string
}

// New returns an instance of WeatherClient, which can be used to access weather data.
func New() WeatherClient {
	// TODO(joshrosso): This abstraction atop NoaaClient generation sits
	// here to support future weather providers or even the ability to mix
	// calls from various providers should we need to.
	return newNoaaClient()
}

func (nc *NoaaClient) getNoaaPointInfo(lat float64, lng float64) (xCoord float64, yCoord float64, stationCode string, err error) {
	// cache this in the future as it will basically never change
	pointUrl := fmt.Sprintf("%s/points/%f,%f", nc.APIEndpoint, lat, lng)
	response, err := http.Get(pointUrl)
	if err != nil {
		return 0, 0, "", err
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, 0, "", err
	}

	pointData := &PointResponse{}
	err = json.Unmarshal(responseBody, pointData)
	if err != nil {
		return 0, 0, "", err
	}

	return float64(pointData.Properties.GridX), float64(pointData.Properties.GridY), pointData.Properties.Cwa, nil
}

func (nc *NoaaClient) GetForecast(lat float64, lng float64) (*ForecastResponse, error) {
	// 40.0294122,-105.3223779 is lat,lng for boulder
	xCoord, yCoord, stationCode, err := nc.getNoaaPointInfo(lat, lng)
	if err != nil {
		return nil, err
	}

	// BOU weather office(boulder), 52X, 75Y is boulder
	url := fmt.Sprintf("%s/gridpoints/%s/%f,%f/forecast/hourly",
		nc.APIEndpoint, stationCode, xCoord, yCoord)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	forecastData := &ForecastResponse{}
	err = json.Unmarshal(responseBody, forecastData)
	if err != nil {
		return nil, err
	}
	return forecastData, nil
}

func newNoaaClient() *NoaaClient {
	return &NoaaClient{
		APIEndpoint: weatherDotGovAPIEndpoint,
	}
}
