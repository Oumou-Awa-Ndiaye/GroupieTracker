package core

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"net/http"
)

// Data structure representing an artist
type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	ImageURL     string   `json:"imageUrl"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	DateCreation int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	DatesConcert string   `json:"concertDates"`
}

// Data structure representing concert locations
type Locations struct {
	Id        int    `json:"id"`
	Locations string `json:"locations"`
}

// Data structure representing concert dates
type Date struct {
	Id           int    `json:"id"`
	ConcertDates string `json:"concertDates"`
}

// Data structure representing a relationship between artist, location, and concert date
type Relation struct {
	ArtistId   int `json:"artistId"`
	LocationId int `json:"locationId"`
	DateId     int `json:"dateId"`
}

// Function to fetch artist data from the API
func GetArtists() []Artist {
	// Performs a GET request to the API to retrieve artist data
	resp, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	// Closes the response once the function completes
	defer resp.Body.Close()

	var artists []Artist
	// Decodes the JSON response into the slice of artists
	json.NewDecoder(resp.Body).Decode(&artists)
	return artists
}

// Loads an image from a URL and prepares it for display in a Fyne application.
func getImageFromURL(urlStr string) *canvas.Image {
	resource, _ := fyne.LoadResourceFromURLString(urlStr)

	image := canvas.NewImageFromResource(resource)
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(200, 200)) // Adjust the size as needed
	return image
}
