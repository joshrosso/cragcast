package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dsauerbrun/cragcast/pkg/weather-client"
)

const (
	errorResponseTemplate = `{ "error": "%s" }`
)

type Controllers struct {
	cl client.WeatherClient
}

func (c *Controllers) GetForecast(w http.ResponseWriter, r *http.Request) {
	forcast, err := c.cl.GetForecast(52, 75)
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
