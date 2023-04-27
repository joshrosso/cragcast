package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// used for retrieval in the noaa-client
	weatherDotGovAPIEndpoint = "https://api.weather.gov"
)

type Elevation struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}

type Dewpoint struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}

type ProbabilityOfPrecipitation struct {
	UnitCode string `json:"unitCode"`
	Value    int    `json:"value"`
}

type RelativeHumidity struct {
	UnitCode string `json:"unitCode"`
	Value    int    `json:"value"`
}

type Period struct {
	Dewpoint                   Dewpoint                   `json:"dewpoint"`
	EndTime                    time.Time                  `json:"endTime"`
	Icon                       string                     `json:"icon"`
	IsDaytime                  bool                       `json:"isDaytime"`
	Name                       string                     `json:"name"`
	Number                     int                        `json:"number"`
	ProbabilityOfPrecipitation ProbabilityOfPrecipitation `json:"probabilityOfPrecipitation"`
	RelativeHumidity           RelativeHumidity           `json:"relativeHumidity"`
	ShortForecast              string                     `json:"shortForecast"`
	StartTime                  time.Time                  `json:"startTime"`
	Temperature                int                        `json:"temperature"`
	TemperatureTrend           interface{}                `json:"temperatureTrend"`
	TemperatureUnit            string                     `json:"temperatureUnit"`
	WindDirection              string                     `json:"windDirection"`
	WindSpeed                  string                     `json:"windSpeed"`
}

type Properties struct {
	Elevation         Elevation `json:"elevation"`
	ForecastGenerator string    `json:"forecastGenerator"`
	GeneratedAt       time.Time `json:"generatedAt"`
	Periods           []Period  `json:"periods"`
	Units             string    `json:"units"`
	UpdateTime        time.Time `json:"updateTime"`
	Updated           time.Time `json:"updated"`
	ValidTimes        string    `json:"validTimes"`
}

type ForecastResponse struct {
	Properties Properties `json:"properties"`
	Type       string     `json:"type"`
}

type WeatherClient interface {
	GetForecast(xCoordinate, yCoordinate float32) (*ForecastResponse, error)
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

func (nc *NoaaClient) GetForecast(xCoordinate float32, yCoordinate float32) (*ForecastResponse, error) {
	// BOU weather office(boulder), 52X, 75Y this is boulder
	url := fmt.Sprintf("%s/gridpoints/BOU/%f,%f/forecast/hourly",
		nc.APIEndpoint, xCoordinate, yCoordinate)
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
