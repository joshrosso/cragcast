package api

import (
	"fmt"
	"log"
	"net/http"
)

const (
	// DefaultPort is the Default setting a port is set to for the server
	// when not otherwise specified. It is publicly available for other
	// clients to lookup, and potentially use.
	DefaultPort = 8081
)

type StartOptions struct {
	Port     int
	BindHost string
}

func Start(opts StartOptions) {
	myController := &Controllers{}
	http.HandleFunc("/forecast/boulder", myController.GetForecast)

	if opts.Port == 0 {
		opts.Port = DefaultPort
	}
	address := fmt.Sprintf("%s:%d", opts.BindHost, opts.Port)
	log.Printf("serving cragcast API at %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
