/*Partie Awa*/

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Groupie_Tracker")

	// Créer une barre de recherche
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher...")

	// Créer un widget de liste pour afficher les suggestions
	suggestionsList := widget.NewList(
		func() int {
			return 0 // Initialisation avec une liste vide
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText("") // La fonction de mise à jour sera définie ultérieurement
		},
	)

	// Créer un conteneur pour la barre de recherche et la liste de suggestions
	content := container.NewVBox(
		searchEntry,
		suggestionsList,
	)
	// Définir le contenu de la fenêtre
	w.SetContent(content)

	w.ShowAndRun()
}
