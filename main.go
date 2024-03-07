package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Recherche...")

	// Exemple de données
	var bands Bands
	err := json.Unmarshal([]byte(sampleData), &bands)
	if err != nil {
		log.Fatal(err)
	}

	// Fonction de mise à jour des résultats en fonction de la recherche
	updateResults := func(query string) {
		filteredResults := filterBySearchQuery(bands, query)
		fmt.Println("Résultats de la recherche :", filteredResults)
		// Mettez à jour l'interface utilisateur avec les résultats de la recherche ici
	}

	// Réagir aux changements dans la barre de recherche
	searchEntry.OnChanged = func(query string) {
		updateResults(query)
	}

	w.SetContent(container.NewVBox(searchEntry, widget.NewLabel("MUUUUSICCCC!")))
	w.ShowAndRun()
}

// Fonction pour filtrer les résultats en fonction de la requête de recherche
func filterBySearchQuery(bands Bands, query string) []Artist {
	var filteredResults []Artist
	for _, band := range bands.Index {
		// Vérifier si le nom du groupe contient la requête de recherche (insensible à la casse)
		if strings.Contains(strings.ToLower(band.Name), strings.ToLower(query)) {
			filteredResults = append(filteredResults, band)
		}
	}
	return filteredResults
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
