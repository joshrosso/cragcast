package noaaclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type NoaaClient struct{}

func (nc *NoaaClient) GetForecast(xCoordinate float32, yCoordinate float32) (ForecastResponse, error) {
	// BOU weather office(boulder), 52X, 75Y this is boulder
	url := fmt.Sprintf("https://api.weather.gov/gridpoints/BOU/%f,%f/forecast/hourly", xCoordinate, yCoordinate)
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject ForecastResponse
	json.Unmarshal(responseData, &responseObject)
	return responseObject, nil
}

func GenerateNoaaClient() *NoaaClient {
	return &NoaaClient{}
}

type ForecastResponse struct {
	Properties struct {
		Elevation struct {
			UnitCode string  `json:"unitCode"`
			Value    float64 `json:"value"`
		} `json:"elevation"`
		ForecastGenerator string    `json:"forecastGenerator"`
		GeneratedAt       time.Time `json:"generatedAt"`
		Periods           []struct {
			Dewpoint struct {
				UnitCode string  `json:"unitCode"`
				Value    float64 `json:"value"`
			} `json:"dewpoint"`
			EndTime                    time.Time `json:"endTime"`
			Icon                       string    `json:"icon"`
			IsDaytime                  bool      `json:"isDaytime"`
			Name                       string    `json:"name"`
			Number                     int       `json:"number"`
			ProbabilityOfPrecipitation struct {
				UnitCode string `json:"unitCode"`
				Value    int    `json:"value"`
			} `json:"probabilityOfPrecipitation"`
			RelativeHumidity struct {
				UnitCode string `json:"unitCode"`
				Value    int    `json:"value"`
			} `json:"relativeHumidity"`
			ShortForecast    string    `json:"shortForecast"`
			StartTime        time.Time `json:"startTime"`
			Temperature      int       `json:"temperature"`
			TemperatureTrend any       `json:"temperatureTrend"`
			TemperatureUnit  string    `json:"temperatureUnit"`
			WindDirection    string    `json:"windDirection"`
			WindSpeed        string    `json:"windSpeed"`
		} `json:"periods"`
		Units      string    `json:"units"`
		UpdateTime time.Time `json:"updateTime"`
		Updated    time.Time `json:"updated"`
		ValidTimes string    `json:"validTimes"`
	} `json:"properties"`
	Type string `json:"type"`
}
