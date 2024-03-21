package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"GroupieTracker/core"
)

var (
	a           = app.New()
	w           = a.NewWindow("Groupie Tracker")
	artistsData []core.Artist
)

func main() {
	artistsData = core.GetArtists() // Utilisez la fonction GetArtists pour récupérer les données des artistes

	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(1000, 600))
	w.SetMainMenu(createMainMenu())

	showHomePage()

	w.ShowAndRun()
}

func createMainMenu() *fyne.MainMenu {
	return fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Quit", func() {
				a.Quit()
			}),
		),
		fyne.NewMenu("Help",
			fyne.NewMenuItem("About", func() {
				widget.NewModalPopUp(
					fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
						widget.NewLabel("Groupie Tracker  Par Ynov Paris  "),
					),
					w.Canvas(),
				).Show()
			}),
		),
	)
}

func showHomePage() {
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search for artists...")
	searchEntry.OnChanged = func(term string) {
		results := searchArtists(term)
		updateArtistGrid(results)
	}

	artistGrid = createArtistGrid()

	// Créer le contenu principal de la page
	content := container.NewVBox(
		createToolbar(),
		canvas.NewText("Groupie Tracker", theme.PrimaryColor()),
		layout.NewSpacer(),
		searchEntry, // Ajout DE la barre de recherche
		artistGrid,  // la grille des artistes
	)

	// Créer un conteneur avec défilement pour le contenu principal
	scrollContainer := container.NewVScroll(content)

	// Définir le contenu de la fenêtre
	w.SetContent(scrollContainer)
}

func updateArtistGrid(artists []core.Artist) {
	artistGrid.Objects = nil // Effacer la grille existante
	for _, artist := range artists {
		artistGrid.Add(createArtistCard(artist))
	}
	artistGrid.Refresh() // Rafraîchir la grille pour afficher les modifications
}
func createToolbar() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			showHomePage()
		}),
	)
}

var artistGrid *fyne.Container

func createArtistGrid() *fyne.Container {
	artistsData := core.GetArtists()

	numColumns := 1

	grid := container.NewGridWithColumns(numColumns)

	for _, artist := range artistsData {
		artistLabel := widget.NewLabel(artist.Name)

		grid.Add(artistLabel)
	}

	// Retourner la grille contenant les artistes
	return grid
}
func createArtistCard(artist core.Artist) fyne.CanvasObject {
	image := canvas.NewImageFromFile(artist.Image)
	image.FillMode = canvas.ImageFillContain

	nameLabel := widget.NewLabel(artist.Name)
	membersLabel := widget.NewLabel(fmt.Sprintf("Members: %s", strings.Join(artist.Members, ", ")))
	firstAlbumLabel := widget.NewLabel(fmt.Sprintf("First Album: %s", artist.FirstAlbum))
	creationDateLabel := widget.NewLabel(fmt.Sprintf("Creation Date: %d", artist.DateCreation))

	return container.NewVBox(
		image,
		nameLabel,
		membersLabel,
		firstAlbumLabel,
		creationDateLabel,
		widget.NewButton("View Details", func() {
			showArtistDetails(artist)
		}),
	)
}

func searchArtists(term string) []core.Artist {
	return core.Searchbar(term, artistsData)
}

func showArtistDetails(artist core.Artist) {
	widget.NewModalPopUp(
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
			canvas.NewText(artist.Name, theme.PrimaryColor()),
			canvas.NewImageFromFile(artist.Image),
			widget.NewLabel(fmt.Sprintf("Members: %s", strings.Join(artist.Members, ", "))),
			widget.NewLabel(fmt.Sprintf("First Album: %s", artist.FirstAlbum)),
			widget.NewLabel(fmt.Sprintf("Creation Date: %d", artist.DateCreation)),
		),
		w.Canvas(),
	).Show()
}
