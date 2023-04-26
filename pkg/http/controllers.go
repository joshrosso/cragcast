package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/dsauerbrun/cragcast/pkg/weather-client"
)

type Controllers struct{}

func (c *Controllers) GetForecast(w http.ResponseWriter, r *http.Request) {
	NoaaClient := client.New()
	gotForecast, err := NoaaClient.GetForecast(52, 75)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	marshalledforecast, err := json.Marshal(gotForecast)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(marshalledforecast))
}
