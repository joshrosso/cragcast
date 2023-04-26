package noaaclient

import "fmt"

type NoaaClient struct {
	Apikey string
}

func (nc *NoaaClient) GetForecast(lat float32, lng float32) (bool, error) {

	fmt.Println("getting forecast")

	return true, nil
}

func GenerateNoaaClient() *NoaaClient {
	return &NoaaClient{
		Apikey: "hey",
	}
}
