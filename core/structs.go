package core

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"net/http"
)

// Structure de données représentant un artiste
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

// Structure de données représentant des lieux de concert
type Locations struct {
	Id        int    `json:"id"`
	Locations string `json:"locations"`
}

// Structure de données représentant des dates de concert
type Date struct {
	Id           int    `json:"id"`
	ConcertDates string `json:"concertDates"`
}

// Structure de données représentant une relation entre artiste, lieu et date de concert
type Relation struct {
	ArtistId   int `json:"artistId"`
	LocationId int `json:"locationId"`
	DateId     int `json:"dateId"`
}

// Fonction pour récupérer les données des artistes depuis l'API
func GetArtists() []Artist {
	// Effectue une requête GET vers l'API pour récupérer les données des artistes
	resp, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	// Ferme la réponse une fois que la fonction est terminée
	defer resp.Body.Close()

	var artists []Artist
	// Décode la réponse JSON dans la slice d'artistes
	json.NewDecoder(resp.Body).Decode(&artists)
	return artists
}

// Fonction pour récupérer les données des lieux de concert depuis l'API
func GetLocations() []Locations {
	resp, _ := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	defer resp.Body.Close()

	var locations []Locations
	json.NewDecoder(resp.Body).Decode(&locations)
	return locations
}

// Fonction pour récupérer les données des dates de concert depuis l'API
func GetDates() []Date {
	resp, _ := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	defer resp.Body.Close()

	var date []Date
	json.NewDecoder(resp.Body).Decode(&date)
	return date
}

// Fonction pour récupérer les relations entre artistes, lieux et dates de concert depuis l'API
func GetRelations() []Relation {
	resp, _ := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	defer resp.Body.Close()

	var relations []Relation
	json.NewDecoder(resp.Body).Decode(&relations)
	return relations
}
func getImageFromURL(urlStr string) *canvas.Image {
	resource, _ := fyne.LoadResourceFromURLString(urlStr)

	image := canvas.NewImageFromResource(resource)
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(200, 200)) // Adjust the size as needed
	return image
}
