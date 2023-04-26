package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dsauerbrun/cragcast/noaaclient"
)

type ForecastResponse struct {
	_Context []any `json:"@context"`
	Geometry struct {
		Coordinates [][][]float64 `json:"coordinates"`
		Type        string        `json:"type"`
	} `json:"geometry"`
	Properties struct {
		Elevation struct {
			UnitCode string  `json:"unitCode"`
			Value    float64 `json:"value"`
		} `json:"elevation"`
		ForecastGenerator string    `json:"forecastGenerator"`
		GeneratedAt       time.Time `json:"generatedAt"`
		Periods           []struct {
			DetailedForecast string `json:"detailedForecast"`
			Dewpoint         struct {
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
				Value    *int   `json:"value"`
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

func main() {
	NoaaClient := noaaclient.GenerateNoaaClient()
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {

		gotForecaset, _ := NoaaClient.GetForecast(23, 23)
		fmt.Println("got forecast", gotForecaset)
		response, err := http.Get("https://api.weather.gov/gridpoints/TOP/31,80/forecast")

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
		fmt.Fprintf(w, responseObject.Geometry.Type)
		//fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
