package api

import (
	"log"
	"net/http"
)

func InstantiateRoutes() {
	myController := &Controllers{}

	http.HandleFunc("/forecast/boulder", myController.GetForecast)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
