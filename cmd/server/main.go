package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dsauerbrun/cragcast/noaaclient"
)

func main() {
	NoaaClient := noaaclient.GenerateNoaaClient()
	http.HandleFunc("/forecast/boulder", func(w http.ResponseWriter, r *http.Request) {

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
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
