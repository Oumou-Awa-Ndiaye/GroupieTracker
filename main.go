package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type Bands struct {
	Index []Artist `json:"index"`
}

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func main() {
	a := app.New()
	w := a.NewWindow("MUUUSIICCC")

	// Exemple de données
	var bands Bands
	err := json.Unmarshal([]byte(sampleData), &bands)
	if err != nil {
		log.Fatal(err)
	}

	// Exemple d'utilisation des fonctions de filtrage
	minCreation := 2000
	maxCreation := 2010
	minMembers := 3
	maxMembers := 5
	concertLocations := []string{"New York", "Los Angeles"}

	filteredByCreationDate := filterByCreationDate(bands, minCreation, maxCreation)
	fmt.Println("Groupes filtrés par date de création :", filteredByCreationDate)

	filteredByMembersCount := filterByMembersCount(bands, minMembers, maxMembers)
	fmt.Println("Groupes filtrés par nombre de membres :", filteredByMembersCount)

	filteredByConcertLocations := filterByConcertLocations(bands, concertLocations)
	fmt.Println("Groupes filtrés par lieux de concerts :", filteredByConcertLocations)

	w.SetContent(widget.NewLabel("MUUUUSICCCC!"))
	w.ShowAndRun()
}

// Fonction pour filtrer les groupes de musique par date de création
func filterByCreationDate(bands Bands, minCreationDate, maxCreationDate int) []Artist {
	var filteredBands []Artist
	for _, band := range bands.Index {
		if band.CreationDate >= minCreationDate && band.CreationDate <= maxCreationDate {
			filteredBands = append(filteredBands, band)
		}
	}
	return filteredBands
}

// Fonction pour filtrer les groupes de musique par nombre de membres
func filterByMembersCount(bands Bands, minMembers, maxMembers int) []Artist {
	var filteredBands []Artist
	for _, band := range bands.Index {
		numMembers := len(band.Members)
		if numMembers >= minMembers && numMembers <= maxMembers {
			filteredBands = append(filteredBands, band)
		}
	}
	return filteredBands
}

// Fonction pour filtrer les groupes de musique par lieux de concerts
func filterByConcertLocations(bands Bands, locations []string) []Artist {
	var filteredBands []Artist
	for _, band := range bands.Index {
		for _, loc := range locations {
			if strings.Contains(band.Locations, loc) {
				filteredBands = append(filteredBands, band)
				break
			}
		}
	}
	return filteredBands
}

// Données JSON d'exemple
const sampleData = `
{
	"index": [
		{
			"id": 1,
			"image": "image1.jpg",
			"name": "Groupe 1",
			"members": ["Membre 1", "Membre 2"],
			"creationDate": 2000,
			"firstAlbum": "Album 1",
			"locations": "New York, Londres",
			"concertDates": "2024-02-10, 2024-02-15",
			"relations": "Relation 1"
		},
		{
			"id": 2,
			"image": "image2.jpg",
			"name": "Groupe 2",
			"members": ["Membre 3", "Membre 4", "Membre 5"],
			"creationDate": 2010,
			"firstAlbum": "Album 2",
			"locations": "Los Angeles",
			"concertDates": "2024-02-12, 2024-02-17",
			"relations": "Relation 2"
		}
	]
}
`
