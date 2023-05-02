package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	client "github.com/dsauerbrun/cragcast/pkg/weather-client"
)

const (
	errorResponseTemplate = `{ "error": "%s" }`
)

type Controllers struct {
	cl client.WeatherClient
}

func (c *Controllers) GetForecast(w http.ResponseWriter, r *http.Request) {
	forcast, err := c.cl.GetForecast(40.0294122, -105.3223779)
	if err != nil {
		// TODO(joshrosso): need to be more thoughtful with the error
		// message we pass back in the body
		respondWithInternalServerError(w, err)
		return
	}

	forcastJSON, err := json.Marshal(forcast)
	if err != nil {
		// TODO(joshrosso): need to be more thoughtful with the error
		// message we pass back in the body
		respondWithInternalServerError(w, err)
		return
	}
	_, err = w.Write(forcastJSON)
	if err != nil {
		// TODO(joshrosso): need to be more thoughtful with the error
		// message we pass back in the body
		respondWithInternalServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", http.DetectContentType(forcastJSON))
}

func NewController() Controllers {
	return Controllers{cl: client.New()}
}

func respondWithInternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(errorResponseTemplate, err.Error())))
}
