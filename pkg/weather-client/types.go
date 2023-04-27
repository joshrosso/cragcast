package client

import "time"

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

type PointResponse struct {
	Properties struct {
		Cwa                 string `json:"cwa"`
		Forecast            string `json:"forecast"`
		ForecastGridData    string `json:"forecastGridData"`
		ForecastHourly      string `json:"forecastHourly"`
		ForecastOffice      string `json:"forecastOffice"`
		ForecastZone        string `json:"forecastZone"`
		GridID              string `json:"gridId"`
		GridX               int    `json:"gridX"`
		GridY               int    `json:"gridY"`
		ObservationStations string `json:"observationStations"`
		RadarStation        string `json:"radarStation"`
		TimeZone            string `json:"timeZone"`
	} `json:"properties"`
	Type string `json:"type"`
}
