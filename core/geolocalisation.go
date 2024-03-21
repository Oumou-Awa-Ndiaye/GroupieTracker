package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const apiKey = "YOUR_API_KEY"

type GeocodeResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}

func main() {
	concerts := map[string]string{
		"Concert 1": "Berlin, Germany",
		"Concert 2": "Paris, France",
		"Concert 3": "New York, USA",
	}

	for name, address := range concerts {
		lat, lng, err := geocodeAddress(address)
		if err != nil {
			log.Printf("Erreur lors de la conversion de l'adresse %s : %v", address, err)
			continue
		}
		fmt.Printf("%s: Latitude %f, Longitude %f\n", name, lat, lng)
	}
}

func geocodeAddress(address string) (float64, float64, error) {
	encodedAddress := url.QueryEscape(address)
	requestURL := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", encodedAddress, apiKey)

	resp, err := http.Get(requestURL)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	var geocodeResp GeocodeResponse
	err = json.NewDecoder(resp.Body).Decode(&geocodeResp)
	if err != nil {
		return 0, 0, err
	}

	if len(geocodeResp.Results) == 0 {
		return 0, 0, fmt.Errorf("aucun résultat trouvé pour l'adresse %s", address)
	}

	return geocodeResp.Results[0].Geometry.Location.Lat, geocodeResp.Results[0].Geometry.Location.Lng, nil
}
