package core

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Fonction pour afficher les données de géolocalisation à partir d'une API
func ShowGeolocalisationFromAPI() {
	// Appeler l'API pour récupérer les données de géolocalisation
	// Code pour appeler l'API et obtenir les données de géolocalisation

	// Supposons que les données de géolocalisation soient stockées dans une liste appelée locations
	locations := []string{"North Carolina, USA", "Georgia, USA", "Los Angeles, USA", "Saitama, Japan", "Osaka, Japan", "Nagoya, Japan", "Penrose, New Zealand", "Dunedin, New Zealand"}

	// Créer une liste pour afficher les localisations
	locationList := widget.NewList(
		// Fonction pour obtenir le contenu de chaque élément de la liste
		func() int {
			return len(locations)
		},
		// Fonction pour créer chaque élément de la liste
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		// Fonction pour mettre à jour le contenu de chaque élément de la liste
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(locations[i])
		},
	)

	// Créer le contenu principal de la page de géolocalisation
	content := container.NewVBox(
		widget.NewLabel("Locations:"),
		locationList, // Ajouter la liste des localisations
	)

	// Créer une nouvelle fenêtre pour afficher les données de géolocalisation
	geoWindow := a.NewWindow("Geolocalisation")
	geoWindow.Resize(fyne.NewSize(400, 300))
	geoWindow.SetContent(content)

	// Afficher la fenêtre
	geoWindow.Show()
}
