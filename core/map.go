package core

import (
	"encoding/json"
	"fmt"
	"image"
	"net/http"
	"strconv"
	"strings"
)

type Map struct {
	latitude  float32
	longitude float32
	markers   []Marker
}

type Response struct {
	Results []struct {
		Longitude float64 `json:"lon"`
		Latitude  float64 `json:"lat"`
	} `json:"results"`
}

type Marker struct {
	lat      float32
	long     float32
	colorHex string
	icon     string
}

func Concertlocation(center string, zoom int) *Map {
	location := strings.Split(center, "-")
	city := strings.Replace(location[0], "_", "%20", -1)
	country := strings.Replace(location[1], "_", "%20", -1)
	url := "https://api.geoapify.com/v1/ipinfo?&apiKey=78d974bfd32a4904b1fd69a4a9354b4e" + city + "&country=" + country + "&format=json&apiKey=LA_CLE"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer res.Body.Close()
	var resp Response
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	if len(resp.Results) > 0 {
		lon := resp.Results[0].Longitude
		lat := resp.Results[0].Latitude
		fmt.Println("Longitude:", lon)
		fmt.Println("Latitude:", lat)
		return &Map{float32(lat), float32(lon), []Marker{}}
	} else {
		fmt.Println("il y a une erreur sur vos coordonnées")
		return &Map{1.0, 1.0, []Marker{}}
	}
}

func (m *Map) GetImg() image.Image {
	res, err := http.Get(m.GetURL())
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer res.Body.Close()

	fmt.Println("Map generation Status:", res.StatusCode)
	img, _, err := image.Decode(res.Body)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return nil
	}
	return img
}

func (m *Map) GetURL() string {
	const (
		baseURL = "https://maps.geoapify.com/v1/staticmap?style=osm-bright-smooth"
		width   = 800
		height  = 600
		zoom    = 4
		apiKey  = "LA_CLE"
	)

	url := fmt.Sprintf("%s&width=%d&height=%d&center=lonlat:%f,%f&zoom=%d",
		baseURL, width, height, m.longitude, m.latitude, zoom)

	if len(m.markers) > 0 {
		url += "&marker="
		for i, marker := range m.markers {
			url += fmt.Sprintf("lonlat:%f,%f;type:awesome;color:%s;size:x-large;icon:%s",
				marker.long, marker.lat, marker.colorHex, marker.icon)
			if i != len(m.markers)-1 {
				url += "|"
			}
		}
	}

	url += "&scaleFactor=2"
	url += "&apiKey=" + apiKey
	return url
}

func (m *Map) AddMarker(lat float32, long float32, color string, icon string) {
	m.markers = append(m.markers, Marker{lat, long, color, icon})
}

func (m *Map) GetLatitude() (lat float32) {
	return m.latitude
}

func (m *Map) GetLongitude() (long float32) {
	return m.longitude
}

func toString(val float32) string {
	return strconv.FormatFloat(float64(val), 'f', -1, 32)
}
